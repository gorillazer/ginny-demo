package client

import (
	"context"
	"fmt"

	consul "github.com/goriller/ginny-consul"
	"github.com/goriller/ginny-demo/internal/config"
	"github.com/goriller/ginny/client"
	"github.com/goriller/ginny/logger"

	pb "github.com/goriller/ginny-demo/api/proto"
)

// NewExampleClient
func NewExampleClient(ctx context.Context, config *config.Config, consul *consul.Client) (pb.SayClient, error) {
	endpoint := config.Client["example"]
	if endpoint == "" {
		return nil, fmt.Errorf("grpc endpoint is undefined")
	}
	cli, err := client.NewGrpcClient(ctx, endpoint, pb.NewSayClient,
		client.WithGrpcResolver(func(ctx context.Context, service, tag string) (addr string, err error) {
			return consul.Resolver(ctx, service, tag)
		}))
	if err != nil {
		logger.Action("NewGrpcCli").Error(fmt.Sprintf("%s: %s", endpoint, err.Error()))
		return nil, err
	}
	if data, ok := cli.(pb.SayClient); ok {
		return data, nil
	}
	return nil, fmt.Errorf("get grpc client error")
}
