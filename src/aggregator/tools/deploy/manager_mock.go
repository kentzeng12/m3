// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/aggregator/tools/deploy/manager.go

// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package deploy is a generated GoMock package.
package deploy

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockInstance is a mock of Instance interface.
type MockInstance struct {
	ctrl     *gomock.Controller
	recorder *MockInstanceMockRecorder
}

// MockInstanceMockRecorder is the mock recorder for MockInstance.
type MockInstanceMockRecorder struct {
	mock *MockInstance
}

// NewMockInstance creates a new mock instance.
func NewMockInstance(ctrl *gomock.Controller) *MockInstance {
	mock := &MockInstance{ctrl: ctrl}
	mock.recorder = &MockInstanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstance) EXPECT() *MockInstanceMockRecorder {
	return m.recorder
}

// ID mocks base method.
func (m *MockInstance) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockInstanceMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockInstance)(nil).ID))
}

// IsDeploying mocks base method.
func (m *MockInstance) IsDeploying() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDeploying")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDeploying indicates an expected call of IsDeploying.
func (mr *MockInstanceMockRecorder) IsDeploying() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDeploying", reflect.TypeOf((*MockInstance)(nil).IsDeploying))
}

// IsHealthy mocks base method.
func (m *MockInstance) IsHealthy() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsHealthy")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsHealthy indicates an expected call of IsHealthy.
func (mr *MockInstanceMockRecorder) IsHealthy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsHealthy", reflect.TypeOf((*MockInstance)(nil).IsHealthy))
}

// Revision mocks base method.
func (m *MockInstance) Revision() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revision")
	ret0, _ := ret[0].(string)
	return ret0
}

// Revision indicates an expected call of Revision.
func (mr *MockInstanceMockRecorder) Revision() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revision", reflect.TypeOf((*MockInstance)(nil).Revision))
}

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Deploy mocks base method.
func (m *MockManager) Deploy(instanceIDs []string, revision string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", instanceIDs, revision)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deploy indicates an expected call of Deploy.
func (mr *MockManagerMockRecorder) Deploy(instanceIDs, revision interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockManager)(nil).Deploy), instanceIDs, revision)
}

// Query mocks base method.
func (m *MockManager) Query(instanceIDs []string) ([]Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", instanceIDs)
	ret0, _ := ret[0].([]Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockManagerMockRecorder) Query(instanceIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockManager)(nil).Query), instanceIDs)
}

// QueryAll mocks base method.
func (m *MockManager) QueryAll() ([]Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryAll")
	ret0, _ := ret[0].([]Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryAll indicates an expected call of QueryAll.
func (mr *MockManagerMockRecorder) QueryAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryAll", reflect.TypeOf((*MockManager)(nil).QueryAll))
}
