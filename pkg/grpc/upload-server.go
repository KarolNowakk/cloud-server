package grpc

import (
	"cloud/pkg/upload"
	"cloud/pkg/upload/uploadpb"
	"fmt"
	"io"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type uploadServer struct {
	uploadpb.UnsafeFileUploadServiceServer

	us upload.Service
}

func (s uploadServer) UploadFile(stream uploadpb.FileUploadService_UploadFileServer) error {
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

	userID := stream.Context().Value("userID").(string)

	//get file info
	req, err := stream.Recv()
	if err != nil {
		log.Printf("Error receving: %v", err)
		return status.Errorf(codes.Internal, "error while receving file info")
	}

	if err := s.us.CreateFileIfNotExistsAndOpen(stream.Context(), req.GetInfo(), userID); err != nil {
		log.Printf("Error createing db model: %v", err)
		return status.Errorf(codes.Internal, "error creating or opening a file")
	}

	fmt.Println(req.GetInfo())

	//save file info to database
	if err := s.us.UpdateOrCreateFile(stream.Context(), userID); err != nil {
		return status.Errorf(codes.Internal, "error while saving file info: %v", err)
	}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error receving: %v", err)
			return status.Errorf(codes.Internal, "error while receving file bytes")
		}

		//write bytes to file on the server
		if err := s.us.WriteBytes(req.GetBody()); err != nil {
			log.Printf("Error writing to file: %v", err)
			return status.Errorf(codes.Internal, "error writing to file: %v", err)
		}
	}

	log.Println("File Uploaded.")

	return nil
}
