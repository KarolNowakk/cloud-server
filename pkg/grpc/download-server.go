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

	var bytesStack [][]byte

	for {
		//get bytes from file
		bytes, err := s.ds.ReadBytes()
		if err == io.EOF {
			// if end of file trim 0 bytes and return byte slice without them
			trimmedBytes := trimBytes(bytesStack[0])

			req := byteRequest(trimmedBytes)

			stream.Send(req)
			break
		}

		if err != nil {
			return status.Errorf(codes.Internal, "error while reading file")
		}

		bytesStack = append(bytesStack, bytes)

		//byte slices are pushed to slice and if len is bigger than 2 request is send
		if len(bytesStack) > 1 {
			//bytes at position 0 are sended
			req := byteRequest(bytesStack[0])

			stream.Send(req)

			//sended bytes are removed from bytesStack
			bytesStack = bytesStack[1:]
		}
	}

	return nil
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
		Body: &downloadpb.FileDownloadBody{
			Bytes: bytes,
		},
	}
}
