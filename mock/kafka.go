// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lovoo/goka/kafka (interfaces: Consumer,TopicManager,Producer)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	kafka "github.com/lovoo/goka/kafka"
	reflect "reflect"
)

// MockConsumer is a mock of Consumer interface
type MockConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockConsumerMockRecorder
}

// MockConsumerMockRecorder is the mock recorder for MockConsumer
type MockConsumerMockRecorder struct {
	mock *MockConsumer
}

// NewMockConsumer creates a new mock instance
func NewMockConsumer(ctrl *gomock.Controller) *MockConsumer {
	mock := &MockConsumer{ctrl: ctrl}
	mock.recorder = &MockConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConsumer) EXPECT() *MockConsumerMockRecorder {
	return m.recorder
}

// AddGroupPartition mocks base method
func (m *MockConsumer) AddGroupPartition(arg0 int32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddGroupPartition", arg0)
}

// AddGroupPartition indicates an expected call of AddGroupPartition
func (mr *MockConsumerMockRecorder) AddGroupPartition(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddGroupPartition", reflect.TypeOf((*MockConsumer)(nil).AddGroupPartition), arg0)
}

// AddPartition mocks base method
func (m *MockConsumer) AddPartition(arg0 string, arg1 int32, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPartition", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPartition indicates an expected call of AddPartition
func (mr *MockConsumerMockRecorder) AddPartition(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPartition", reflect.TypeOf((*MockConsumer)(nil).AddPartition), arg0, arg1, arg2)
}

// Close mocks base method
func (m *MockConsumer) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockConsumerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConsumer)(nil).Close))
}

// Commit mocks base method
func (m *MockConsumer) Commit(arg0 string, arg1 int32, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockConsumerMockRecorder) Commit(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockConsumer)(nil).Commit), arg0, arg1, arg2)
}

// Events mocks base method
func (m *MockConsumer) Events() <-chan kafka.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Events")
	ret0, _ := ret[0].(<-chan kafka.Event)
	return ret0
}

// Events indicates an expected call of Events
func (mr *MockConsumerMockRecorder) Events() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Events", reflect.TypeOf((*MockConsumer)(nil).Events))
}

// RemovePartition mocks base method
func (m *MockConsumer) RemovePartition(arg0 string, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePartition", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePartition indicates an expected call of RemovePartition
func (mr *MockConsumerMockRecorder) RemovePartition(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePartition", reflect.TypeOf((*MockConsumer)(nil).RemovePartition), arg0, arg1)
}

// Subscribe mocks base method
func (m *MockConsumer) Subscribe(arg0 map[string]int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockConsumerMockRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockConsumer)(nil).Subscribe), arg0)
}

// MockTopicManager is a mock of TopicManager interface
type MockTopicManager struct {
	ctrl     *gomock.Controller
	recorder *MockTopicManagerMockRecorder
}

// MockTopicManagerMockRecorder is the mock recorder for MockTopicManager
type MockTopicManagerMockRecorder struct {
	mock *MockTopicManager
}

// NewMockTopicManager creates a new mock instance
func NewMockTopicManager(ctrl *gomock.Controller) *MockTopicManager {
	mock := &MockTopicManager{ctrl: ctrl}
	mock.recorder = &MockTopicManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTopicManager) EXPECT() *MockTopicManagerMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockTopicManager) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockTopicManagerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockTopicManager)(nil).Close))
}

// EnsureStreamExists mocks base method
func (m *MockTopicManager) EnsureStreamExists(arg0 string, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureStreamExists", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureStreamExists indicates an expected call of EnsureStreamExists
func (mr *MockTopicManagerMockRecorder) EnsureStreamExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureStreamExists", reflect.TypeOf((*MockTopicManager)(nil).EnsureStreamExists), arg0, arg1)
}

// EnsureTableExists mocks base method
func (m *MockTopicManager) EnsureTableExists(arg0 string, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureTableExists", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureTableExists indicates an expected call of EnsureTableExists
func (mr *MockTopicManagerMockRecorder) EnsureTableExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureTableExists", reflect.TypeOf((*MockTopicManager)(nil).EnsureTableExists), arg0, arg1)
}

// EnsureTopicExists mocks base method
func (m *MockTopicManager) EnsureTopicExists(arg0 string, arg1, arg2 int, arg3 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureTopicExists", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureTopicExists indicates an expected call of EnsureTopicExists
func (mr *MockTopicManagerMockRecorder) EnsureTopicExists(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureTopicExists", reflect.TypeOf((*MockTopicManager)(nil).EnsureTopicExists), arg0, arg1, arg2, arg3)
}

// Partitions mocks base method
func (m *MockTopicManager) Partitions(arg0 string) ([]int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Partitions", arg0)
	ret0, _ := ret[0].([]int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Partitions indicates an expected call of Partitions
func (mr *MockTopicManagerMockRecorder) Partitions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Partitions", reflect.TypeOf((*MockTopicManager)(nil).Partitions), arg0)
}

// MockProducer is a mock of Producer interface
type MockProducer struct {
	ctrl     *gomock.Controller
	recorder *MockProducerMockRecorder
}

// MockProducerMockRecorder is the mock recorder for MockProducer
type MockProducerMockRecorder struct {
	mock *MockProducer
}

// NewMockProducer creates a new mock instance
func NewMockProducer(ctrl *gomock.Controller) *MockProducer {
	mock := &MockProducer{ctrl: ctrl}
	mock.recorder = &MockProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProducer) EXPECT() *MockProducerMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockProducer) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockProducerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockProducer)(nil).Close))
}

// Emit mocks base method
func (m *MockProducer) Emit(arg0, arg1 string, arg2 []byte) *kafka.Promise {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Emit", arg0, arg1, arg2)
	ret0, _ := ret[0].(*kafka.Promise)
	return ret0
}

// Emit indicates an expected call of Emit
func (mr *MockProducerMockRecorder) Emit(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Emit", reflect.TypeOf((*MockProducer)(nil).Emit), arg0, arg1, arg2)
}