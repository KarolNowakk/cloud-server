package main

import (
	"cloud/pkg/auth"
	"cloud/pkg/download"
	"cloud/pkg/grpc"
	"cloud/pkg/permissions"
	"cloud/pkg/scanner"
	storage "cloud/pkg/storage/mongo"
	"cloud/pkg/upload"
	"context"
	"log"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
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
	//load env config variables
	JWTkey, tokenDuration, dbUsername, dbPassword, dbHost, dbPort := getConfig()

	//open mongo db
	cloudDatabase := getMongoConnection(dbUsername, dbPassword, dbHost, dbPort)
	defer cloudDatabase.Client().Disconnect(context.TODO())

	//init storages
	fileStorageService := storage.NewFileStorageService(cloudDatabase)
	downloadStorageService := storage.NewDownloadStorageService(cloudDatabase)
	authStorageService := storage.NewAuthStorageService(cloudDatabase)
	scannerStorageService := storage.NewScannerStorageService(cloudDatabase)

	//init services
	uploadService := upload.NewService(fileStorageService)
	downloadService := download.NewService(downloadStorageService)
	authService := auth.NewService(authStorageService, JWTkey, tokenDuration)
	scannerService := scanner.NewService(scannerStorageService)

	downloadPermissions := permissions.NewDownloadPermissions()
	uploadPermissions := permissions.NewUploadPermissions()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	grpcServer := grpc.NewServer(
		downloadPermissions,
		uploadPermissions,
		uploadService,
		downloadService,
		authService,
		scannerService)

	log.Println("Listening...")
	log.Fatal(grpcServer.Serve(lis))
}

func getConfig() ([]byte, time.Duration, string, string, string, string) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("unable to load .env file")
	}

	JWTKey := []byte(os.Getenv("JWTKEY"))

	dbUsername := os.Getenv("MONGO_USER")
	dbPassword := os.Getenv("MONGO_PASSWORD")
	dbHost := os.Getenv("MONGO_HOST")
	dbPort := os.Getenv("MONGO_PORT")

	return JWTKey, 7 * 24 * time.Hour, dbUsername, dbPassword, dbHost, dbPort
}

func getMongoConnection(dbUsername, dbPassword, dbHost, dbPort string) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(
		"mongodb://" + dbUsername + ":" + dbPassword + "@" + dbHost + ":" + dbPort))

	if err != nil {
		log.Fatal("mongo connection not established: " + err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		log.Fatal("mongo connection not established: " + err.Error())
	}

	db := client.Database("cloud")

	list, err := db.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		log.Fatalf("mongo connection not established: %v", err)
	}

	log.Println(list)

	return db
}
