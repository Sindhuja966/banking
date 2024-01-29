package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sindhuja966/banking/dto"
	"github.com/Sindhuja966/banking/errs"
	"github.com/Sindhuja966/banking/mocks/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func Test_customers_status_code_200(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dummyCustomers := []dto.CustomerResponse{
		{Id: "105", Name: "Ashish", City: "New Delhi", Zipcode: "110011", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "106", Name: "Rob", City: "New Delhi", Zipcode: "110011", DateofBirth: "2000-01-01", Status: "1"},
	}
	mockService := service.NewMockCustomerService(ctrl)
	mockService.EXPECT().GetAllCustomer("").Return(dummyCustomers, nil)
	ch := CustomerHandlers{mockService}

	router := gin.Default()
	router.GET("/customers", ch.GetAllCustomers)

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_status_code_500_with_error_message(t *testing.T) {
	// Arrange

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockCustomerService(ctrl)
	mockService.EXPECT().GetAllCustomer("").Return(nil, errs.NewUnexpectedError("some database error"))

	ch := CustomerHandlers{mockService}
	// Act
	router := gin.Default()
	router.GET("/customers", ch.GetAllCustomers)

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code == http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}

}

func Test_bycustomerid_status_code_200(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dummyCustomers := []dto.CustomerResponse{
		{Id: "105", Name: "Ashish", City: "New Delhi", Zipcode: "110011", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "106", Name: "Rob", City: "New Delhi", Zipcode: "110011", DateofBirth: "2000-01-01", Status: "1"},
	}
	mockService := service.NewMockCustomerService(ctrl)
	mockService.EXPECT().GetCustomer("").Return(dummyCustomers, nil)
	ch := CustomerHandlers{mockService}

	router := gin.Default()
	router.GET("/customers/105", ch.GetCustomer)

	request, _ := http.NewRequest(http.MethodGet, "/customers/105", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}
