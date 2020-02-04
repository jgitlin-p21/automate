// Code generated by MockGen. DO NOT EDIT.
// Source: ../gateway/clients.go

// Package mock_gateway is a generated GoMock package.
package mock_gateway

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	applications "github.com/chef/automate/api/external/applications"
	data_feed "github.com/chef/automate/api/external/data_feed"
	secrets "github.com/chef/automate/api/external/secrets"
	authn "github.com/chef/automate/api/interservice/authn"
	authz "github.com/chef/automate/api/interservice/authz"
	v2 "github.com/chef/automate/api/interservice/authz/v2"
	service "github.com/chef/automate/api/interservice/cfgmgmt/service"
	data_lifecycle "github.com/chef/automate/api/interservice/data_lifecycle"
	deployment "github.com/chef/automate/api/interservice/deployment"
	event_feed "github.com/chef/automate/api/interservice/event_feed"
	service0 "github.com/chef/automate/api/interservice/infra_proxy/service"
	ingest "github.com/chef/automate/api/interservice/ingest"
	license_control "github.com/chef/automate/api/interservice/license_control"
	local_user "github.com/chef/automate/api/interservice/local_user"
	manager "github.com/chef/automate/api/interservice/nodemanager/manager"
	nodes "github.com/chef/automate/api/interservice/nodemanager/nodes"
	v20 "github.com/chef/automate/api/interservice/teams/v2"
	jobs "github.com/chef/automate/components/compliance-service/api/jobs"
	profiles "github.com/chef/automate/components/compliance-service/api/profiles"
	reporting "github.com/chef/automate/components/compliance-service/api/reporting"
	stats "github.com/chef/automate/components/compliance-service/api/stats"
	version "github.com/chef/automate/components/compliance-service/api/version"
	ingest0 "github.com/chef/automate/components/compliance-service/ingest/ingest"
	api "github.com/chef/automate/components/notifications-client/api"
	notifier "github.com/chef/automate/components/notifications-client/notifier"
)

// MockClientsFactory is a mock of ClientsFactory interface
type MockClientsFactory struct {
	ctrl     *gomock.Controller
	recorder *MockClientsFactoryMockRecorder
}

// MockClientsFactoryMockRecorder is the mock recorder for MockClientsFactory
type MockClientsFactoryMockRecorder struct {
	mock *MockClientsFactory
}

// NewMockClientsFactory creates a new mock instance
func NewMockClientsFactory(ctrl *gomock.Controller) *MockClientsFactory {
	mock := &MockClientsFactory{ctrl: ctrl}
	mock.recorder = &MockClientsFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientsFactory) EXPECT() *MockClientsFactoryMockRecorder {
	return m.recorder
}

