package main

// Response represents the response for this application
//
// A response is use a model to json response.
//
// swagger:model
type Response struct {
	// the http status code
	//
	// required: false
	code int64

	// the http status description
	//
	// required: false
	status string

	// the message for this response
	//
	// required: false
	message string

	// the array of customers for this response
	//
	// required: false
	payload []*Customer
}
