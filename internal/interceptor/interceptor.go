package interceptor

import (
	"context"
	"time"

	"github.com/aclgo/grpc-jwt/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Interceptor interface {
	Logger(context.Context, any, *grpc.UnaryServerInfo, *grpc.UnaryHandler) (any, error)
	Metrics(context.Context, any, *grpc.UnaryServerInfo, *grpc.UnaryHandler) (any, error)
}

type InterceptorManager struct {
	logger logger.Logger
}

func NewInterceptor(logger logger.Logger) *InterceptorManager {
	return &InterceptorManager{
		logger: logger,
	}
}

func (i *InterceptorManager) Logger(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	now := time.Now()

	meta, _ := metadata.FromIncomingContext(ctx)
	reply, err := handler(ctx, req)

	i.logger.Infof("Method: %v, Time: %v, Metadata: %v, Err: %v", info.FullMethod, time.Since(now), meta, err)

	return reply, err
}

// func (i *InterceptorManager) Metrics(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
// 	start := time.Now()
// 	resp, err := handler(ctx, req)
// 	var status = http.StatusOK
// 	if err != nil {

// 	}

// 	i.Metrics(.)
// }
