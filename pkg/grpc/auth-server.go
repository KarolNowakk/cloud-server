package grpc

import (
	"cloud/pkg/auth"
	"cloud/pkg/auth/authpb"
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authServer struct {
	as auth.Service
}

func (s *authServer) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	res := authpb.RegisterResponse{
		Success: true,
		Msg:     "Successfully registering a new user",
	}

	err := s.as.Validate(req)

	if err != nil {
		log.Printf("Register error: %s", err.Error())
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	err = s.as.Register(req)

	if err != nil {
		log.Printf("Register error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "error while registering")
	}

	return &res, nil
}
func (s *authServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	token, err := s.as.Login(req)
	if err != nil {
		if err == auth.ErrBadCredentials {
			return nil, status.Error(codes.Unauthenticated, "bad credentials")
		}

		return nil, status.Error(codes.Internal, "something went wrong")
	}

	return &authpb.LoginResponse{
		Type:           "Bearer ",
		ExpirationTime: token.ExpirationTime,
		Token:          token.TokenString,
		UserID:         token.UserID,
	}, nil
}
