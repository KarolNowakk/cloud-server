package interceptors

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

//AuthInterceptor is interface with intercepotrs methods
type AuthInterceptor interface {
	Unary() grpc.UnaryServerInterceptor
	Stream() grpc.StreamServerInterceptor
}

//AuthService is interface that must be satistified to be used as a auth service
type AuthService interface {
	Verify(ctx context.Context, tokenString string) (string, error)
}

//NewAuthInterceptor creates new auth interceptor
func NewAuthInterceptor(s AuthService, metodsExcludedFromAuth []string) AuthInterceptor {
	return &authInterceptor{service: s, metodsExcludedFromAuth: metodsExcludedFromAuth}
}

type authInterceptor struct {
	service                AuthService
	metodsExcludedFromAuth []string
}

func (i authInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		id, err := i.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		ctx = context.WithValue(ctx, "userID", id)

		return handler(ctx, req)
	}
}

type wrapedStream struct {
	grpc.ServerStream
	userID string
}

func newwrapedStream(ss grpc.ServerStream) *wrapedStream {
	return &wrapedStream{ServerStream: ss}
}

func (s wrapedStream) Context() context.Context {
	return context.WithValue(s.ServerStream.Context(), "userID", s.userID) // TODO: add some struct as key
}

func (i authInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		id, err := i.authorize(ss.Context(), info.FullMethod)
		fmt.Println(id, err)
		if err != nil {
			return err
		}

		wStream := newwrapedStream(ss)
		wStream.userID = id

		return handler(srv, wStream)
	}
}

func (i authInterceptor) authorize(ctx context.Context, currentMethod string) (string, error) {
	//check if method do not requires authentication
	for _, excludedMethod := range i.metodsExcludedFromAuth {
		if excludedMethod == currentMethod {
			return "TOKEN", nil
		}
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "metadata does not exists")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	id, err := i.service.Verify(ctx, accessToken)
	if err != nil {
		return "", status.Errorf(codes.Unauthenticated, "unauthorized")
	}

	return id, nil
}
