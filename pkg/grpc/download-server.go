package grpc

import (
	"cloud/pkg/download"
	"cloud/pkg/download/downloadpb"
	"context"
	"fmt"
	"io"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type permission interface {
	CanDelete(userID, ownerID string) error
	CanDownload(userID, ownerID string) error
}
type downloadServer struct {
	downloadpb.UnsafeFileDownloadServiceServer

	ds download.Service
	p  permission
}

func (s *downloadServer) DownloadFile(req *downloadpb.FileDownloadRequest, stream downloadpb.FileDownloadService_DownloadFileServer) error {
	userID := stream.Context().Value("userID").(string)
	fileID := req.GetId()

	fileData, err := s.ds.FindFile(stream.Context(), fileID)
	if err != nil {
		return status.Errorf(codes.NotFound, "File not found.")
	}

	if err := s.p.CanDownload(userID, fileData.Owner); err != nil {
		return status.Errorf(codes.PermissionDenied, "You don't have acces to requested data.")
	}

	file, err := s.ds.OpenFile(fileData.Path)
	if err != nil {
		return status.Errorf(codes.NotFound, "File not foud")
	}

	stream.Send(infoRequest(fileData.Name))

	bufferSize := 2 * 1024

	var bytesStack [][]byte

	var offset int64 = 0

	for {
		bytesRead, err := s.ds.ReadChunk(file, bufferSize, offset, 0)

		if err == io.EOF {
			if len(bytesStack) < 1 {
				return status.Errorf(codes.Internal, "Invalid file.")
			}

			trimmedBytes := trimBytes(bytesStack[0])

			stream.Send(byteRequest(trimmedBytes))

			break
		}

		if err != nil {
			return status.Errorf(codes.Internal, "Error reading bytes.")
		}

		bytesStack = append(bytesStack, bytesRead)

		if len(bytesStack) > 1 {
			if err := stream.Send(byteRequest(bytesStack[0])); err != nil {
				return status.Errorf(codes.Internal, "Error sending bytes.")
			}

			bytesStack = bytesStack[1:]
		}

		offset += int64(bufferSize)
	}

	return nil
}

func infoRequest(name string) *downloadpb.FileDownloadResponse {
	return &downloadpb.FileDownloadResponse{
		Data: &downloadpb.FileDownloadResponse_Info{
			Info: &downloadpb.FileDownloadInfo{Name: name},
		},
	}
}

func trimBytes(bytes []byte) []byte {
	var firstZeroByteIndex int
	firstZeroWasHitted := false

	for i, singleByte := range bytes {
		if singleByte == 0 && !firstZeroWasHitted {
			firstZeroByteIndex = i
			firstZeroWasHitted = true
		}
		if singleByte != 0 {
			firstZeroByteIndex = -1
			firstZeroWasHitted = false
		}
	}

	if firstZeroByteIndex == -1 {
		firstZeroByteIndex = len(bytes)
	}

	bytes = bytes[:firstZeroByteIndex]

	return bytes
}

func byteRequest(bytes []byte) *downloadpb.FileDownloadResponse {
	return &downloadpb.FileDownloadResponse{
		Data: &downloadpb.FileDownloadResponse_Body{
			Body: &downloadpb.FileDownloadBody{Bytes: bytes},
		},
	}
}

func (s downloadServer) DeleteFile(ctx context.Context, in *downloadpb.FileDeleteRequest) (*downloadpb.FileDeleteResponse, error) {
	userID := ctx.Value("userID").(string)
	fileID := in.GetId()

	fileData, err := s.ds.FindFile(ctx, fileID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "File not found.")
	}

	if err := s.p.CanDelete(userID, fileData.Owner); err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "you cannot delete this file")
	}

	fmt.Println(fileData.Path)
	err = os.Remove(fileData.Path)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error deleting file %v", err)
	}

	if err := s.ds.DeleteFile(ctx, fileID); err != nil {
		return nil, status.Errorf(codes.Internal, "Error deleting file %v", err)
	}

	res := &downloadpb.FileDeleteResponse{
		Ok:  true,
		Msg: fmt.Sprintf("File %s has been succesfully deleted", fileData.Name),
	}

	return res, nil
}
