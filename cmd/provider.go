// +build wireinject

package main

import (
	"github.com/gorillazer/ginny-demo/internal/handlers"
	"github.com/gorillazer/ginny-demo/internal/repositories"
	rpc "github.com/gorillazer/ginny-demo/internal/rpc"
	rpc_client "github.com/gorillazer/ginny-demo/internal/rpc/client"
	rpc_server "github.com/gorillazer/ginny-demo/internal/rpc/server"
	"github.com/gorillazer/ginny-demo/internal/services"

	"github.com/google/wire"
	"github.com/gorillazer/ginny"
	config "github.com/gorillazer/ginny-config"
	jaeger "github.com/gorillazer/ginny-jaeger"
	log "github.com/gorillazer/ginny-log"

	consul "github.com/gorillazer/ginny-consul"
	grpc "github.com/gorillazer/ginny-serve/grpc"
	http "github.com/gorillazer/ginny-serve/http"
)

var appProvider = wire.NewSet(newServe, ginny.AppProviderSet)

// Create http/grpc Serve
func newServe(
	hs *http.Server,
	cli *consul.Client,
	gs *grpc.Server,
) ([]ginny.Serve, error) {
	return []ginny.Serve{
		ginny.HttpServe(hs),
		ginny.GrpcServeWithConsul(gs, cli),
	}, nil
}

// CreateApp
func CreateApp(name string) (*ginny.Application, error) {
	panic(wire.Build(wire.NewSet(
		log.ProviderSet,
		config.ProviderSet,
		jaeger.ProviderSet,

		handlers.ProviderSet,
		// grpc
		rpc.ProviderSet,
		rpc_server.ProviderSet,
		rpc_client.ProviderSet,

		services.ProviderSet,
		repositories.ProviderSet,
		appProvider,
	)))
}
