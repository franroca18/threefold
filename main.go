package main

import (
	"net/http"
	"threefold/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Threefold Tech Challenge API
// @version 1.0
// @description This is a API to handle a full CRUDL microservice for a customer

// @host localhost:9000
// @BasePath /customers
// @query.collection.format multi

// @contact.name Francisco Rodriguez
// @contact.email franroca18@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	docs.SwaggerInfo.Host = "localhost:9000"

	r := gin.Default()

	r.GET("/customers/:id", handleGetCustomer)
	r.GET("/customers/", handleGetCustomers)
	r.PUT("/customers/", handleCreateOrUpdateCustomer)
	r.DELETE("/customers/:id", handleDeleteCustomer)

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":9000")
}

// handleGetCustomer godoc
// @Summary Get a single record of customer
// @Description get customer by ID
// @ID get-primitive.ObjectID-to-delete-by-string
// @Accept  plain
// @Produce  json
// @Param id path string true "ID Number"
// @Success 200 {object} main.Response
// @Failure 400 {object} main.Response
// @Failure 404 {object} main.Response
// @Failure 500 {object} main.Response
// @Router /customers/{id} [get]
func handleGetCustomer(c *gin.Context) {
	var objectID, error = getPrimitiveObjectIDFromContext(c)
	if error == nil {
		var loadedCustomer, err = GetCustomerByID(objectID)
		if err != nil {
			generateResponseError(c, http.StatusNotFound, err.Error())
			return
		}
		generateResponseSuccessful(c, http.StatusOK, []*Customer{loadedCustomer})
	}
}

// handleGetCustomers godoc
// @Summary Get a list record of customer
// @Description get customer recorded
// @Accept  json
// @Produce  json
// @Success 200 {object} main.Response
// @Failure 400 {object} main.Response
// @Failure 404 {object} main.Response
// @Router /customers/ [get]
func handleGetCustomers(c *gin.Context) {
	var loadedCustomers, err = GetAllCustomers(c.Query(ParamPage), c.Query(ParamLimit))
	if err != nil {
		generateResponseError(c, http.StatusNotFound, err.Error())
		return
	}
	generateResponseSuccessful(c, http.StatusOK, loadedCustomers)
}

// CreateOrUpdateCustomer godoc
// @Summary Create or update a customer from a given json
// @Description Create a customer if IDNumber is not provided otherwise customer is updated if exist
// @ID main.Customer-by-json
// @Accept  json
// @Produce  json
// @Success 200 {object} main.Response
// @Failure 500 {object} main.Response
// @Router /customers/ [put]
// @Path("/customers")
func handleCreateOrUpdateCustomer(c *gin.Context) {
	var customer *Customer
	customer, err := getCustomerFromContext(c)
	if err == nil {
		if isObjectIDValid(customer.IDNumber) {
			handleCreateCustomer(c, customer)
		} else {
			handleUpdateCustomer(c, customer)
		}
	}
}

// CreateCustomer godoc
// @Summary create a new customer
// @Description create customer by customer given, it is an internal method
// @Produce  json
// @Success 200 {object} main.Response
// @Failure 500 {object} main.Response
func handleCreateCustomer(c *gin.Context, customer *Customer) {
	id, err := Create(customer)
	if err != nil {
		generateResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	var savedCustomer, errGettingCustomer = GetCustomerByID(id)
	if errGettingCustomer != nil {
		generateResponseSuccessful(c, http.StatusOK, nil)
		return
	}
	generateResponseSuccessful(c, http.StatusOK, []*Customer{savedCustomer})
}

// UpdateCustomer godoc
// @Summary update a customer
// @Description Update a customer by IDNumber, if customer does not exist an error is raised
// @Produce  json
// @Success 200 {object} main.Response
// @Failure 500 {object} main.Response
func handleUpdateCustomer(c *gin.Context, customer *Customer) {
	savedCustomer, err := Update(customer)
	if err != nil {
		generateResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	generateResponseSuccessful(c, http.StatusOK, []*Customer{savedCustomer})
}

// DeleteCustomer godoc
// @Summary Delete a single record of customer
// @Description delete customer by ID
// @ID get-primitive.ObjectID-by-string
// @Accept  plain
// @Produce  json
// @Param id path string true "ID Number"
// @Success 200 {object} main.Response
// @Failure 500 {object} main.Response
// @Router /customers/{id} [delete]
func handleDeleteCustomer(c *gin.Context) {
	var objectID, error = getPrimitiveObjectIDFromContext(c)
	if error == nil {
		err := Delete(objectID)
		if err != nil {
			generateResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		generateResponseSuccessful(c, http.StatusOK, nil)
	}
}
