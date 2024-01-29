package app

import (
	"fmt"
	"net/http"

	"github.com/Sindhuja966/banking/domain"
	"github.com/Sindhuja966/banking/service"
	"github.com/gin-gonic/gin"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomers(c *gin.Context) {
	status := c.Request.URL.Query().Get("status")
	customers, _ := ch.service.GetAllCustomer(status)

	c.JSON(http.StatusOK, customers)
}

func (ch *CustomerHandlers) GetCustomer(c *gin.Context) {
	id := c.Param("customer_id")
	fmt.Printf("Received request for Customer ID: %s", id)
	customers, err := ch.service.GetCustomer(id)

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

	savedCustomer, err := ch.service.SaveCustomer(newCustomer)
	if err != nil {
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.JSON(http.StatusCreated, savedCustomer)

}
