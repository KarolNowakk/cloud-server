package grpc

import (
	"cloud/pkg/download"
	"cloud/pkg/download/downloadpb"
	"cloud/pkg/permissions"
	"io"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type downloadServer struct {
	ds download.Service
	p  permissions.DownladPermissions
}

func (s *downloadServer) DownloadFile(req *downloadpb.FileDownloadRequest, stream downloadpb.FileDownloadService_DownloadFileServer) error {
	userID := stream.Context().Value("userID").(string)

	if err := s.p.CanDownload(userID, req); err != nil {
		return status.Errorf(codes.PermissionDenied, "you don't have acces to requested data")
	}

	if err := s.ds.OpenFile(req, userID); err != nil {
		return status.Errorf(codes.NotFound, "file not foud")
	}

	for {
		bytes, err := s.ds.ReadBytes()
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Errorf(codes.Internal, "error while reading file")
		}

		res := &downloadpb.FileDownloadResponse{
			Body: &downloadpb.FileDownloadBody{
				Bytes: bytes,
			},
		}

		stream.Send(res)
	}

	go s.ds.RecordDownloadFile()

	return nil
}
