// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"context"
	"github.com/goriller/ginny"
	"github.com/goriller/ginny-broker"
	"github.com/goriller/ginny-consul"
	"github.com/goriller/ginny-demo/internal/cache"
	config2 "github.com/goriller/ginny-demo/internal/config"
	"github.com/goriller/ginny-demo/internal/repo"
	"github.com/goriller/ginny-demo/internal/service"
	"github.com/goriller/ginny-demo/internal/task"
	"github.com/goriller/ginny-jaeger"
	"github.com/goriller/ginny-mysql"
	"github.com/goriller/ginny-redis"
	"github.com/goriller/ginny/config"
	"github.com/goriller/ginny/logger"
	"github.com/goriller/ginny/server"
	"github.com/opentracing/opentracing-go"
)

import (
	_ "go.uber.org/automaxprocs/maxprocs"
)

// Injectors from app.go:

// NewApp
func NewApp(ctx context.Context) (*ginny.Application, error) {
	viper, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	option, err := ginny.NewOption(viper)
	if err != nil {
		return nil, err
	}
	zapLogger := logger.Default()
	configConfig, err := config2.NewConfig(viper)
	if err != nil {
		return nil, err
	}
	redisConfig, err := redis.NewConfig(viper)
	if err != nil {
		return nil, err
	}
	redisRedis, err := redis.NewRedis(ctx, redisConfig, zapLogger)
	if err != nil {
		return nil, err
	}
	redisCache := cache.NewRedisCache(redisRedis)
	brokerConfig, err := broker.NewConfiguration(viper)
	if err != nil {
		return nil, err
	}
	brokerBroker, err := broker.NewBroker(ctx, zapLogger, brokerConfig)
	if err != nil {
		return nil, err
	}
	taskTask := &task.Task{
		Broker: brokerBroker,
	}
	mysqlConfig, err := mysql.NewConfig(viper)
	if err != nil {
		return nil, err
	}
	sqlBuilder, err := mysql.NewSqlBuilder(ctx, mysqlConfig, zapLogger)
	if err != nil {
		return nil, err
	}
	userRepo := repo.NewUserRepo(configConfig, sqlBuilder, redisCache)
	serviceService, err := service.NewService(ctx, configConfig, redisCache, taskTask, userRepo)
	if err != nil {
		return nil, err
	}
	registrarFunc := service.RegisterService(ctx, serviceService)
	apiConfig, err := consul.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	client, err := consul.NewClient(ctx, apiConfig)
	if err != nil {
		return nil, err
	}
	configuration, err := jaeger.NewConfiguration(viper)
	if err != nil {
		return nil, err
	}
	tracer, err := jaeger.NewJaegerTracer(ctx, configuration)
	if err != nil {
		return nil, err
	}
	v := serverOption(client, tracer)
	application, err := ginny.NewApp(ctx, option, zapLogger, registrarFunc, v...)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// app.go:

func serverOption(consul2 *consul.Client,
	tracer opentracing.Tracer,
) (opts []server.Option) {
	opts = append(opts, server.WithDiscover(consul2))
	opts = append(opts, server.WithTracer(tracer))
	return
}
