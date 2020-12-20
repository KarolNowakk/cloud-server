package grpc

import (
	"cloud/pkg/download"
	"cloud/pkg/download/downloadpb"
	"io"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type downloadServer struct {
	ds download.Service
}

func (s *downloadServer) DownloadFile(req *downloadpb.FileDownloadRequest, stream downloadpb.FileDownloadService_DownloadFileServer) error {
	fileData := &download.FileDownload{
		Name:      req.GetName(),
		Extension: req.GetExtension(),
		Path:      req.GetPath(),
	}
	if err := s.ds.OpenFile(fileData); err != nil {
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
