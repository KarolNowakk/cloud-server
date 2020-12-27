package grpc

import (
	"cloud/pkg/scanner"
	"cloud/pkg/scanner/scannerpb"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type scannerServer struct {
	ss scanner.Service
}

func (s *scannerServer) ScanForNewFiles(ctx context.Context, req *scannerpb.ScannerRequest) (*scannerpb.ScannerResponse, error) {
	userID := ctx.Value("userID").(string)

	date := req.GetDateTime().AsTime()

	files, err := s.ss.Scan(date, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "something went wrong while looking for new files")
	}

	res := mapToResponse(files)

	return res, nil
}

func mapToResponse(files []scanner.ScannedFile) *scannerpb.ScannerResponse {
	var res scannerpb.ScannerResponse
	var responseElements []*scannerpb.ScannerFileInfo

	for _, file := range files {
		responseElement := &scannerpb.ScannerFileInfo{
			Id:        file.ID,
			Name:      file.Name,
			Extension: file.Extension,
			Path:      file.Path,
			Owner:     file.Owner,
		}

		responseElements = append(responseElements, responseElement)
	}

	res.Files = responseElements

	return &res
}
