package app

import (
	"net/http"
	"net/http/httptest"
	"strings"
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
		{Id: 105, Name: "Ashish", City: "New Delhi", Zipcode: "110011", DateofBirth: "2000-01-01", Status: "1"},
		{Id: 106, Name: "Rob", City: "New Delhi", Zipcode: "110011", DateofBirth: "2000-01-01", Status: "1"},
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
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}

}

func Test_getCustomer_Success(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockCustomerService(ctrl)
	ch := CustomerHandlers{service: mockService}

	router := gin.Default()
	router.GET("/customers/:customer_id", ch.GetCustomer)

	expectedCustomer := &dto.CustomerResponse{
		Id:          100,
		Name:        "Sindhu",
		City:        "HYD",
		Zipcode:     "53243",
		DateofBirth: "1999-01-22",
		Status:      "1",
	}
	mockService.EXPECT().GetCustomer(100).Return(expectedCustomer, nil)

	request, _ := http.NewRequest(http.MethodGet, "/customers/100", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_getcustomer_status_code_500_with_error_message(t *testing.T) {
	// Arrange

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockCustomerService(ctrl)
	ch := CustomerHandlers{mockService}
	mockService.EXPECT().GetCustomer(100).Return(nil, errs.NewUnexpectedError("some database error"))

	// Act
	router := gin.Default()
	router.GET("/customers/:customer_id", ch.GetCustomer)

	request, _ := http.NewRequest(http.MethodGet, "/customers/100", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}

}

func Test_postCustomer_Success(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockCustomerService(ctrl)
	ch := CustomerHandlers{service: mockService}

	router := gin.Default()
	router.POST("/postCustomer", ch.CreateCustomer)

	expectedCustomer := &dto.CustomerResponse{
		Id:          109,
		Name:        "John",
		DateofBirth: "1999-02-13",
		City:        "GNT",
		Zipcode:     "12345",
		Status:      "1",
	}

	mockService.EXPECT().CreateCustomer(gomock.Any()).Return(expectedCustomer, nil)

	expectedRequestBody :=
		`{
		"Id": 109,
		"Name": "John",
		"DateofBirth": "1999-02-13",
		"City": "GNT",
		"Zipcode": "12345",
		"Status": "1"
	}`

	request, _ := http.NewRequest(http.MethodPost, "/postCustomer", strings.NewReader(expectedRequestBody))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusCreated {
		t.Error("Failed while testing the status code")
	}
}

func Test_postcustomer_status_code_500(t *testing.T) {
	// Arrange

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockCustomerService(ctrl)
	ch := CustomerHandlers{mockService}
	mockService.EXPECT().CreateCustomer(gomock.Any()).Return(nil, errs.NewUnexpectedError("some database error"))

	// Act
	router := gin.Default()
	router.POST("/postCustomer", ch.CreateCustomer)
	expectedRequestBody :=
		`{
	"Id": 109,
	"Name": "John",
	"DateofBirth": "1999-02-13",
	"City": "GNT",
	"Zipcode": "12345",
	"Status": "1"
}`
	request, _ := http.NewRequest(http.MethodPost, "/postCustomer", strings.NewReader(expectedRequestBody))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}

}
