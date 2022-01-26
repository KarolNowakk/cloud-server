package grpc

import (
	"cloud/pkg/search"
	"cloud/pkg/search/searchpb"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type searchServer struct {
	searchpb.UnsafeFileSearchServiceServer

	ss search.Service
}

func (s *searchServer) SearchFiles(ctx context.Context, req *searchpb.SearchRequest) (*searchpb.SearchResponse, error) {
	userID := ctx.Value("userID").(string)
	phrase := req.GetSearchPhrase()

	searchResults, err := s.ss.Search(phrase, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "something went wrong while looking for new files: %v", err)
	}

	fmt.Println(searchResults)
	res := mapToResponse(searchResults)

	return res, nil
}

func mapToResponse(files []search.SearchedFile) *searchpb.SearchResponse {
	var res searchpb.SearchResponse
	var responseElements []*searchpb.SearchFileInfo

	for _, file := range files {
		responseElement := &searchpb.SearchFileInfo{
			Id:         file.ID,
			Name:       file.Name,
			SearchTags: file.SearchTags,
		}

		responseElements = append(responseElements, responseElement)
	}

	res.Files = responseElements

	return &res
}
