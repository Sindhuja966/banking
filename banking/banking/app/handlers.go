package app

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Sindhuja966/banking/domain"
	"github.com/Sindhuja966/banking/dto"
	"github.com/Sindhuja966/banking/service"
	"github.com/gin-gonic/gin"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func convertToNewCustomerRequest(customer domain.Customer) *dto.NewCustomerRequest {

	return &dto.NewCustomerRequest{
		Id:          customer.Id,
		Name:        customer.Name,
		City:        customer.City,
		Zipcode:     customer.Zipcode,
		DateofBirth: customer.DateofBirth,
		Status:      customer.Status,
	}
}

func (ch *CustomerHandlers) GetAllCustomers(c *gin.Context) {

	status := c.Query("status")
	ctx, cancel := context.WithTimeout(c, 5*time.Microsecond)
	customers, err := ch.service.GetAllCustomer(ctx, status)
	defer cancel()

	if err != nil {
		fmt.Printf("Error: %+v", err)
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	} else {
		c.JSON(http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) GetCustomer(c *gin.Context) {
	id := c.Param("customer_id")
	fmt.Printf("Received request for Customer ID: %s", id)
	customerid, _ := strconv.Atoi(id)
	customers, err := ch.service.GetCustomer(customerid)

	if err != nil {
		fmt.Printf("Error: %+v", err)
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	} else {
		c.JSON(http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) CreateCustomer(c *gin.Context) {
	var newCustomer domain.Customer

	//Bind the JSON data from the request body to the newOrder struct
	if err := c.BindJSON(&newCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	parsedRequest := convertToNewCustomerRequest(newCustomer)
	createdCustomer, err := ch.service.CreateCustomer(parsedRequest)
	if err != nil {
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.JSON(http.StatusCreated, createdCustomer)

}
