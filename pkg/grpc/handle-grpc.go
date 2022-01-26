package grpc

import (
	"cloud/pkg/auth"
	"cloud/pkg/auth/authpb"
	"cloud/pkg/download"
	"cloud/pkg/download/downloadpb"
	"cloud/pkg/interceptors"
	"cloud/pkg/permissions"
	"cloud/pkg/search"
	"cloud/pkg/search/searchpb"
	"cloud/pkg/upload"
	"cloud/pkg/upload/uploadpb"

	"google.golang.org/grpc"
)

//NewServer return new grpc server instance
func NewServer(
	dp permissions.Permissions,
	us upload.Service,
	ds download.Service,
	as auth.Service,
	ss search.Service) *grpc.Server {
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
	downloadpb.RegisterFileDownloadServiceServer(s, &downloadServer{ds: ds, p: dp})
	authpb.RegisterAuthServiceServer(s, &authServer{as: as})
	searchpb.RegisterFileSearchServiceServer(s, &searchServer{ss: ss})

	return s
}
