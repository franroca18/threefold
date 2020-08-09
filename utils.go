package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//generateResponse return respose with basic attributes
func generateResponse(code int) (response Response) {
	response.code = int64(code)
	response.status = http.StatusText(code)
	return response
}

//generateResponseError add to context json response with error message
func generateResponseError(c *gin.Context, code int, err string) {
	var response = generateResponse(code)
	response.message = err
	generateResponseJSON(c, code, response)
}

//generateResponseError add to context json response response with error message
func generateResponseSuccessful(c *gin.Context, code int, customers []*Customer) {
	var response = generateResponse(code)
	response.message = SuccessfulMessage
	response.payload = customers
	generateResponseJSON(c, code, response)
}

//generateResponseJSON add to context json response from code and response given
func generateResponseJSON(c *gin.Context, code int, response Response) {
	c.JSON(code, generateBodyResponseJSON(response))
}

//generateBodyResponseJSON return json gin.H from response
func generateBodyResponseJSON(response Response) gin.H {
	return gin.H{AttributeResponseCode: response.code,
		AttributeResponseStatus:  response.status,
		AttributeResponseMessage: response.message,
		AttributeResponsePayload: response.payload}
}

//getPrimitiveObjectIDFromContext return primitive.ObjectID provided from context
func getPrimitiveObjectIDFromContext(c *gin.Context) (primitive.ObjectID, error) {
	var idNumber = strings.Trim(c.Param(ParamID), ParamID+"=")

	//Check if param is empty
	if len(idNumber) == 0 {
		generateResponseError(c, http.StatusBadRequest, ErrorMessageMissingParamID)
		return primitive.NilObjectID, errors.New(ErrorMessageMissingParamID)
	}

	//transform string to primitive.ObjectID
	var objectID, error = primitive.ObjectIDFromHex(idNumber)
	if error != nil {
		generateResponseError(c, http.StatusBadRequest, error.Error())
		return primitive.NilObjectID, error
	}
	return objectID, nil
}

//getPrimitiveObjectIDFromContext return paramvalue as string
func getValueFromContext(c *gin.Context, paramKey string) string {
	//Get param from context
	return fmt.Sprintf("%v", c.Keys[paramKey])
}

//stringToInt64WithDefaultValue return converted string to int64
func stringToInt64WithDefaultValue(value string, defaultValue int64) int64 {
	n, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return defaultValue
	}
	return n
}

//getCustomerFromContext return *Customer, error from Context binding JSON to Customer
func getCustomerFromContext(c *gin.Context) (*Customer, error) {
	var customer Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		log.Print(err)
		generateResponseError(c, http.StatusBadRequest, err.Error())
		return nil, err
	}
	return &customer, nil
}

//isObjectIDValid return if ObjectId is empty or nil
func isObjectIDValid(id primitive.ObjectID) bool {
	return id.IsZero()
}

//existCustomer return if ObjectID is into database
func existCustomer(id primitive.ObjectID) bool {
	returnedCustomer, err := GetCustomerByID(id)
	log.Printf(InfoMessageCustomerExist, returnedCustomer)
	return err == nil
}
