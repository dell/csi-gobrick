/*
 Copyright © 2020-2022 Dell Inc. or its subsidiaries. All Rights Reserved.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
      http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/
// Code generated by MockGen. DO NOT EDIT.
// Source: wrappers.go

// Package wrappers is a generated GoMock package.
package wrappers

import (
	context "context"
	os "os"
	reflect "reflect"

	goiscsi "github.com/dell/goiscsi"
	gonvme "github.com/dell/gonvme"
	gomock "github.com/golang/mock/gomock"
)

// MockLimitedFileInfo is a mock of LimitedFileInfo interface
type MockLimitedFileInfo struct {
	ctrl     *gomock.Controller
	recorder *MockLimitedFileInfoMockRecorder
}

// MockLimitedFileInfoMockRecorder is the mock recorder for MockLimitedFileInfo
type MockLimitedFileInfoMockRecorder struct {
	mock *MockLimitedFileInfo
}

// NewMockLimitedFileInfo creates a new mock instance
func NewMockLimitedFileInfo(ctrl *gomock.Controller) *MockLimitedFileInfo {
	mock := &MockLimitedFileInfo{ctrl: ctrl}
	mock.recorder = &MockLimitedFileInfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLimitedFileInfo) EXPECT() *MockLimitedFileInfoMockRecorder {
	return m.recorder
}

// IsDir mocks base method
func (m *MockLimitedFileInfo) IsDir() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDir")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDir indicates an expected call of IsDir
func (mr *MockLimitedFileInfoMockRecorder) IsDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDir", reflect.TypeOf((*MockLimitedFileInfo)(nil).IsDir))
}

// MockLimitedFile is a mock of LimitedFile interface
type MockLimitedFile struct {
	ctrl     *gomock.Controller
	recorder *MockLimitedFileMockRecorder
}

// MockLimitedFileMockRecorder is the mock recorder for MockLimitedFile
type MockLimitedFileMockRecorder struct {
	mock *MockLimitedFile
}

// NewMockLimitedFile creates a new mock instance
func NewMockLimitedFile(ctrl *gomock.Controller) *MockLimitedFile {
	mock := &MockLimitedFile{ctrl: ctrl}
	mock.recorder = &MockLimitedFileMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLimitedFile) EXPECT() *MockLimitedFileMockRecorder {
	return m.recorder
}

// WriteString mocks base method
func (m *MockLimitedFile) WriteString(s string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteString", s)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteString indicates an expected call of WriteString
func (mr *MockLimitedFileMockRecorder) WriteString(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteString", reflect.TypeOf((*MockLimitedFile)(nil).WriteString), s)
}

// Close mocks base method
func (m *MockLimitedFile) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockLimitedFileMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockLimitedFile)(nil).Close))
}

// MockLimitedOSExec is a mock of LimitedOSExec interface
type MockLimitedOSExec struct {
	ctrl     *gomock.Controller
	recorder *MockLimitedOSExecMockRecorder
}

// MockLimitedOSExecMockRecorder is the mock recorder for MockLimitedOSExec
type MockLimitedOSExecMockRecorder struct {
	mock *MockLimitedOSExec
}

// NewMockLimitedOSExec creates a new mock instance
func NewMockLimitedOSExec(ctrl *gomock.Controller) *MockLimitedOSExec {
	mock := &MockLimitedOSExec{ctrl: ctrl}
	mock.recorder = &MockLimitedOSExecMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLimitedOSExec) EXPECT() *MockLimitedOSExecMockRecorder {
	return m.recorder
}

// CommandContext mocks base method
func (m *MockLimitedOSExec) CommandContext(ctx context.Context, name string, arg ...string) LimitedOSExecCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, name}
	for _, a := range arg {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CommandContext", varargs...)
	ret0, _ := ret[0].(LimitedOSExecCmd)
	return ret0
}

// CommandContext indicates an expected call of CommandContext
func (mr *MockLimitedOSExecMockRecorder) CommandContext(ctx, name interface{}, arg ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, name}, arg...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandContext", reflect.TypeOf((*MockLimitedOSExec)(nil).CommandContext), varargs...)
}

// MockLimitedOSExecCmd is a mock of LimitedOSExecCmd interface
type MockLimitedOSExecCmd struct {
	ctrl     *gomock.Controller
	recorder *MockLimitedOSExecCmdMockRecorder
}

// MockLimitedOSExecCmdMockRecorder is the mock recorder for MockLimitedOSExecCmd
type MockLimitedOSExecCmdMockRecorder struct {
	mock *MockLimitedOSExecCmd
}

// NewMockLimitedOSExecCmd creates a new mock instance
func NewMockLimitedOSExecCmd(ctrl *gomock.Controller) *MockLimitedOSExecCmd {
	mock := &MockLimitedOSExecCmd{ctrl: ctrl}
	mock.recorder = &MockLimitedOSExecCmdMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLimitedOSExecCmd) EXPECT() *MockLimitedOSExecCmdMockRecorder {
	return m.recorder
}

// CombinedOutput mocks base method
func (m *MockLimitedOSExecCmd) CombinedOutput() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CombinedOutput")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CombinedOutput indicates an expected call of CombinedOutput
func (mr *MockLimitedOSExecCmdMockRecorder) CombinedOutput() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CombinedOutput", reflect.TypeOf((*MockLimitedOSExecCmd)(nil).CombinedOutput))
}

// MockLimitedIOUtil is a mock of LimitedIOUtil interface
type MockLimitedIOUtil struct {
	ctrl     *gomock.Controller
	recorder *MockLimitedIOUtilMockRecorder
}

// MockLimitedIOUtilMockRecorder is the mock recorder for MockLimitedIOUtil
type MockLimitedIOUtilMockRecorder struct {
	mock *MockLimitedIOUtil
}

// NewMockLimitedIOUtil creates a new mock instance
func NewMockLimitedIOUtil(ctrl *gomock.Controller) *MockLimitedIOUtil {
	mock := &MockLimitedIOUtil{ctrl: ctrl}
	mock.recorder = &MockLimitedIOUtilMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLimitedIOUtil) EXPECT() *MockLimitedIOUtilMockRecorder {
	return m.recorder
}

// ReadFile mocks base method
func (m *MockLimitedOS) ReadFile(filename string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFile", filename)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile
func (mr *MockLimitedOSMockRecorder) ReadFile(filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockLimitedOS)(nil).ReadFile), filename)
}

// MockLimitedFilepath is a mock of LimitedFilepath interface
type MockLimitedFilepath struct {
	ctrl     *gomock.Controller
	recorder *MockLimitedFilepathMockRecorder
}

// MockLimitedFilepathMockRecorder is the mock recorder for MockLimitedFilepath
type MockLimitedFilepathMockRecorder struct {
	mock *MockLimitedFilepath
}

// NewMockLimitedFilepath creates a new mock instance
func NewMockLimitedFilepath(ctrl *gomock.Controller) *MockLimitedFilepath {
	mock := &MockLimitedFilepath{ctrl: ctrl}
	mock.recorder = &MockLimitedFilepathMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLimitedFilepath) EXPECT() *MockLimitedFilepathMockRecorder {
	return m.recorder
}

// Glob mocks base method
func (m *MockLimitedFilepath) Glob(pattern string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Glob", pattern)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Glob indicates an expected call of Glob
func (mr *MockLimitedFilepathMockRecorder) Glob(pattern interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Glob", reflect.TypeOf((*MockLimitedFilepath)(nil).Glob), pattern)
}

// EvalSymlinks mocks base method
func (m *MockLimitedFilepath) EvalSymlinks(path string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EvalSymlinks", path)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EvalSymlinks indicates an expected call of EvalSymlinks
func (mr *MockLimitedFilepathMockRecorder) EvalSymlinks(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvalSymlinks", reflect.TypeOf((*MockLimitedFilepath)(nil).EvalSymlinks), path)
}

// MockLimitedOS is a mock of LimitedOS interface
type MockLimitedOS struct {
	ctrl     *gomock.Controller
	recorder *MockLimitedOSMockRecorder
}

// MockLimitedOSMockRecorder is the mock recorder for MockLimitedOS
type MockLimitedOSMockRecorder struct {
	mock *MockLimitedOS
}

// NewMockLimitedOS creates a new mock instance
func NewMockLimitedOS(ctrl *gomock.Controller) *MockLimitedOS {
	mock := &MockLimitedOS{ctrl: ctrl}
	mock.recorder = &MockLimitedOSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLimitedOS) EXPECT() *MockLimitedOSMockRecorder {
	return m.recorder
}

// OpenFile mocks base method
func (m *MockLimitedOS) OpenFile(name string, flag int, perm os.FileMode) (LimitedFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenFile", name, flag, perm)
	ret0, _ := ret[0].(LimitedFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenFile indicates an expected call of OpenFile
func (mr *MockLimitedOSMockRecorder) OpenFile(name, flag, perm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenFile", reflect.TypeOf((*MockLimitedOS)(nil).OpenFile), name, flag, perm)
}

// Stat mocks base method
func (m *MockLimitedOS) Stat(name string) (LimitedFileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stat", name)
	ret0, _ := ret[0].(LimitedFileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stat indicates an expected call of Stat
func (mr *MockLimitedOSMockRecorder) Stat(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockLimitedOS)(nil).Stat), name)
}

// IsNotExist mocks base method
func (m *MockLimitedOS) IsNotExist(err error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNotExist", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNotExist indicates an expected call of IsNotExist
func (mr *MockLimitedOSMockRecorder) IsNotExist(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNotExist", reflect.TypeOf((*MockLimitedOS)(nil).IsNotExist), err)
}

// Mkdir mocks base method
func (m *MockLimitedOS) Mkdir(name string, perm os.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mkdir", name, perm)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mkdir indicates an expected call of Mkdir
func (mr *MockLimitedOSMockRecorder) Mkdir(name, perm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mkdir", reflect.TypeOf((*MockLimitedOS)(nil).Mkdir), name, perm)
}

// Remove mocks base method
func (m *MockLimitedOS) Remove(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockLimitedOSMockRecorder) Remove(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockLimitedOS)(nil).Remove), name)
}

// MockISCSILib is a mock of ISCSILib interface
type MockISCSILib struct {
	ctrl     *gomock.Controller
	recorder *MockISCSILibMockRecorder
}

// MockISCSILibMockRecorder is the mock recorder for MockISCSILib
type MockISCSILibMockRecorder struct {
	mock *MockISCSILib
}

// NewMockISCSILib creates a new mock instance
func NewMockISCSILib(ctrl *gomock.Controller) *MockISCSILib {
	mock := &MockISCSILib{ctrl: ctrl}
	mock.recorder = &MockISCSILibMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockISCSILib) EXPECT() *MockISCSILibMockRecorder {
	return m.recorder
}

// DiscoverTargets mocks base method
func (m *MockISCSILib) DiscoverTargets(address string, login bool) ([]goiscsi.ISCSITarget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiscoverTargets", address, login)
	ret0, _ := ret[0].([]goiscsi.ISCSITarget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiscoverTargets indicates an expected call of DiscoverTargets
func (mr *MockISCSILibMockRecorder) DiscoverTargets(address, login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoverTargets", reflect.TypeOf((*MockISCSILib)(nil).DiscoverTargets), address, login)
}

// GetInitiators mocks base method
func (m *MockISCSILib) GetInitiators(filename string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInitiators", filename)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInitiators indicates an expected call of GetInitiators
func (mr *MockISCSILibMockRecorder) GetInitiators(filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInitiators", reflect.TypeOf((*MockISCSILib)(nil).GetInitiators), filename)
}

// PerformLogin mocks base method
func (m *MockISCSILib) PerformLogin(target goiscsi.ISCSITarget) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PerformLogin", target)
	ret0, _ := ret[0].(error)
	return ret0
}

// PerformLogin indicates an expected call of PerformLogin
func (mr *MockISCSILibMockRecorder) PerformLogin(target interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PerformLogin", reflect.TypeOf((*MockISCSILib)(nil).PerformLogin), target)
}

// PerformLogout mocks base method
func (m *MockISCSILib) PerformLogout(target goiscsi.ISCSITarget) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PerformLogout", target)
	ret0, _ := ret[0].(error)
	return ret0
}

// PerformLogout indicates an expected call of PerformLogout
func (mr *MockISCSILibMockRecorder) PerformLogout(target interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PerformLogout", reflect.TypeOf((*MockISCSILib)(nil).PerformLogout), target)
}

// PerformRescan mocks base method
func (m *MockISCSILib) PerformRescan() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PerformRescan")
	ret0, _ := ret[0].(error)
	return ret0
}

// PerformRescan indicates an expected call of PerformRescan
func (mr *MockISCSILibMockRecorder) PerformRescan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PerformRescan", reflect.TypeOf((*MockISCSILib)(nil).PerformRescan))
}

// GetSessions mocks base method
func (m *MockISCSILib) GetSessions() ([]goiscsi.ISCSISession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessions")
	ret0, _ := ret[0].([]goiscsi.ISCSISession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessions indicates an expected call of GetSessions
func (mr *MockISCSILibMockRecorder) GetSessions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessions", reflect.TypeOf((*MockISCSILib)(nil).GetSessions))
}

// GetNodes mocks base method
func (m *MockISCSILib) GetNodes() ([]goiscsi.ISCSINode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodes")
	ret0, _ := ret[0].([]goiscsi.ISCSINode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodes indicates an expected call of GetNodes
func (mr *MockISCSILibMockRecorder) GetNodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodes", reflect.TypeOf((*MockISCSILib)(nil).GetNodes))
}

// CreateOrUpdateNode mocks base method
func (m *MockISCSILib) CreateOrUpdateNode(target goiscsi.ISCSITarget, options map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateNode", target, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateNode indicates an expected call of CreateOrUpdateNode
func (mr *MockISCSILibMockRecorder) CreateOrUpdateNode(target, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateNode", reflect.TypeOf((*MockISCSILib)(nil).CreateOrUpdateNode), target, options)
}

// DeleteNode mocks base method
func (m *MockISCSILib) DeleteNode(target goiscsi.ISCSITarget) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNode", target)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNode indicates an expected call of DeleteNode
func (mr *MockISCSILibMockRecorder) DeleteNode(target interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNode", reflect.TypeOf((*MockISCSILib)(nil).DeleteNode), target)
}

// MockNVMe is a mock of NVMe interface.
type MockNVMe struct {
	ctrl     *gomock.Controller
	recorder *MockNVMeMockRecorder
}

// MockNVMeMockRecorder is the mock recorder for MockNVMe.
type MockNVMeMockRecorder struct {
	mock *MockNVMe
}

// NewMockNVMe creates a new mock instance.
func NewMockNVMe(ctrl *gomock.Controller) *MockNVMe {
	mock := &MockNVMe{ctrl: ctrl}
	mock.recorder = &MockNVMeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNVMe) EXPECT() *MockNVMeMockRecorder {
	return m.recorder
}

// DiscoverNVMeFCTargets mocks base method.
func (m *MockNVMe) DiscoverNVMeFCTargets(address string, login bool) ([]gonvme.NVMeTarget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiscoverNVMeFCTargets", address, login)
	ret0, _ := ret[0].([]gonvme.NVMeTarget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiscoverNVMeFCTargets indicates an expected call of DiscoverNVMeFCTargets.
func (mr *MockNVMeMockRecorder) DiscoverNVMeFCTargets(address, login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoverNVMeFCTargets", reflect.TypeOf((*MockNVMe)(nil).DiscoverNVMeFCTargets), address, login)
}

// DiscoverNVMeTCPTargets mocks base method.
func (m *MockNVMe) DiscoverNVMeTCPTargets(address string, login bool) ([]gonvme.NVMeTarget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiscoverNVMeTCPTargets", address, login)
	ret0, _ := ret[0].([]gonvme.NVMeTarget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiscoverNVMeTCPTargets indicates an expected call of DiscoverNVMeTCPTargets.
func (mr *MockNVMeMockRecorder) DiscoverNVMeTCPTargets(address, login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoverNVMeTCPTargets", reflect.TypeOf((*MockNVMe)(nil).DiscoverNVMeTCPTargets), address, login)
}

// GetInitiators mocks base method.
func (m *MockNVMe) GetInitiators(filename string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInitiators", filename)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInitiators indicates an expected call of GetInitiators.
func (mr *MockNVMeMockRecorder) GetInitiators(filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInitiators", reflect.TypeOf((*MockNVMe)(nil).GetInitiators), filename)
}

// GetNVMeDeviceData mocks base method.
func (m *MockNVMe) GetNVMeDeviceData(path string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNVMeDeviceData", path)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetNVMeDeviceData indicates an expected call of GetNVMeDeviceData.
func (mr *MockNVMeMockRecorder) GetNVMeDeviceData(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNVMeDeviceData", reflect.TypeOf((*MockNVMe)(nil).GetNVMeDeviceData), path)
}

// GetSessions mocks base method.
func (m *MockNVMe) GetSessions() ([]gonvme.NVMESession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessions")
	ret0, _ := ret[0].([]gonvme.NVMESession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessions indicates an expected call of GetSessions.
func (mr *MockNVMeMockRecorder) GetSessions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessions", reflect.TypeOf((*MockNVMe)(nil).GetSessions))
}

// ListNVMeDeviceAndNamespace mocks base method.
func (m *MockNVMe) ListNVMeDeviceAndNamespace() ([]gonvme.DevicePathAndNamespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNVMeDeviceAndNamespace")
	ret0, _ := ret[0].([]gonvme.DevicePathAndNamespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNVMeDeviceAndNamespace indicates an expected call of ListNVMeDeviceAndNamespace.
func (mr *MockNVMeMockRecorder) ListNVMeDeviceAndNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNVMeDeviceAndNamespace", reflect.TypeOf((*MockNVMe)(nil).ListNVMeDeviceAndNamespace))
}

// NVMeDisconnect mocks base method.
func (m *MockNVMe) NVMeDisconnect(target gonvme.NVMeTarget) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NVMeDisconnect", target)
	ret0, _ := ret[0].(error)
	return ret0
}

// NVMeDisconnect indicates an expected call of NVMeDisconnect.
func (mr *MockNVMeMockRecorder) NVMeDisconnect(target interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NVMeDisconnect", reflect.TypeOf((*MockNVMe)(nil).NVMeDisconnect), target)
}

// NVMeFCConnect mocks base method.
func (m *MockNVMe) NVMeFCConnect(target gonvme.NVMeTarget, duplicateConnect bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NVMeFCConnect", target, duplicateConnect)
	ret0, _ := ret[0].(error)
	return ret0
}

// NVMeFCConnect indicates an expected call of NVMeFCConnect.
func (mr *MockNVMeMockRecorder) NVMeFCConnect(target, duplicateConnect interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NVMeFCConnect", reflect.TypeOf((*MockNVMe)(nil).NVMeFCConnect), target, duplicateConnect)
}

// NVMeTCPConnect mocks base method.
func (m *MockNVMe) NVMeTCPConnect(target gonvme.NVMeTarget, duplicateConnect bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NVMeTCPConnect", target, duplicateConnect)
	ret0, _ := ret[0].(error)
	return ret0
}

// NVMeTCPConnect indicates an expected call of NVMeTCPConnect.
func (mr *MockNVMeMockRecorder) NVMeTCPConnect(target, duplicateConnect interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NVMeTCPConnect", reflect.TypeOf((*MockNVMe)(nil).NVMeTCPConnect), target, duplicateConnect)
}
