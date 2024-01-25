package app

import (
	"github.com/Sindhuja966/banking/domain"
	"github.com/Sindhuja966/banking/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	//mux := http.NewServeMux()
	router := gin.Default()
	// create an instance of the repository
	customerRepo := domain.NewCustomerRepositoryDb()

	// create an instance of the service and inject the repository
	customerService := service.NewCustomerService(customerRepo)
	//unit testcases
	//customerService := service.NewMockCustomerService(customerRepo)

	// create an instance of the handlers and inject the services
	customerHandlers := &CustomerHandlers{service: customerService}

	router.GET("/customers", customerHandlers.GetAllCustomers)          // To List all Order
	router.GET("/customers/:customer_id", customerHandlers.GetCustomer) // To Get Order Request for particular ID
	//router.POST("/postCustomer", customerHandlers.CreateCustomer)       // To Post Order request
	return router

}

// starting server
func Start() {
	r := SetupRouter()
	r.Run("localhost:8080")
}
