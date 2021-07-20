// Code generated by MockGen. DO NOT EDIT.
// Source: code.cloudfoundry.org/cli/plugin (interfaces: CliConnection)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	plugin_models "code.cloudfoundry.org/cli/plugin/models"
	gomock "github.com/golang/mock/gomock"
)

// MockCliConnection is a mock of CliConnection interface.
type MockCliConnection struct {
	ctrl     *gomock.Controller
	recorder *MockCliConnectionMockRecorder
}

// MockCliConnectionMockRecorder is the mock recorder for MockCliConnection.
type MockCliConnectionMockRecorder struct {
	mock *MockCliConnection
}

// NewMockCliConnection creates a new mock instance.
func NewMockCliConnection(ctrl *gomock.Controller) *MockCliConnection {
	mock := &MockCliConnection{ctrl: ctrl}
	mock.recorder = &MockCliConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCliConnection) EXPECT() *MockCliConnectionMockRecorder {
	return m.recorder
}

// AccessToken mocks base method.
func (m *MockCliConnection) AccessToken() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessToken")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccessToken indicates an expected call of AccessToken.
func (mr *MockCliConnectionMockRecorder) AccessToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessToken", reflect.TypeOf((*MockCliConnection)(nil).AccessToken))
}

// ApiEndpoint mocks base method.
func (m *MockCliConnection) ApiEndpoint() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApiEndpoint")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApiEndpoint indicates an expected call of ApiEndpoint.
func (mr *MockCliConnectionMockRecorder) ApiEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApiEndpoint", reflect.TypeOf((*MockCliConnection)(nil).ApiEndpoint))
}

// ApiVersion mocks base method.
func (m *MockCliConnection) ApiVersion() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApiVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApiVersion indicates an expected call of ApiVersion.
func (mr *MockCliConnectionMockRecorder) ApiVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApiVersion", reflect.TypeOf((*MockCliConnection)(nil).ApiVersion))
}

// CliCommand mocks base method.
func (m *MockCliConnection) CliCommand(arg0 ...string) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CliCommand", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CliCommand indicates an expected call of CliCommand.
func (mr *MockCliConnectionMockRecorder) CliCommand(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CliCommand", reflect.TypeOf((*MockCliConnection)(nil).CliCommand), arg0...)
}

// CliCommandWithoutTerminalOutput mocks base method.
func (m *MockCliConnection) CliCommandWithoutTerminalOutput(arg0 ...string) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CliCommandWithoutTerminalOutput", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CliCommandWithoutTerminalOutput indicates an expected call of CliCommandWithoutTerminalOutput.
func (mr *MockCliConnectionMockRecorder) CliCommandWithoutTerminalOutput(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CliCommandWithoutTerminalOutput", reflect.TypeOf((*MockCliConnection)(nil).CliCommandWithoutTerminalOutput), arg0...)
}

// DopplerEndpoint mocks base method.
func (m *MockCliConnection) DopplerEndpoint() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DopplerEndpoint")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DopplerEndpoint indicates an expected call of DopplerEndpoint.
func (mr *MockCliConnectionMockRecorder) DopplerEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DopplerEndpoint", reflect.TypeOf((*MockCliConnection)(nil).DopplerEndpoint))
}

