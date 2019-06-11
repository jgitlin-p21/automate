// Code generated by MockGen. DO NOT EDIT.
// Source: authz/v2/project.pb.go

// Package v2 is a generated GoMock package.
package v2

import (
	event "github.com/chef/automate/api/interservice/event"
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockProjectsClient is a mock of ProjectsClient interface
type MockProjectsClient struct {
	ctrl     *gomock.Controller
	recorder *MockProjectsClientMockRecorder
}

// MockProjectsClientMockRecorder is the mock recorder for MockProjectsClient
type MockProjectsClientMockRecorder struct {
	mock *MockProjectsClient
}

// NewMockProjectsClient creates a new mock instance
func NewMockProjectsClient(ctrl *gomock.Controller) *MockProjectsClient {
	mock := &MockProjectsClient{ctrl: ctrl}
	mock.recorder = &MockProjectsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProjectsClient) EXPECT() *MockProjectsClientMockRecorder {
	return m.recorder
}

// UpdateProject mocks base method
func (m *MockProjectsClient) UpdateProject(ctx context.Context, in *UpdateProjectReq, opts ...grpc.CallOption) (*UpdateProjectResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateProject", varargs...)
	ret0, _ := ret[0].(*UpdateProjectResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProject indicates an expected call of UpdateProject
func (mr *MockProjectsClientMockRecorder) UpdateProject(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockProjectsClient)(nil).UpdateProject), varargs...)
}

// CreateProject mocks base method
func (m *MockProjectsClient) CreateProject(ctx context.Context, in *CreateProjectReq, opts ...grpc.CallOption) (*CreateProjectResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateProject", varargs...)
	ret0, _ := ret[0].(*CreateProjectResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject
func (mr *MockProjectsClientMockRecorder) CreateProject(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockProjectsClient)(nil).CreateProject), varargs...)
}

// GetProject mocks base method
func (m *MockProjectsClient) GetProject(ctx context.Context, in *GetProjectReq, opts ...grpc.CallOption) (*GetProjectResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProject", varargs...)
	ret0, _ := ret[0].(*GetProjectResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProject indicates an expected call of GetProject
func (mr *MockProjectsClientMockRecorder) GetProject(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockProjectsClient)(nil).GetProject), varargs...)
}

// DeleteProject mocks base method
func (m *MockProjectsClient) DeleteProject(ctx context.Context, in *DeleteProjectReq, opts ...grpc.CallOption) (*DeleteProjectResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteProject", varargs...)
	ret0, _ := ret[0].(*DeleteProjectResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProject indicates an expected call of DeleteProject
func (mr *MockProjectsClientMockRecorder) DeleteProject(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockProjectsClient)(nil).DeleteProject), varargs...)
}

// ListProjects mocks base method
func (m *MockProjectsClient) ListProjects(ctx context.Context, in *ListProjectsReq, opts ...grpc.CallOption) (*ListProjectsResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjects", varargs...)
	ret0, _ := ret[0].(*ListProjectsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjects indicates an expected call of ListProjects
func (mr *MockProjectsClientMockRecorder) ListProjects(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockProjectsClient)(nil).ListProjects), varargs...)
}