// CfgMgmtClient mocks base method
func (m *MockClientsFactory) CfgMgmtClient() (service.CfgMgmtClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CfgMgmtClient")
	ret0, _ := ret[0].(service.CfgMgmtClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CfgMgmtClient indicates an expected call of CfgMgmtClient
func (mr *MockClientsFactoryMockRecorder) CfgMgmtClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CfgMgmtClient", reflect.TypeOf((*MockClientsFactory)(nil).CfgMgmtClient))
}

// IngestStatusClient mocks base method
func (m *MockClientsFactory) IngestStatusClient() (ingest.IngestStatusClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IngestStatusClient")
	ret0, _ := ret[0].(ingest.IngestStatusClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IngestStatusClient indicates an expected call of IngestStatusClient
func (mr *MockClientsFactoryMockRecorder) IngestStatusClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IngestStatusClient", reflect.TypeOf((*MockClientsFactory)(nil).IngestStatusClient))
}

// ChefIngesterClient mocks base method
func (m *MockClientsFactory) ChefIngesterClient() (ingest.ChefIngesterClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChefIngesterClient")
	ret0, _ := ret[0].(ingest.ChefIngesterClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChefIngesterClient indicates an expected call of ChefIngesterClient
func (mr *MockClientsFactoryMockRecorder) ChefIngesterClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChefIngesterClient", reflect.TypeOf((*MockClientsFactory)(nil).ChefIngesterClient))
}

// ChefIngesterJobSchedulerClient mocks base method
func (m *MockClientsFactory) ChefIngesterJobSchedulerClient() (ingest.JobSchedulerClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChefIngesterJobSchedulerClient")
	ret0, _ := ret[0].(ingest.JobSchedulerClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChefIngesterJobSchedulerClient indicates an expected call of ChefIngesterJobSchedulerClient
func (mr *MockClientsFactoryMockRecorder) ChefIngesterJobSchedulerClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChefIngesterJobSchedulerClient", reflect.TypeOf((*MockClientsFactory)(nil).ChefIngesterJobSchedulerClient))
}

// ComplianceIngesterClient mocks base method
func (m *MockClientsFactory) ComplianceIngesterClient() (ingest0.ComplianceIngesterClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComplianceIngesterClient")
	ret0, _ := ret[0].(ingest0.ComplianceIngesterClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComplianceIngesterClient indicates an expected call of ComplianceIngesterClient
func (mr *MockClientsFactoryMockRecorder) ComplianceIngesterClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComplianceIngesterClient", reflect.TypeOf((*MockClientsFactory)(nil).ComplianceIngesterClient))
}

// NotificationsClient mocks base method
func (m *MockClientsFactory) NotificationsClient() (api.NotificationsClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotificationsClient")
	ret0, _ := ret[0].(api.NotificationsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotificationsClient indicates an expected call of NotificationsClient
func (mr *MockClientsFactoryMockRecorder) NotificationsClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotificationsClient", reflect.TypeOf((*MockClientsFactory)(nil).NotificationsClient))
}

// AuthenticationClient mocks base method
func (m *MockClientsFactory) AuthenticationClient() (authn.AuthenticationClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticationClient")
	ret0, _ := ret[0].(authn.AuthenticationClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthenticationClient indicates an expected call of AuthenticationClient
func (mr *MockClientsFactoryMockRecorder) AuthenticationClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticationClient", reflect.TypeOf((*MockClientsFactory)(nil).AuthenticationClient))
}

// AuthorizationClient mocks base method
func (m *MockClientsFactory) AuthorizationClient() (authz.AuthorizationClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizationClient")
	ret0, _ := ret[0].(authz.AuthorizationClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthorizationClient indicates an expected call of AuthorizationClient
func (mr *MockClientsFactoryMockRecorder) AuthorizationClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizationClient", reflect.TypeOf((*MockClientsFactory)(nil).AuthorizationClient))
}

// AuthorizationV2Client mocks base method
func (m *MockClientsFactory) AuthorizationV2Client() (v2.AuthorizationClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizationV2Client")
	ret0, _ := ret[0].(v2.AuthorizationClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthorizationV2Client indicates an expected call of AuthorizationV2Client
func (mr *MockClientsFactoryMockRecorder) AuthorizationV2Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizationV2Client", reflect.TypeOf((*MockClientsFactory)(nil).AuthorizationV2Client))
}

// PoliciesClient mocks base method
func (m *MockClientsFactory) PoliciesClient() (v2.PoliciesClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PoliciesClient")
	ret0, _ := ret[0].(v2.PoliciesClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PoliciesClient indicates an expected call of PoliciesClient
func (mr *MockClientsFactoryMockRecorder) PoliciesClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PoliciesClient", reflect.TypeOf((*MockClientsFactory)(nil).PoliciesClient))
}

// ProjectsClient mocks base method
func (m *MockClientsFactory) ProjectsClient() (v2.ProjectsClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectsClient")
	ret0, _ := ret[0].(v2.ProjectsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectsClient indicates an expected call of ProjectsClient
func (mr *MockClientsFactoryMockRecorder) ProjectsClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectsClient", reflect.TypeOf((*MockClientsFactory)(nil).ProjectsClient))
}

// TeamsClient mocks base method
func (m *MockClientsFactory) TeamsClient() (v20.TeamsV2Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamsClient")
	ret0, _ := ret[0].(v20.TeamsV2Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamsClient indicates an expected call of TeamsClient
func (mr *MockClientsFactoryMockRecorder) TeamsClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamsClient", reflect.TypeOf((*MockClientsFactory)(nil).TeamsClient))
}

// TokensMgmtClient mocks base method
func (m *MockClientsFactory) TokensMgmtClient() (authn.TokensMgmtClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokensMgmtClient")
	ret0, _ := ret[0].(authn.TokensMgmtClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokensMgmtClient indicates an expected call of TokensMgmtClient
func (mr *MockClientsFactoryMockRecorder) TokensMgmtClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokensMgmtClient", reflect.TypeOf((*MockClientsFactory)(nil).TokensMgmtClient))
}

// UsersMgmtClient mocks base method
func (m *MockClientsFactory) UsersMgmtClient() (local_user.UsersMgmtClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UsersMgmtClient")
	ret0, _ := ret[0].(local_user.UsersMgmtClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UsersMgmtClient indicates an expected call of UsersMgmtClient
func (mr *MockClientsFactoryMockRecorder) UsersMgmtClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsersMgmtClient", reflect.TypeOf((*MockClientsFactory)(nil).UsersMgmtClient))
}

// Notifier mocks base method
func (m *MockClientsFactory) Notifier() (notifier.Notifier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Notifier")
	ret0, _ := ret[0].(notifier.Notifier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Notifier indicates an expected call of Notifier
func (mr *MockClientsFactoryMockRecorder) Notifier() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notifier", reflect.TypeOf((*MockClientsFactory)(nil).Notifier))
}

// ApplicationsClient mocks base method
func (m *MockClientsFactory) ApplicationsClient() (applications.ApplicationsServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationsClient")
	ret0, _ := ret[0].(applications.ApplicationsServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationsClient indicates an expected call of ApplicationsClient
func (mr *MockClientsFactoryMockRecorder) ApplicationsClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationsClient", reflect.TypeOf((*MockClientsFactory)(nil).ApplicationsClient))
}

// SecretClient mocks base method
func (m *MockClientsFactory) SecretClient() (secrets.SecretsServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretClient")
	ret0, _ := ret[0].(secrets.SecretsServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecretClient indicates an expected call of SecretClient
func (mr *MockClientsFactoryMockRecorder) SecretClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretClient", reflect.TypeOf((*MockClientsFactory)(nil).SecretClient))
}

// DatafeedClient mocks base method
func (m *MockClientsFactory) DatafeedClient() (data_feed.DatafeedServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatafeedClient")
	ret0, _ := ret[0].(data_feed.DatafeedServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DatafeedClient indicates an expected call of DatafeedClient
func (mr *MockClientsFactoryMockRecorder) DatafeedClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatafeedClient", reflect.TypeOf((*MockClientsFactory)(nil).DatafeedClient))
}

// NodesClient mocks base method
func (m *MockClientsFactory) NodesClient() (nodes.NodesServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodesClient")
	ret0, _ := ret[0].(nodes.NodesServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodesClient indicates an expected call of NodesClient
func (mr *MockClientsFactoryMockRecorder) NodesClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodesClient", reflect.TypeOf((*MockClientsFactory)(nil).NodesClient))
}

// FeedClient mocks base method
func (m *MockClientsFactory) FeedClient() (event_feed.EventFeedServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FeedClient")
	ret0, _ := ret[0].(event_feed.EventFeedServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FeedClient indicates an expected call of FeedClient
func (mr *MockClientsFactoryMockRecorder) FeedClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FeedClient", reflect.TypeOf((*MockClientsFactory)(nil).FeedClient))
}

// ComplianceReportingServiceClient mocks base method
func (m *MockClientsFactory) ComplianceReportingServiceClient() (reporting.ReportingServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComplianceReportingServiceClient")
	ret0, _ := ret[0].(reporting.ReportingServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComplianceReportingServiceClient indicates an expected call of ComplianceReportingServiceClient
func (mr *MockClientsFactoryMockRecorder) ComplianceReportingServiceClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComplianceReportingServiceClient", reflect.TypeOf((*MockClientsFactory)(nil).ComplianceReportingServiceClient))
}

// ComplianceProfilesServiceClient mocks base method
func (m *MockClientsFactory) ComplianceProfilesServiceClient() (profiles.ProfilesServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComplianceProfilesServiceClient")
	ret0, _ := ret[0].(profiles.ProfilesServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComplianceProfilesServiceClient indicates an expected call of ComplianceProfilesServiceClient
func (mr *MockClientsFactoryMockRecorder) ComplianceProfilesServiceClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComplianceProfilesServiceClient", reflect.TypeOf((*MockClientsFactory)(nil).ComplianceProfilesServiceClient))
}

// ComplianceJobsServiceClient mocks base method
func (m *MockClientsFactory) ComplianceJobsServiceClient() (jobs.JobsServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComplianceJobsServiceClient")
	ret0, _ := ret[0].(jobs.JobsServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComplianceJobsServiceClient indicates an expected call of ComplianceJobsServiceClient
func (mr *MockClientsFactoryMockRecorder) ComplianceJobsServiceClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComplianceJobsServiceClient", reflect.TypeOf((*MockClientsFactory)(nil).ComplianceJobsServiceClient))
}

// ComplianceStatsServiceClient mocks base method
func (m *MockClientsFactory) ComplianceStatsServiceClient() (stats.StatsServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComplianceStatsServiceClient")
	ret0, _ := ret[0].(stats.StatsServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComplianceStatsServiceClient indicates an expected call of ComplianceStatsServiceClient
func (mr *MockClientsFactoryMockRecorder) ComplianceStatsServiceClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComplianceStatsServiceClient", reflect.TypeOf((*MockClientsFactory)(nil).ComplianceStatsServiceClient))
}

// ComplianceVersionServiceClient mocks base method
func (m *MockClientsFactory) ComplianceVersionServiceClient() (version.VersionServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComplianceVersionServiceClient")
	ret0, _ := ret[0].(version.VersionServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComplianceVersionServiceClient indicates an expected call of ComplianceVersionServiceClient
func (mr *MockClientsFactoryMockRecorder) ComplianceVersionServiceClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComplianceVersionServiceClient", reflect.TypeOf((*MockClientsFactory)(nil).ComplianceVersionServiceClient))
}

// NodeManagerClient mocks base method
func (m *MockClientsFactory) NodeManagerClient() (manager.NodeManagerServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeManagerClient")
	ret0, _ := ret[0].(manager.NodeManagerServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodeManagerClient indicates an expected call of NodeManagerClient
func (mr *MockClientsFactoryMockRecorder) NodeManagerClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeManagerClient", reflect.TypeOf((*MockClientsFactory)(nil).NodeManagerClient))
}

// LicenseControlClient mocks base method
func (m *MockClientsFactory) LicenseControlClient() (license_control.LicenseControlClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LicenseControlClient")
	ret0, _ := ret[0].(license_control.LicenseControlClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LicenseControlClient indicates an expected call of LicenseControlClient
func (mr *MockClientsFactoryMockRecorder) LicenseControlClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LicenseControlClient", reflect.TypeOf((*MockClientsFactory)(nil).LicenseControlClient))
}

// DeploymentServiceClient mocks base method
func (m *MockClientsFactory) DeploymentServiceClient() (deployment.DeploymentClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeploymentServiceClient")
	ret0, _ := ret[0].(deployment.DeploymentClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeploymentServiceClient indicates an expected call of DeploymentServiceClient
func (mr *MockClientsFactoryMockRecorder) DeploymentServiceClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploymentServiceClient", reflect.TypeOf((*MockClientsFactory)(nil).DeploymentServiceClient))
}

// PurgeClient mocks base method
func (m *MockClientsFactory) PurgeClient(service string) (data_lifecycle.PurgeClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PurgeClient", service)
	ret0, _ := ret[0].(data_lifecycle.PurgeClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PurgeClient indicates an expected call of PurgeClient
func (mr *MockClientsFactoryMockRecorder) PurgeClient(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PurgeClient", reflect.TypeOf((*MockClientsFactory)(nil).PurgeClient), service)
}

// Close mocks base method
func (m *MockClientsFactory) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	return ret[0].(error)
}

// Close indicates an expected call of Close
func (mr *MockClientsFactoryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClientsFactory)(nil).Close))
}

// InfraProxyClient mocks base method
func (m *MockClientsFactory) InfraProxyClient() (service0.InfraProxyClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InfraProxyClient")
	ret0, _ := ret[0].(service0.InfraProxyClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InfraProxyClient indicates an expected call of InfraProxyClient
func (mr *MockClientsFactoryMockRecorder) InfraProxyClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InfraProxyClient", reflect.TypeOf((*MockClientsFactory)(nil).InfraProxyClient))
}
