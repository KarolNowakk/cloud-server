package main

import (
	"cloud-watcher/proto/authpb"
	"cloud-watcher/proto/downloadpb"
	"cloud-watcher/proto/searchpb"
	"cloud-watcher/proto/uploadpb"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
	"google.golang.org/grpc"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	db := openDatabase()

	token, err := findConfigString(db, keyToken)
	if err != nil {
		log.Fatal(err)
	}

	auth := newAuthInterceptor(token, []string{"/auth.AuthService/Login", "/auth.AuthService/Register"})

	cc, err := grpc.Dial("127.0.0.1:50051",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(auth.Unary()),
		grpc.WithStreamInterceptor(auth.Stream()),
	)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	if len(os.Args) > 1 {
		resolveArguments(cc, db)
		return
	}
}

func openDatabase() *bolt.DB {
	// Open database file.
	db, err := bolt.Open("./db/bolt", 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	// Start writable transaction.
	tx, err := db.Begin(true)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	if _, err := tx.CreateBucketIfNotExists([]byte("config")); err != nil {
		log.Fatal(err)
	}
	if _, err := tx.CreateBucketIfNotExists([]byte("files")); err != nil {
		log.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	return db
}

func resolveArguments(cc *grpc.ClientConn, db *bolt.DB) {
	ac := authpb.NewAuthServiceClient(cc)
	sc := searchpb.NewFileSearchServiceClient(cc)
	uc := uploadpb.NewFileUploadServiceClient(cc)
	dc := downloadpb.NewFileDownloadServiceClient(cc)

	switch arg := os.Args[1]; arg {
	case "register":
		auth := MakeAuth(ac)
		auth.Register()
	case "login":
		auth := MakeAuth(ac)
		auth.Login(db)
	case "delete":
		delete := MakeDelete(dc)
		delete.Delete()
	case "search":
		search := MakeSearch(sc)
		search.Search()
	case "upload":
		uploader := MakeUploader(uc, os.Args)
		uploader.Upload()
	case "download":
		downloader := MakeDownloader(dc)
		downloader.Download()
	default:
		fmt.Println("Action not found.")
	}
}
