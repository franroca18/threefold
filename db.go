package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	paginate "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Timeout operations after N seconds
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

// GetConnection - Retrieves a client to the DocumentDB
func getConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	/*username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	clusterEndpoint := os.Getenv("MONGODB_ENDPOINT")*/
	username := "franroca"
	password := "franrocaMONGODB"
	clusterEndpoint := "localhost:27017"

	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Printf(ErrorMessageCreateClient, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf(ErrorMessageConnectCluster, err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf(ErrorMessagePingCluster, err)
	}

	fmt.Println(InfoMessageConnectedMongo)
	return client, ctx, cancel
}

// GetAllCustomers Retrives all customers from the db
func GetAllCustomers(pageS string, limitS string) ([]*Customer, error) {
	var customers []*Customer
	var page int64 = stringToInt64WithDefaultValue(pageS, DefaultPage)
	var limit int64 = stringToInt64WithDefaultValue(limitS, DefaultLimit)
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	db := client.Database(DatabaseName)
	collection := db.Collection(DatabaseName)
	paginatedData, err := paginate.New(collection).Limit(limit).Page(page).Filter(bson.M{}).Find()
	if err != nil {
		return nil, err
	}

	for _, raw := range paginatedData.Data {
		var customer *Customer
		if marshallErr := bson.Unmarshal(raw, &customer); marshallErr == nil {
			customers = append(customers, customer)
		}

	}

	return customers, nil
}

// GetCustomerByID Retrives a Customer by its id from the db
func GetCustomerByID(id primitive.ObjectID) (*Customer, error) {
	var customer *Customer

	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	db := client.Database(DatabaseName)
	collection := db.Collection(DatabaseName)
	result := collection.FindOne(ctx, bson.M{"_id": id})
	if result == nil {
		return nil, errors.New(ErrorMessageNotFoundCustomer)
	}
	err := result.Decode(&customer)

	if err != nil {
		log.Printf(ErrorMessageFailMarshalling, err)
		return nil, err
	}
	log.Printf("Customer: %v", customer)
	return customer, nil
}

//Create creating a Customer in a mongo or document db
func Create(customer *Customer) (primitive.ObjectID, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	customer.IDNumber = primitive.NewObjectID()
	customer.Lastupdated = 0

	result, err := client.Database(DatabaseName).Collection(DatabaseName).InsertOne(ctx, customer)
	if err != nil {
		log.Printf(ErrorMessageCreateCustomer, err)
		return primitive.NilObjectID, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid, nil
}

//Update updating an existing customer in a mongo
func Update(customer *Customer) (*Customer, error) {
	var updatedCustomer *Customer
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	if !existCustomer(customer.IDNumber) {
		return nil, errors.New(ErrorMessageNotFoundCustomer)
	}

	customer.Lastupdated = time.Now().Unix()
	update := bson.M{
		"$set": customer,
	}

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		Upsert:         &upsert,
		ReturnDocument: &after,
	}

	err := client.Database(DatabaseName).Collection(DatabaseName).FindOneAndUpdate(ctx, bson.M{"_id": customer.IDNumber}, update, &opt).Decode(&updatedCustomer)
	if err != nil {
		log.Printf(ErrorMessageSaveCustomer, err)
		return nil, err
	}
	return updatedCustomer, nil
}

//Delete deleting an existing customer in a mongo
func Delete(id primitive.ObjectID) error {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	if !existCustomer(id) {
		return errors.New(ErrorMessageNotFoundCustomer)
	}
	deletedCustomer, err := client.Database(DatabaseName).Collection(DatabaseName).DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Printf(ErrorMessageDeleteCustomer, err)
		return err
	}
	log.Printf("Deleted Customer: %v", deletedCustomer)
	return nil
}
