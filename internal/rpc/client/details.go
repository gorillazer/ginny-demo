package client

import (
	"github.com/gorillazer/ginny-demo/api/proto"

	consul "github.com/gorillazer/ginny-consul"
	"github.com/gorillazer/ginny-serve/grpc"
	"github.com/pkg/errors"
)

func NewDetailsClient(
	client *grpc.Client,
	consul *consul.Client,
) (proto.DetailsClient, error) {
	conn, err := client.Dial("Details", grpc.WithTarget("10.95.19.67:3000"))
	// conn, err := client.Dial("Details", grpc.WithConsulConfig(consul.Config))
	if err != nil {
		return nil, errors.Wrap(err, "detail client dial error")
	}
	c := proto.NewDetailsClient(conn)

	return c, nil
}
