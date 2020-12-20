package grpc

import (
	"cloud/pkg/auth"
	"cloud/pkg/auth/authpb"
	"cloud/pkg/download"
	"cloud/pkg/download/downloadpb"
	"cloud/pkg/interceptors"
	"cloud/pkg/upload"
	"cloud/pkg/upload/uploadpb"

	"google.golang.org/grpc"
)

//NewServer return new grpc server instance
func NewServer(us upload.Service, ds download.Service, as auth.Service) *grpc.Server {
	methodsExcludedFromAuth := []string{
		"/auth.AuthService/Login",
		"/auth.AuthService/Register",
	}

	interceptor := interceptors.NewAuthInterceptor(as, methodsExcludedFromAuth)
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}

	s := grpc.NewServer(serverOptions...)

	uploadpb.RegisterFileUploadServiceServer(s, &uploadServer{us: us})
	downloadpb.RegisterFileDownloadServiceServer(s, &downloadServer{ds: ds})
	authpb.RegisterAuthServiceServer(s, &authServer{as: as})

	return s
}
