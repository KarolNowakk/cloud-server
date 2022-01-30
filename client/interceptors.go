package main

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type authInterceptor struct {
	accessToken             string
	methodsExcludedFromAuth []string
}

func newAuthInterceptor(token string, methods []string) *authInterceptor {
	return &authInterceptor{
		accessToken:             token,
		methodsExcludedFromAuth: methods,
	}
}

func (i *authInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, currentMethod string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		for _, excludedMethod := range i.methodsExcludedFromAuth {
			if excludedMethod == currentMethod {
				return invoker(ctx, currentMethod, req, reply, cc, opts...)
			}
		}

		return invoker(i.attachToken(ctx), currentMethod, req, reply, cc, opts...)
	}
}

func (i *authInterceptor) Stream() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		return streamer(i.attachToken(ctx), desc, cc, method, opts...)
	}
}

func (i *authInterceptor) attachToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", i.accessToken)
}
