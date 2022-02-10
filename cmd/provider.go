// +build wireinject

package main

import (
	"github.com/gorillazer/ginny-demo/internal/handlers"
	"github.com/gorillazer/ginny-demo/internal/repositories"
	rpc "github.com/gorillazer/ginny-demo/internal/rpc"
	"github.com/gorillazer/ginny-demo/internal/services"
	// CMD_IMPORT 锚点请勿删除! Do not delete this line!

	"github.com/google/wire"
	"github.com/gorillazer/ginny"
	config "github.com/gorillazer/ginny-config"
	consul "github.com/gorillazer/ginny-consul"
	jaeger "github.com/gorillazer/ginny-jaeger"
	log "github.com/gorillazer/ginny-log"
	// grpc "github.com/gorillazer/ginny-serve/grpc"
	http "github.com/gorillazer/ginny-serve/http"
)

// Create http/grpc Serve
func newServe(
	hs *http.Server,
	cli *consul.Client,
	// gs *grpc.Server,
	// CMD_SERVEPARAM 锚点请勿删除! Do not delete this line!

) ([]ginny.Serve, error) {
	return []ginny.Serve{
		ginny.HttpServe(hs),
		// ginny.GrpcServeWithConsul(gs, cli),
		// CMD_SERVEFUNC 锚点请勿删除! Do not delete this line!

	}, nil
}

// appProvider
var appProvider = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	jaeger.ProviderSet,
	newServe, ginny.AppProviderSet)

// CreateApp
func CreateApp(name string) (*ginny.Application, error) {
	panic(wire.Build(wire.NewSet(

		// grpc
		rpc.ProviderSet,

		handlers.ProviderSet,
		repositories.ProviderSet,
		services.ProviderSet,
		// CMD_PROVIDERSET 锚点请勿删除! Do not delete this line!

		appProvider,
	)))
}
