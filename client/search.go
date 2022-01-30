package main

import (
	"cloud-watcher/proto/searchpb"
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Search struct {
	phrase string
	c      searchpb.FileSearchServiceClient
}

func MakeSearch(c searchpb.FileSearchServiceClient) Search {
	return Search{c: c, phrase: readInput("Enter search phrase (leave empty when want to see all files): ")}
}

func (s Search) Search() {
	req := &searchpb.SearchRequest{
		SearchPhrase: s.phrase,
	}

	resp, err := s.c.SearchFiles(context.Background(), req)
	if err != nil {
		log.Fatal("Error searching files: ", err)
	}

	for _, file := range resp.Files {
		fmt.Printf("ID: %s || Name: %s || Tags: %s \n", file.Id, file.Name, file.SearchTags)
		fmt.Println("--------------------------------------------------------------------")
	}
}