// GetApp mocks base method.
func (m *MockCliConnection) GetApp(arg0 string) (plugin_models.GetAppModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApp", arg0)
	ret0, _ := ret[0].(plugin_models.GetAppModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApp indicates an expected call of GetApp.
func (mr *MockCliConnectionMockRecorder) GetApp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApp", reflect.TypeOf((*MockCliConnection)(nil).GetApp), arg0)
}

// GetApps mocks base method.
func (m *MockCliConnection) GetApps() ([]plugin_models.GetAppsModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApps")
	ret0, _ := ret[0].([]plugin_models.GetAppsModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApps indicates an expected call of GetApps.
func (mr *MockCliConnectionMockRecorder) GetApps() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApps", reflect.TypeOf((*MockCliConnection)(nil).GetApps))
}

// GetCurrentOrg mocks base method.
func (m *MockCliConnection) GetCurrentOrg() (plugin_models.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentOrg")
	ret0, _ := ret[0].(plugin_models.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentOrg indicates an expected call of GetCurrentOrg.
func (mr *MockCliConnectionMockRecorder) GetCurrentOrg() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentOrg", reflect.TypeOf((*MockCliConnection)(nil).GetCurrentOrg))
}

// GetCurrentSpace mocks base method.
func (m *MockCliConnection) GetCurrentSpace() (plugin_models.Space, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentSpace")
	ret0, _ := ret[0].(plugin_models.Space)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentSpace indicates an expected call of GetCurrentSpace.
func (mr *MockCliConnectionMockRecorder) GetCurrentSpace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentSpace", reflect.TypeOf((*MockCliConnection)(nil).GetCurrentSpace))
}

// GetOrg mocks base method.
func (m *MockCliConnection) GetOrg(arg0 string) (plugin_models.GetOrg_Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrg", arg0)
	ret0, _ := ret[0].(plugin_models.GetOrg_Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrg indicates an expected call of GetOrg.
func (mr *MockCliConnectionMockRecorder) GetOrg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrg", reflect.TypeOf((*MockCliConnection)(nil).GetOrg), arg0)
}

// GetOrgUsers mocks base method.
func (m *MockCliConnection) GetOrgUsers(arg0 string, arg1 ...string) ([]plugin_models.GetOrgUsers_Model, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrgUsers", varargs...)
	ret0, _ := ret[0].([]plugin_models.GetOrgUsers_Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgUsers indicates an expected call of GetOrgUsers.
func (mr *MockCliConnectionMockRecorder) GetOrgUsers(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgUsers", reflect.TypeOf((*MockCliConnection)(nil).GetOrgUsers), varargs...)
}

// GetOrgs mocks base method.
func (m *MockCliConnection) GetOrgs() ([]plugin_models.GetOrgs_Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgs")
	ret0, _ := ret[0].([]plugin_models.GetOrgs_Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgs indicates an expected call of GetOrgs.
func (mr *MockCliConnectionMockRecorder) GetOrgs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgs", reflect.TypeOf((*MockCliConnection)(nil).GetOrgs))
}

// GetService mocks base method.
func (m *MockCliConnection) GetService(arg0 string) (plugin_models.GetService_Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", arg0)
	ret0, _ := ret[0].(plugin_models.GetService_Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService.
func (mr *MockCliConnectionMockRecorder) GetService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockCliConnection)(nil).GetService), arg0)
}

// GetServices mocks base method.
func (m *MockCliConnection) GetServices() ([]plugin_models.GetServices_Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServices")
	ret0, _ := ret[0].([]plugin_models.GetServices_Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServices indicates an expected call of GetServices.
func (mr *MockCliConnectionMockRecorder) GetServices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServices", reflect.TypeOf((*MockCliConnection)(nil).GetServices))
}

// GetSpace mocks base method.
func (m *MockCliConnection) GetSpace(arg0 string) (plugin_models.GetSpace_Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpace", arg0)
	ret0, _ := ret[0].(plugin_models.GetSpace_Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpace indicates an expected call of GetSpace.
func (mr *MockCliConnectionMockRecorder) GetSpace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpace", reflect.TypeOf((*MockCliConnection)(nil).GetSpace), arg0)
}

// GetSpaceUsers mocks base method.
func (m *MockCliConnection) GetSpaceUsers(arg0, arg1 string) ([]plugin_models.GetSpaceUsers_Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpaceUsers", arg0, arg1)
	ret0, _ := ret[0].([]plugin_models.GetSpaceUsers_Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpaceUsers indicates an expected call of GetSpaceUsers.
func (mr *MockCliConnectionMockRecorder) GetSpaceUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpaceUsers", reflect.TypeOf((*MockCliConnection)(nil).GetSpaceUsers), arg0, arg1)
}

// GetSpaces mocks base method.
func (m *MockCliConnection) GetSpaces() ([]plugin_models.GetSpaces_Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpaces")
	ret0, _ := ret[0].([]plugin_models.GetSpaces_Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpaces indicates an expected call of GetSpaces.
func (mr *MockCliConnectionMockRecorder) GetSpaces() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpaces", reflect.TypeOf((*MockCliConnection)(nil).GetSpaces))
}

// HasAPIEndpoint mocks base method.
func (m *MockCliConnection) HasAPIEndpoint() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasAPIEndpoint")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasAPIEndpoint indicates an expected call of HasAPIEndpoint.
func (mr *MockCliConnectionMockRecorder) HasAPIEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasAPIEndpoint", reflect.TypeOf((*MockCliConnection)(nil).HasAPIEndpoint))
}

// HasOrganization mocks base method.
func (m *MockCliConnection) HasOrganization() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasOrganization")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasOrganization indicates an expected call of HasOrganization.
func (mr *MockCliConnectionMockRecorder) HasOrganization() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasOrganization", reflect.TypeOf((*MockCliConnection)(nil).HasOrganization))
}

// HasSpace mocks base method.
func (m *MockCliConnection) HasSpace() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasSpace")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasSpace indicates an expected call of HasSpace.
func (mr *MockCliConnectionMockRecorder) HasSpace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasSpace", reflect.TypeOf((*MockCliConnection)(nil).HasSpace))
}

// IsLoggedIn mocks base method.
func (m *MockCliConnection) IsLoggedIn() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLoggedIn")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsLoggedIn indicates an expected call of IsLoggedIn.
func (mr *MockCliConnectionMockRecorder) IsLoggedIn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLoggedIn", reflect.TypeOf((*MockCliConnection)(nil).IsLoggedIn))
}

// IsSSLDisabled mocks base method.
func (m *MockCliConnection) IsSSLDisabled() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSSLDisabled")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSSLDisabled indicates an expected call of IsSSLDisabled.
func (mr *MockCliConnectionMockRecorder) IsSSLDisabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSSLDisabled", reflect.TypeOf((*MockCliConnection)(nil).IsSSLDisabled))
}

// LoggregatorEndpoint mocks base method.
func (m *MockCliConnection) LoggregatorEndpoint() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoggregatorEndpoint")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoggregatorEndpoint indicates an expected call of LoggregatorEndpoint.
func (mr *MockCliConnectionMockRecorder) LoggregatorEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoggregatorEndpoint", reflect.TypeOf((*MockCliConnection)(nil).LoggregatorEndpoint))
}

// UserEmail mocks base method.
func (m *MockCliConnection) UserEmail() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserEmail")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserEmail indicates an expected call of UserEmail.
func (mr *MockCliConnectionMockRecorder) UserEmail() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserEmail", reflect.TypeOf((*MockCliConnection)(nil).UserEmail))
}

// UserGuid mocks base method.
func (m *MockCliConnection) UserGuid() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserGuid")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserGuid indicates an expected call of UserGuid.
func (mr *MockCliConnectionMockRecorder) UserGuid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserGuid", reflect.TypeOf((*MockCliConnection)(nil).UserGuid))
}

// Username mocks base method.
func (m *MockCliConnection) Username() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Username")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Username indicates an expected call of Username.
func (mr *MockCliConnectionMockRecorder) Username() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Username", reflect.TypeOf((*MockCliConnection)(nil).Username))
}
