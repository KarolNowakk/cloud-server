package main

import (
	"cloud-watcher/proto/downloadpb"
	"io"
	"os"

	"github.com/gookit/color"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type Downloader struct {
	c downloadpb.FileDownloadServiceClient

	fileID string
}

func MakeDownloader(c downloadpb.FileDownloadServiceClient) Downloader {
	return Downloader{c: c, fileID: readInput("Provide file ID: ")}
}

func (d Downloader) Download() {
	stream, err := d.c.DownloadFile(context.Background(), request(d.fileID))
	if err != nil {
		log.Fatalf("Error opening a stream: %v", err)
	}

	res, err := stream.Recv()
	if err != nil {
		log.Fatalf("Error receving info request: %v", err)
	}

	name := res.GetInfo().GetName()

	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("download file error: %v", err)
		}

		if err := writeBytes(file, res); err != nil {
			log.Fatalf("Error writing bytes to a file: %v", err)
		}
	}

	color.Success.Tips("Successfully downloaded: " + name)
}

func request(ID string) *downloadpb.FileDownloadRequest {
	return &downloadpb.FileDownloadRequest{Id: ID}
}

func writeBytes(file *os.File, res *downloadpb.FileDownloadResponse) error {
	bytes := res.GetBody().GetBytes()

	if _, err := file.Write(bytes); err != nil {
		return err
	}

	return nil
}
