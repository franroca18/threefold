package main

// SuccessfulMessage return Successful
const SuccessfulMessage string = "Successful"

// ErrorMessageMissingParamID return error message
const ErrorMessageMissingParamID string = "Missing param: id"

// ErrorMessageCreateCustomer return error message
const ErrorMessageCreateCustomer string = "Could not create Customer: %v"

// ErrorMessageCreateClient return error message
const ErrorMessageCreateClient string = "Failed to create client: %v"

// ErrorMessageConnectCluster return error message
const ErrorMessageConnectCluster string = "Failed to connect to cluster: %v"

// ErrorMessageNotFoundCustomer return error message
const ErrorMessageNotFoundCustomer string = "Could not find a Customer"

// ErrorMessageFailMarshalling return error message
const ErrorMessageFailMarshalling string = "Failed marshalling %v"

// ErrorMessagePingCluster return error message
const ErrorMessagePingCluster string = "Failed to ping cluster: %v"

// ErrorMessageSaveCustomer return error message
const ErrorMessageSaveCustomer string = "Could not save Customer: %v"

// ErrorMessageDeleteCustomer return error message
const ErrorMessageDeleteCustomer string = "Could not delete Customer: %v"

// InfoMessageDeletedCustomer return info message
const InfoMessageDeletedCustomer string = "Deleted Customer: %v"

// InfoMessageCustomerExist return info message
const InfoMessageCustomerExist string = "Customer exist: %v"

// InfoMessageConnectedMongo return info message
const InfoMessageConnectedMongo string = "Connected to MongoDB!"

// DatabaseName return string database name
const DatabaseName string = "customers"

// ParamID return string paramKey
const ParamID string = "id"

// ParamLimit return string paramKey
const ParamLimit string = "limit"

// ParamPage return string paramKey
const ParamPage string = "page"

// AttributeResponseCode return name of response attribute
const AttributeResponseCode string = "code"

// AttributeResponseStatus return name of response attribute
const AttributeResponseStatus string = "status"

// AttributeResponseMessage return name of response attribute
const AttributeResponseMessage string = "message"

// AttributeResponsePayload return name of response attribute
const AttributeResponsePayload string = "payload"

// DefaultLimit return default limit
const DefaultLimit int64 = 2

// DefaultPage return default page
const DefaultPage int64 = 1
