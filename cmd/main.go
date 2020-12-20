package main

import (
	"cloud/pkg/auth"
	"cloud/pkg/download"
	"cloud/pkg/grpc"
	storage "cloud/pkg/storage/mongo"
	"cloud/pkg/upload"
	"context"
	"log"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type connection struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func main() {
	//open badger db
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://karol01:secret@mongo_db:27017"))
	if err != nil {
		log.Fatal("mongo connection not established: " + err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("mongo connection not established: " + err.Error())
	}
	defer client.Disconnect(context.TODO())

	cloudDatabase := client.Database("cloud")

	authConfig := getConfig()

	//init storages
	fileStorageService := storage.NewFileStorageService(cloudDatabase)
	downloadStorageService := storage.NewDownloadStorageService(cloudDatabase)
	authStorageService := storage.NewAuthStorageService(cloudDatabase)

	//init services
	uploadService := upload.NewService(fileStorageService)
	downloadService := download.NewService(downloadStorageService)
	authService := auth.NewService(authStorageService, authConfig)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Faild to listen %v", err)
	}

	grpcServer := grpc.NewServer(uploadService, downloadService, authService)

	log.Println("Listening...")
	log.Fatal(grpcServer.Serve(lis))
}

func getConfig() *auth.Config {
	JWTkey := []byte("fdfiasomcmd1232e1m32d12jnd24do1idnoijn531oi5xmj535m341232x2445")

	return &auth.Config{
		Key:      JWTkey,
		Duration: 7 * 24 * time.Hour,
	}
}
