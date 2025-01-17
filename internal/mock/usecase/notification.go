// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/domain/usecase/notification.go
//
// Generated by this command:
//
//	mockgen -source=./internal/domain/usecase/notification.go -destination=./internal/mock/usecase/notification.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	usecase "parse_server/internal/domain/usecase"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockNotification is a mock of Notification interface.
type MockNotification struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationMockRecorder
}

// MockNotificationMockRecorder is the mock recorder for MockNotification.
type MockNotificationMockRecorder struct {
	mock *MockNotification
}

// NewMockNotification creates a new mock instance.
func NewMockNotification(ctrl *gomock.Controller) *MockNotification {
	mock := &MockNotification{ctrl: ctrl}
	mock.recorder = &MockNotificationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotification) EXPECT() *MockNotificationMockRecorder {
	return m.recorder
}

// Notify mocks base method.
func (m *MockNotification) Notify(address string, tx usecase.Transaction) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Notify", address, tx)
}

// Notify indicates an expected call of Notify.
func (mr *MockNotificationMockRecorder) Notify(address, tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockNotification)(nil).Notify), address, tx)
}
