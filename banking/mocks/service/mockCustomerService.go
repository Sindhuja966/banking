package service

import "github.com/golang/mock/gomock"

type MockCustomerService struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerServiceMockRecorder
}

type MockCustomerServiceMockRecorder struct {
	mock *MockCustomerService
}
