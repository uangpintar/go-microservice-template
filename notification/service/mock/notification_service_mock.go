// Code generated by MockGen. DO NOT EDIT.
// Source: service/notification_service.go

// Package mock_svc is a generated GoMock package.
package mock_svc

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNotificationSvc is a mock of NotificationSvc interface.
type MockNotificationSvc struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationSvcMockRecorder
}

// MockNotificationSvcMockRecorder is the mock recorder for MockNotificationSvc.
type MockNotificationSvcMockRecorder struct {
	mock *MockNotificationSvc
}

// NewMockNotificationSvc creates a new mock instance.
func NewMockNotificationSvc(ctrl *gomock.Controller) *MockNotificationSvc {
	mock := &MockNotificationSvc{ctrl: ctrl}
	mock.recorder = &MockNotificationSvcMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationSvc) EXPECT() *MockNotificationSvcMockRecorder {
	return m.recorder
}

// ListenAndSendEmail mocks base method.
func (m *MockNotificationSvc) ListenAndSendEmail(ctx context.Context, isEndlessly bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ListenAndSendEmail", ctx, isEndlessly)
}

// ListenAndSendEmail indicates an expected call of ListenAndSendEmail.
func (mr *MockNotificationSvcMockRecorder) ListenAndSendEmail(ctx, isEndlessly interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenAndSendEmail", reflect.TypeOf((*MockNotificationSvc)(nil).ListenAndSendEmail), ctx, isEndlessly)
}