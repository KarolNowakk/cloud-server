package grpc

import (
	"cloud/pkg/upload"
	"cloud/pkg/upload/uploadpb"
	"io"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type uploadServer struct {
	us upload.Service
}

func (s *uploadServer) UploadFile(stream uploadpb.FileUploadService_UploadFileServer) error {
	//send response
	defer func() {
		response := uploadpb.FileUploadResponse{Ok: true, Msg: "file has been uploaded"}

		//handle panics if any
		if r := recover(); r != nil {
			log.Println(r)
			response = uploadpb.FileUploadResponse{Ok: false, Msg: "something went wrong"}
		}

		stream.SendAndClose(&response)
	}()

	//get file info
	req, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Internal, "error while receving file info")
	}

	if err := s.us.CreateFileIfNotExistsAndOpen(req.GetInfo()); err != nil {
		return status.Errorf(codes.Internal, "error creating or opening a file")
	}

	//save file info to database
	if err := s.us.UpdateOrCreateFile(); err != nil {
		return status.Errorf(codes.Internal, "error while saving file info")
	}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Errorf(codes.Internal, "error while receving file bytes")
		}

		//write bytes to file on the server
		if err := s.us.WriteBytes(req.GetBody()); err != nil {
			return status.Errorf(codes.Internal, "error writing to file")
		}
	}

	return nil
}
