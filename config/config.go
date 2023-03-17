package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	_ "github.com/lib/pq"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectMongo() {
	// clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017,127.0.0.1:27017,127.0.0.1:27017/?replicaSet=myReplicaSet") //.SetServerSelectionTimeout(30*time.Second)
	clientOptions := options.Client().ApplyURI(mongoConnectionString).SetDirect(true)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
	}

	defer client.Disconnect(context.Background())
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	oplogCollection := client.Database(local).Collection(oplog)
	filter := bson.M{
		"op": bson.M{"$nin": []string{"n", "c"}},
		"$and": []bson.M{
			bson.M{"ns": bson.M{"$not": bson.M{"$regex": "^(admin|config)\\."}}},
			bson.M{"ns": bson.M{"$not": bson.M{"$eq": ""}}},
		},
	}

	cursor, err := oplogCollection.Find(context.Background(), filter)
	if err != nil {
		panic(err)
	}

	// Write documents to JSON file
	var results []bson.M
	if err = cursor.All(context.Background(), &results); err != nil {
		panic(err)
	}

	if err = writeOplogToJSON(outputFile, results); err != nil {
		panic(err)
	}
}

func writeOplogToJSON(filename string, data interface{}) error {
	// Encode data as JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Write JSON data to file
	if err = ioutil.WriteFile(filename, jsonData, 0644); err != nil {
		return err
	}

	return nil
}
