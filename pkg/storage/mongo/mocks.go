package storage

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func dbMock() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://karol01:secret@0.0.0.0:27017"))

	if err != nil {
		log.Fatal("mongo connection not established: " + err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		log.Fatal("mongo connection not established: " + err.Error())
	}

	db := client.Database("cloudTest")

	list, err := db.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		log.Fatalf("mongo connection not established: %v", err)
	}

	log.Println(list)

	return db
}

func clearCollection(coll *mongo.Collection) {
	filter := bson.M{}

	_, err := coll.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
}