// ListProjectsForIntrospection mocks base method
func (m *MockProjectsClient) ListProjectsForIntrospection(ctx context.Context, in *ListProjectsReq, opts ...grpc.CallOption) (*ListProjectsResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjectsForIntrospection", varargs...)
	ret0, _ := ret[0].(*ListProjectsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjectsForIntrospection indicates an expected call of ListProjectsForIntrospection
func (mr *MockProjectsClientMockRecorder) ListProjectsForIntrospection(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectsForIntrospection", reflect.TypeOf((*MockProjectsClient)(nil).ListProjectsForIntrospection), varargs...)
}

// HandleEvent mocks base method
func (m *MockProjectsClient) HandleEvent(ctx context.Context, in *event.EventMsg, opts ...grpc.CallOption) (*event.EventResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HandleEvent", varargs...)
	ret0, _ := ret[0].(*event.EventResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleEvent indicates an expected call of HandleEvent
func (mr *MockProjectsClientMockRecorder) HandleEvent(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleEvent", reflect.TypeOf((*MockProjectsClient)(nil).HandleEvent), varargs...)
}

// ProjectUpdateStatus mocks base method
func (m *MockProjectsClient) ProjectUpdateStatus(ctx context.Context, in *ProjectUpdateStatusReq, opts ...grpc.CallOption) (*ProjectUpdateStatusResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ProjectUpdateStatus", varargs...)
	ret0, _ := ret[0].(*ProjectUpdateStatusResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectUpdateStatus indicates an expected call of ProjectUpdateStatus
func (mr *MockProjectsClientMockRecorder) ProjectUpdateStatus(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectUpdateStatus", reflect.TypeOf((*MockProjectsClient)(nil).ProjectUpdateStatus), varargs...)
}

// ProjectUpdateCancel mocks base method
func (m *MockProjectsClient) ProjectUpdateCancel(ctx context.Context, in *ProjectUpdateStatusReq, opts ...grpc.CallOption) (*ProjectUpdateCancelResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ProjectUpdateCancel", varargs...)
	ret0, _ := ret[0].(*ProjectUpdateCancelResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectUpdateCancel indicates an expected call of ProjectUpdateCancel
func (mr *MockProjectsClientMockRecorder) ProjectUpdateCancel(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectUpdateCancel", reflect.TypeOf((*MockProjectsClient)(nil).ProjectUpdateCancel), varargs...)
}

// CreateRule mocks base method
func (m *MockProjectsClient) CreateRule(ctx context.Context, in *CreateRuleReq, opts ...grpc.CallOption) (*CreateRuleResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateRule", varargs...)
	ret0, _ := ret[0].(*CreateRuleResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRule indicates an expected call of CreateRule
func (mr *MockProjectsClientMockRecorder) CreateRule(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRule", reflect.TypeOf((*MockProjectsClient)(nil).CreateRule), varargs...)
}

// UpdateRule mocks base method
func (m *MockProjectsClient) UpdateRule(ctx context.Context, in *UpdateRuleReq, opts ...grpc.CallOption) (*UpdateRuleResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateRule", varargs...)
	ret0, _ := ret[0].(*UpdateRuleResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRule indicates an expected call of UpdateRule
func (mr *MockProjectsClientMockRecorder) UpdateRule(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRule", reflect.TypeOf((*MockProjectsClient)(nil).UpdateRule), varargs...)
}

// GetRule mocks base method
func (m *MockProjectsClient) GetRule(ctx context.Context, in *GetRuleReq, opts ...grpc.CallOption) (*GetRuleResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRule", varargs...)
	ret0, _ := ret[0].(*GetRuleResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRule indicates an expected call of GetRule
func (mr *MockProjectsClientMockRecorder) GetRule(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRule", reflect.TypeOf((*MockProjectsClient)(nil).GetRule), varargs...)
}

// ListRules mocks base method
func (m *MockProjectsClient) ListRules(ctx context.Context, in *ListRulesReq, opts ...grpc.CallOption) (*ListRulesResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRules", varargs...)
	ret0, _ := ret[0].(*ListRulesResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRules indicates an expected call of ListRules
func (mr *MockProjectsClientMockRecorder) ListRules(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRules", reflect.TypeOf((*MockProjectsClient)(nil).ListRules), varargs...)
}

// ListRulesForProject mocks base method
func (m *MockProjectsClient) ListRulesForProject(ctx context.Context, in *ListRulesForProjectReq, opts ...grpc.CallOption) (*ListRulesForProjectResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRulesForProject", varargs...)
	ret0, _ := ret[0].(*ListRulesForProjectResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRulesForProject indicates an expected call of ListRulesForProject
func (mr *MockProjectsClientMockRecorder) ListRulesForProject(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRulesForProject", reflect.TypeOf((*MockProjectsClient)(nil).ListRulesForProject), varargs...)
}

// DeleteRule mocks base method
func (m *MockProjectsClient) DeleteRule(ctx context.Context, in *DeleteRuleReq, opts ...grpc.CallOption) (*DeleteRuleResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteRule", varargs...)
	ret0, _ := ret[0].(*DeleteRuleResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRule indicates an expected call of DeleteRule
func (mr *MockProjectsClientMockRecorder) DeleteRule(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRule", reflect.TypeOf((*MockProjectsClient)(nil).DeleteRule), varargs...)
}

// ListRulesForAllProjects mocks base method
func (m *MockProjectsClient) ListRulesForAllProjects(ctx context.Context, in *ListRulesForAllProjectsReq, opts ...grpc.CallOption) (*ListRulesForAllProjectsResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRulesForAllProjects", varargs...)
	ret0, _ := ret[0].(*ListRulesForAllProjectsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRulesForAllProjects indicates an expected call of ListRulesForAllProjects
func (mr *MockProjectsClientMockRecorder) ListRulesForAllProjects(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRulesForAllProjects", reflect.TypeOf((*MockProjectsClient)(nil).ListRulesForAllProjects), varargs...)
}

// MockProjectsServer is a mock of ProjectsServer interface
type MockProjectsServer struct {
	ctrl     *gomock.Controller
	recorder *MockProjectsServerMockRecorder
}

// MockProjectsServerMockRecorder is the mock recorder for MockProjectsServer
type MockProjectsServerMockRecorder struct {
	mock *MockProjectsServer
}

// NewMockProjectsServer creates a new mock instance
func NewMockProjectsServer(ctrl *gomock.Controller) *MockProjectsServer {
	mock := &MockProjectsServer{ctrl: ctrl}
	mock.recorder = &MockProjectsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProjectsServer) EXPECT() *MockProjectsServerMockRecorder {
	return m.recorder
}

// UpdateProject mocks base method
func (m *MockProjectsServer) UpdateProject(arg0 context.Context, arg1 *UpdateProjectReq) (*UpdateProjectResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProject", arg0, arg1)
	ret0, _ := ret[0].(*UpdateProjectResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProject indicates an expected call of UpdateProject
func (mr *MockProjectsServerMockRecorder) UpdateProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockProjectsServer)(nil).UpdateProject), arg0, arg1)
}

// CreateProject mocks base method
func (m *MockProjectsServer) CreateProject(arg0 context.Context, arg1 *CreateProjectReq) (*CreateProjectResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", arg0, arg1)
	ret0, _ := ret[0].(*CreateProjectResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject
func (mr *MockProjectsServerMockRecorder) CreateProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockProjectsServer)(nil).CreateProject), arg0, arg1)
}

// GetProject mocks base method
func (m *MockProjectsServer) GetProject(arg0 context.Context, arg1 *GetProjectReq) (*GetProjectResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProject", arg0, arg1)
	ret0, _ := ret[0].(*GetProjectResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProject indicates an expected call of GetProject
func (mr *MockProjectsServerMockRecorder) GetProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockProjectsServer)(nil).GetProject), arg0, arg1)
}

// DeleteProject mocks base method
func (m *MockProjectsServer) DeleteProject(arg0 context.Context, arg1 *DeleteProjectReq) (*DeleteProjectResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", arg0, arg1)
	ret0, _ := ret[0].(*DeleteProjectResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProject indicates an expected call of DeleteProject
func (mr *MockProjectsServerMockRecorder) DeleteProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockProjectsServer)(nil).DeleteProject), arg0, arg1)
}

// ListProjects mocks base method
func (m *MockProjectsServer) ListProjects(arg0 context.Context, arg1 *ListProjectsReq) (*ListProjectsResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjects", arg0, arg1)
	ret0, _ := ret[0].(*ListProjectsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjects indicates an expected call of ListProjects
func (mr *MockProjectsServerMockRecorder) ListProjects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockProjectsServer)(nil).ListProjects), arg0, arg1)
}

// ListProjectsForIntrospection mocks base method
func (m *MockProjectsServer) ListProjectsForIntrospection(arg0 context.Context, arg1 *ListProjectsReq) (*ListProjectsResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjectsForIntrospection", arg0, arg1)
	ret0, _ := ret[0].(*ListProjectsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjectsForIntrospection indicates an expected call of ListProjectsForIntrospection
func (mr *MockProjectsServerMockRecorder) ListProjectsForIntrospection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectsForIntrospection", reflect.TypeOf((*MockProjectsServer)(nil).ListProjectsForIntrospection), arg0, arg1)
}

// HandleEvent mocks base method
func (m *MockProjectsServer) HandleEvent(arg0 context.Context, arg1 *event.EventMsg) (*event.EventResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleEvent", arg0, arg1)
	ret0, _ := ret[0].(*event.EventResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleEvent indicates an expected call of HandleEvent
func (mr *MockProjectsServerMockRecorder) HandleEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleEvent", reflect.TypeOf((*MockProjectsServer)(nil).HandleEvent), arg0, arg1)
}

// ProjectUpdateStatus mocks base method
func (m *MockProjectsServer) ProjectUpdateStatus(arg0 context.Context, arg1 *ProjectUpdateStatusReq) (*ProjectUpdateStatusResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectUpdateStatus", arg0, arg1)
	ret0, _ := ret[0].(*ProjectUpdateStatusResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectUpdateStatus indicates an expected call of ProjectUpdateStatus
func (mr *MockProjectsServerMockRecorder) ProjectUpdateStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectUpdateStatus", reflect.TypeOf((*MockProjectsServer)(nil).ProjectUpdateStatus), arg0, arg1)
}

// ProjectUpdateCancel mocks base method
func (m *MockProjectsServer) ProjectUpdateCancel(arg0 context.Context, arg1 *ProjectUpdateStatusReq) (*ProjectUpdateCancelResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectUpdateCancel", arg0, arg1)
	ret0, _ := ret[0].(*ProjectUpdateCancelResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectUpdateCancel indicates an expected call of ProjectUpdateCancel
func (mr *MockProjectsServerMockRecorder) ProjectUpdateCancel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectUpdateCancel", reflect.TypeOf((*MockProjectsServer)(nil).ProjectUpdateCancel), arg0, arg1)
}

// CreateRule mocks base method
func (m *MockProjectsServer) CreateRule(arg0 context.Context, arg1 *CreateRuleReq) (*CreateRuleResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRule", arg0, arg1)
	ret0, _ := ret[0].(*CreateRuleResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRule indicates an expected call of CreateRule
func (mr *MockProjectsServerMockRecorder) CreateRule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRule", reflect.TypeOf((*MockProjectsServer)(nil).CreateRule), arg0, arg1)
}

// UpdateRule mocks base method
func (m *MockProjectsServer) UpdateRule(arg0 context.Context, arg1 *UpdateRuleReq) (*UpdateRuleResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRule", arg0, arg1)
	ret0, _ := ret[0].(*UpdateRuleResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRule indicates an expected call of UpdateRule
func (mr *MockProjectsServerMockRecorder) UpdateRule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRule", reflect.TypeOf((*MockProjectsServer)(nil).UpdateRule), arg0, arg1)
}

// GetRule mocks base method
func (m *MockProjectsServer) GetRule(arg0 context.Context, arg1 *GetRuleReq) (*GetRuleResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRule", arg0, arg1)
	ret0, _ := ret[0].(*GetRuleResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRule indicates an expected call of GetRule
func (mr *MockProjectsServerMockRecorder) GetRule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRule", reflect.TypeOf((*MockProjectsServer)(nil).GetRule), arg0, arg1)
}

// ListRules mocks base method
func (m *MockProjectsServer) ListRules(arg0 context.Context, arg1 *ListRulesReq) (*ListRulesResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRules", arg0, arg1)
	ret0, _ := ret[0].(*ListRulesResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRules indicates an expected call of ListRules
func (mr *MockProjectsServerMockRecorder) ListRules(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRules", reflect.TypeOf((*MockProjectsServer)(nil).ListRules), arg0, arg1)
}

// ListRulesForProject mocks base method
func (m *MockProjectsServer) ListRulesForProject(arg0 context.Context, arg1 *ListRulesForProjectReq) (*ListRulesForProjectResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRulesForProject", arg0, arg1)
	ret0, _ := ret[0].(*ListRulesForProjectResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRulesForProject indicates an expected call of ListRulesForProject
func (mr *MockProjectsServerMockRecorder) ListRulesForProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRulesForProject", reflect.TypeOf((*MockProjectsServer)(nil).ListRulesForProject), arg0, arg1)
}

// DeleteRule mocks base method
func (m *MockProjectsServer) DeleteRule(arg0 context.Context, arg1 *DeleteRuleReq) (*DeleteRuleResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRule", arg0, arg1)
	ret0, _ := ret[0].(*DeleteRuleResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRule indicates an expected call of DeleteRule
func (mr *MockProjectsServerMockRecorder) DeleteRule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRule", reflect.TypeOf((*MockProjectsServer)(nil).DeleteRule), arg0, arg1)
}

// ListRulesForAllProjects mocks base method
func (m *MockProjectsServer) ListRulesForAllProjects(arg0 context.Context, arg1 *ListRulesForAllProjectsReq) (*ListRulesForAllProjectsResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRulesForAllProjects", arg0, arg1)
	ret0, _ := ret[0].(*ListRulesForAllProjectsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRulesForAllProjects indicates an expected call of ListRulesForAllProjects
func (mr *MockProjectsServerMockRecorder) ListRulesForAllProjects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRulesForAllProjects", reflect.TypeOf((*MockProjectsServer)(nil).ListRulesForAllProjects), arg0, arg1)
}
