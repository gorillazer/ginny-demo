package rpc

import (
	"github.com/google/wire"
	consul "github.com/gorillazer/ginny-consul"
	rpc_client "github.com/gorillazer/ginny-demo/internal/rpc/client"
	"github.com/gorillazer/ginny-serve/grpc"
	// rpc_server "github.com/gorillazer/ginny-demo/internal/rpc/server"
)

// ProviderSet
var ProviderSet = wire.NewSet(
	grpc.ProviderSet,
	consul.ProviderSet,
	// rpc_server.ProviderSet,
	rpc_client.ProviderSet,
)
