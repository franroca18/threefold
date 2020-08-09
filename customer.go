package main

import "go.mongodb.org/mongo-driver/bson/primitive"

// Customer represents the customers for this application
//
// A customer is the principal object for this application.
//
// swagger:model
type Customer struct {
	// the id for this customer
	//
	// required: true
	IDNumber primitive.ObjectID `json:"idNumber,omitempty" bson:"_id,omitempty"`

	// the name for this customer
	//
	// required: false
	Name string `json:"name,omitempty" bson:"name,omitempty"`

	// the surname for this customer
	//
	// required: false
	Surname string `json:"surname,omitempty" bson:"surname,omitempty"`

	// the email for this customer
	//
	// required: false
	Email string `json:"email,omitempty" bson:"email,omitempty"`

	// the intials for this customer
	//
	// required: false
	Initials string `json:"initials,omitempty" bson:"initials,omitempty"`

	// the mobile for this customer
	//
	// required: false
	Mobile string `json:"mobile,omitempty" bson:"mobile,omitempty"`

	// the lastUpadate for this customer
	//
	// required: false
	Lastupdated int64 `json:"lastUptade,omitempty" bson:"lastUpdate,omitempty"`
}
