package main

import (
	"cloud-watcher/proto/downloadpb"
	"context"

	"github.com/gookit/color"
	log "github.com/sirupsen/logrus"
)

type Delete struct {
	c  downloadpb.FileDownloadServiceClient
	id string
}

func MakeDelete(c downloadpb.FileDownloadServiceClient) Delete {
	return Delete{
		c:  c,
		id: readInput("Provide ID of a file that you want to delete: "),
	}
}

func (d *Delete) Delete() {
	req := downloadpb.FileDeleteRequest{
		Id: d.id,
	}

	resp, err := d.c.DeleteFile(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error deleting the file: %v", err)
	}

	color.Success.Tips(resp.GetMsg())
}
