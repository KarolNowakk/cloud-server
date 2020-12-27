package storage

import (
	"cloud/pkg/upload"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func dbMock() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://karol01:secret@127.0.0.1:27017"))

	if err != nil {
		log.Fatal("mongo connection not established: " + err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		log.Fatal("mongo connection not established: " + err.Error())
	}

	db := client.Database("cloudTest")

	return db
}

func clearDatabase(db *mongo.Database) {
	_, _ = db.Collection("users").DeleteMany(context.Background(), bson.M{})
	_, _ = db.Collection("files").DeleteMany(context.Background(), bson.M{})
	_, _ = db.Collection("folders").DeleteMany(context.Background(), bson.M{})
	_, _ = db.Collection("tokens").DeleteMany(context.Background(), bson.M{})
}

func getSampleInsertedUser(db *mongo.Database) *userModel {
	user := &userModel{
		Username: "test01110101",
		Email:    "email@test.com",
		Password: "2a$10$tSJtkEIc3S1fN4MKfUmEWOcpK2fbgzJ1O3t14OYZqm.sBPNwBVXKu", //PassWord12
	}

	res, _ := db.Collection("users").InsertOne(context.Background(), user)
	oid, _ := res.InsertedID.(primitive.ObjectID)
	user.ID = oid

	return user
}

func getSampleInsertedFile(db *mongo.Database, owner primitive.ObjectID) *fileModel {
	file := getSampleFile(db, owner)

	res, _ := db.Collection("files").InsertOne(context.Background(), file)
	oid, _ := res.InsertedID.(primitive.ObjectID)
	file.ID = oid

	return file
}

func getSampleFile(db *mongo.Database, owner primitive.ObjectID) *fileModel {
	file := &fileModel{
		Name:      "file",
		Extension: ".pdf",
		FullPath:  "testing/tester/testosteron/file.pdf",
		Owner:     owner,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	return file
}

func getSampleFileAsUploadFile(db *mongo.Database, owner primitive.ObjectID) *upload.File {
	file := getSampleFile(db, owner)

	return &upload.File{
		Name:      file.Name,
		Extension: file.Extension,
		FullPath:  file.FullPath,
		Owner:     file.Owner.Hex(),
	}
}
