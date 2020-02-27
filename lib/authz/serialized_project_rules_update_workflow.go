package authz

import (
	"fmt"
	"strings"
	"time"

	"github.com/chef/automate/lib/cereal"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type SerializedProjectUpdateWorkflowExecutor struct {
	PollInterval time.Duration
}

type SerializedProjectUpdateTask struct {
	Priority int64
	Params   map[string]string
}

type SerializedProjectUpdateTaskID string

const SerializedProjectUpdateTaskIDUnknown = ""

type SerializedProjectUpdateTaskState string

const (
	SerializedProjectUpdateTaskRunning SerializedProjectUpdateTaskState = "running"
	SerializedProjectUpdateTaskFailed  SerializedProjectUpdateTaskState = "failed"
	SerializedProjectUpdateTaskSuccess SerializedProjectUpdateTaskState = "success"
)

type SerializedProjectUpdateTaskStatus struct {
	State                 SerializedProjectUpdateTaskState
	PercentageComplete    float32
	EstimatedEndTimeInSec int64
	Metadata              map[string]string
	Error                 string
}

type SerializedProjectUpdate interface {
	ListTasks() ([]SerializedProjectUpdateTask, error)
	RunTask(params map[string]string) (SerializedProjectUpdateTaskID, SerializedProjectUpdateTaskStatus, error)
	MonitorTask(id SerializedProjectUpdateTaskID) (SerializedProjectUpdateTaskStatus, error)
	CancelTask(id SerializedProjectUpdateTaskID) error
}

type SerializedProjectUpdateWorkflowParams struct {
	ProjectUpdateID string
	DomainServices  []string
}
type SerializedProjectUpdateWorkflowPhase string

const (
	SerializedProjectUpdateWorkflowPhaseStarting  = "starting"
	SerializedProjectUpdateWorkflowPhaseRunning   = "running"
	SerializedProjectUpdateWorkflowPhaseComplete  = "complete"
	SerializedProjectUpdateWorkflowPhaseCanceling = "canceling"
)

type serializedProjectUpdateWorkflowPhaseStartingData struct {
	ListTasksResults map[string]SerializedProjectUpdateListTasksTaskResult
}

type serializedProjectUpdateTask struct {
	SerializedProjectUpdateTask `json:"t"`
	DomainService               string `json:"svc"`
}

type serializedProjectUpdateRunningTask struct {
	ID   SerializedProjectUpdateTaskID
	Task serializedProjectUpdateTask
}

type serializedProjectUpdateWorkflowPhaseRunningData struct {
	// TODO: need to add domain service to each of these
	RunningTask    serializedProjectUpdateRunningTask
	RemainingTasks []serializedProjectUpdateTask
}

type serializedProjectUpdateWorkflowState struct {
	ProjectUpdateID string
	DomainServices  []string
	Phase           SerializedProjectUpdateWorkflowPhase
	Starting        serializedProjectUpdateWorkflowPhaseStartingData
	Running         serializedProjectUpdateWorkflowPhaseRunningData
	MonitorFailures int
}

func (s *serializedProjectUpdateWorkflowState) isStartPhaseDone() bool {
	if s.Phase != SerializedProjectUpdateWorkflowPhaseStarting {
		return true
	}

	if s.Starting.ListTasksResults == nil {
		return false
	}

	for _, svc := range s.DomainServices {
		if _, ok := s.Starting.ListTasksResults[svc]; !ok {
			return false
		}
	}

	return true
}

type SerializedProjectUpdateListTasksTaskParams struct{}
type SerializedProjectUpdateListTasksTaskResult struct {
	Tasks []SerializedProjectUpdateTask
}

type SerializedProjectUpdateRunTaskTaskParams struct {
	Params map[string]string
}
type SerializedProjectUpdateRunTaskTaskResult struct {
	ID     SerializedProjectUpdateTaskID
	Status SerializedProjectUpdateTaskStatus
}

type SerializedProjectUpdateMonitorTaskTaskParams struct {
	ID SerializedProjectUpdateTaskID
}
type SerializedProjectUpdateMonitorTaskTaskResult struct {
	Status SerializedProjectUpdateTaskStatus
}

type SerializedProjectUpdateCancelTaskTaskParams struct {
	ID SerializedProjectUpdateTaskID
}
type SerializedProjectUpdateCancelTaskTaskResult struct {
	Status SerializedProjectUpdateTaskStatus
}

func (m *SerializedProjectUpdateWorkflowExecutor) OnStart(
	w cereal.WorkflowInstance, ev cereal.StartEvent) cereal.Decision {

	params := SerializedProjectUpdateWorkflowParams{}
	if err := w.GetParameters(&params); err != nil {
		return w.Fail(err)
	}

	if len(params.DomainServices) == 0 {
		return w.Fail(errors.New("no services specified"))
	}

	logrus.Infof("Started SerializedProjectUpdateWorkflow for %s",
		params.ProjectUpdateID)

	for _, svc := range params.DomainServices {
		if err := w.EnqueueTask(m.listTasksTaskNameForService(svc), nil); err != nil {
			return w.Fail(err)
		}
	}

	return w.Continue(serializedProjectUpdateWorkflowState{
		ProjectUpdateID: params.ProjectUpdateID,
		DomainServices:  params.DomainServices,
		Phase:           SerializedProjectUpdateWorkflowPhaseStarting,
	})
}

func (m *SerializedProjectUpdateWorkflowExecutor) OnTaskComplete(
	w cereal.WorkflowInstance, ev cereal.TaskCompleteEvent) cereal.Decision {

	state := serializedProjectUpdateWorkflowState{}
	if err := w.GetPayload(&state); err != nil {
		logrus.WithError(err).Error("Failed to deserialize payload")
		return w.Fail(err)
	}
	switch m.taskType(ev.TaskName) {
	case SerializedProjectUpdateListTaskType:
		return m.handleListTasksComplete(&state, w, ev)
	case SerializedProjectUpdateRunTaskType:
		return m.handleRunTaskComplete(&state, w, ev)
	case SerializedProjectUpdateMonitorTaskType:
		return m.handleMonitorTaskComplete(&state, w, ev)
	}

	return w.Fail(errors.New("somethign"))
}

func (m *SerializedProjectUpdateWorkflowExecutor) OnCancel(
	w cereal.WorkflowInstance, ev cereal.CancelEvent) cereal.Decision {
	state := serializedProjectUpdateWorkflowState{}
	if err := w.GetPayload(&state); err != nil {
		logrus.WithError(err).Error("Failed to deserialize payload")
		return w.Fail(err)
	}
	if state.Phase == SerializedProjectUpdateWorkflowPhaseCanceling {
		return w.Continue(state)
	}
	return m.cancelRunningTask(&state, w)
}

func (m *SerializedProjectUpdateWorkflowExecutor) cancelRunningTask(
	state *serializedProjectUpdateWorkflowState, w cereal.WorkflowInstance) cereal.Decision {
	state.Phase = SerializedProjectUpdateWorkflowPhaseCanceling
	state.Running.RemainingTasks = []serializedProjectUpdateTask{}
	if state.Running.RunningTask.ID != SerializedProjectUpdateTaskIDUnknown {
		w.EnqueueTask(
			m.cancelTaskTaskNameForService(state.Running.RunningTask.Task.DomainService),
			SerializedProjectUpdateCancelTaskTaskParams{
				ID: state.Running.RunningTask.ID,
			},
		)
	}
	return w.Continue(state)
}

func (m *SerializedProjectUpdateWorkflowExecutor) handleListTasksComplete(
	state *serializedProjectUpdateWorkflowState, w cereal.WorkflowInstance,
	ev cereal.TaskCompleteEvent) cereal.Decision {
	if ev.Result.Err() != nil {
		return w.Fail(ev.Result.Err())
	}

	svc, err := m.svcName(ev.TaskName)
	if err != nil {
		return w.Fail(err)
	}

	var res SerializedProjectUpdateListTasksTaskResult
	if err := ev.Result.Get(&res); err != nil {
		return w.Fail(err)
	}

	if state.Starting.ListTasksResults == nil {
		state.Starting.ListTasksResults = make(map[string]SerializedProjectUpdateListTasksTaskResult)
	}

	state.Starting.ListTasksResults[svc] = res

	if state.isStartPhaseDone() {
		tasks := []serializedProjectUpdateTask{}
		for svc, resForSvc := range state.Starting.ListTasksResults {
			for _, t := range resForSvc.Tasks {
				tasks = append(tasks, serializedProjectUpdateTask{
					DomainService:               svc,
					SerializedProjectUpdateTask: t,
				})
			}
		}

		if len(tasks) == 0 {
			state.Phase = SerializedProjectUpdateWorkflowPhaseComplete
			return w.Complete(cereal.WithResult(state))
		}

		state.Phase = SerializedProjectUpdateWorkflowPhaseRunning
		state.Running.RunningTask = serializedProjectUpdateRunningTask{
			ID:   SerializedProjectUpdateTaskIDUnknown,
			Task: tasks[0],
		}
		if err := w.EnqueueTask(m.runTaskTaskNameForService(tasks[0].DomainService),
			SerializedProjectUpdateRunTaskTaskParams{
				Params: tasks[0].Params,
			}); err != nil {
			return w.Fail(err)
		}
		state.Running.RemainingTasks = tasks[1:]
	}

	return w.Continue(state)
}

func (m *SerializedProjectUpdateWorkflowExecutor) handleRunTaskComplete(
	state *serializedProjectUpdateWorkflowState, w cereal.WorkflowInstance,
	ev cereal.TaskCompleteEvent) cereal.Decision {

	if ev.Result.Err() != nil {
		return w.Fail(ev.Result.Err())
	}

	var res SerializedProjectUpdateRunTaskTaskResult
	if err := ev.Result.Get(&res); err != nil {
		return w.Fail(err)
	}

	switch res.Status.State {
	case SerializedProjectUpdateTaskSuccess:
		return m.startNextTask(state, w)
	case SerializedProjectUpdateTaskFailed:
		return w.Fail(errors.New(res.Status.Error))
	case SerializedProjectUpdateTaskRunning:
		state.Running.RunningTask.ID = res.ID
	default:
		return w.Fail(errors.New("Unknown task state"))
	}

	if state.Phase == SerializedProjectUpdateWorkflowPhaseCanceling {
		return m.cancelRunningTask(state, w)
	}

	// Start task monitor
	return m.startMonitorTask(state, w)
}

const maxMonitorFailures = 10

func (m *SerializedProjectUpdateWorkflowExecutor) handleMonitorTaskComplete(
	state *serializedProjectUpdateWorkflowState, w cereal.WorkflowInstance,
	ev cereal.TaskCompleteEvent) cereal.Decision {

	var params SerializedProjectUpdateMonitorTaskTaskResult
	if err := ev.Result.GetParameters(&params); err != nil {
		return w.Fail(err)
	}

	var res SerializedProjectUpdateMonitorTaskTaskResult
	if err := ev.Result.Get(&res); err != nil {
		state.MonitorFailures++
		if state.MonitorFailures > maxMonitorFailures {

		}
		return m.startMonitorTask(state, w)
	} else {
		state.MonitorFailures = 0
		switch res.Status.State {
		case SerializedProjectUpdateTaskSuccess:
			return m.startNextTask(state, w)
		case SerializedProjectUpdateTaskFailed:
			return w.Fail(errors.New(res.Status.Error))
		case SerializedProjectUpdateTaskRunning:
			return m.startMonitorTask(state, w)
		default:
			return w.Fail(errors.New("Unknown task state"))
		}
	}

}

const maxFailureBackoffDuration = 2 * time.Minute

func (m *SerializedProjectUpdateWorkflowExecutor) startMonitorTask(
	state *serializedProjectUpdateWorkflowState, w cereal.WorkflowInstance) cereal.Decision {

	if state.Phase == SerializedProjectUpdateWorkflowPhaseRunning {
		failureBackoff := time.Duration(0)
		if state.MonitorFailures > 0 {
			failureBackoff = (1 << state.MonitorFailures) * time.Second
			if failureBackoff > maxFailureBackoffDuration {
				failureBackoff = maxFailureBackoffDuration
			}
		}

		w.EnqueueTask(m.monitorTaskTaskNameForService(state.Running.RunningTask.Task.DomainService), // nolint: errcheck
			SerializedProjectUpdateMonitorTaskTaskParams{
				ID: state.Running.RunningTask.ID,
			}, cereal.StartAfter(time.Now().Add(m.PollInterval).Add(failureBackoff)))
	}

	return w.Continue(state)
}

func (m *SerializedProjectUpdateWorkflowExecutor) startNextTask(
	state *serializedProjectUpdateWorkflowState, w cereal.WorkflowInstance) cereal.Decision {

	if state.Phase == SerializedProjectUpdateWorkflowPhaseCanceling {
		return w.Complete(cereal.WithResult(state))
	}

	if len(state.Running.RemainingTasks) == 0 {
		state.Phase = SerializedProjectUpdateWorkflowPhaseComplete
		return w.Complete(cereal.WithResult(state))
	}

	state.Phase = SerializedProjectUpdateWorkflowPhaseRunning
	state.Running.RunningTask = serializedProjectUpdateRunningTask{
		ID:   SerializedProjectUpdateTaskIDUnknown,
		Task: state.Running.RemainingTasks[0],
	}
	if err := w.EnqueueTask(m.runTaskTaskNameForService(state.Running.RunningTask.Task.DomainService),
		SerializedProjectUpdateRunTaskTaskParams{
			Params: state.Running.RunningTask.Task.Params,
		}); err != nil {
		return w.Fail(err)
	}
	state.Running.RemainingTasks = state.Running.RemainingTasks[1:]

	return w.Continue(state)
}

const (
	SerializedProjectUpdateListTaskType    = "project-update-list-tasks"
	SerializedProjectUpdateRunTaskType     = "project-update-run-task"
	SerializedProjectUpdateMonitorTaskType = "project-update-monitor-task"
	SerializedProjectUpdateCancelTaskType  = "project-update-cancel-task"
)

func (m *SerializedProjectUpdateWorkflowExecutor) taskType(taskName cereal.TaskName) string {
	parts := strings.SplitN(taskName.String(), "/", 2)
	return parts[0]
}

func (m *SerializedProjectUpdateWorkflowExecutor) svcName(taskName cereal.TaskName) (string, error) {
	parts := strings.Split(taskName.String(), "/")
	if len(parts) < 2 {
		return "", errors.New("Failed to parse task name for service")
	}
	return parts[1], nil
}

func (m *SerializedProjectUpdateWorkflowExecutor) listTasksTaskNameForService(
	svc string) cereal.TaskName {
	return cereal.NewTaskName(fmt.Sprintf("%s/%s", SerializedProjectUpdateListTaskType, svc))
}

func (m *SerializedProjectUpdateWorkflowExecutor) runTaskTaskNameForService(
	svc string) cereal.TaskName {
	return cereal.NewTaskName(fmt.Sprintf("%s/%s", SerializedProjectUpdateRunTaskType, svc))
}

func (m *SerializedProjectUpdateWorkflowExecutor) monitorTaskTaskNameForService(
	svc string) cereal.TaskName {
	return cereal.NewTaskName(fmt.Sprintf("%s/%s", SerializedProjectUpdateMonitorTaskType, svc))
}

func (m *SerializedProjectUpdateWorkflowExecutor) cancelTaskTaskNameForService(
	svc string) cereal.TaskName {
	return cereal.NewTaskName(fmt.Sprintf("%s/%s", SerializedProjectUpdateCancelTaskType, svc))
}
