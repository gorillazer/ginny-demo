package service

import (
	"context"

	"github.com/google/wire"
	"github.com/goriller/ginny"
	broker "github.com/goriller/ginny-broker"
	_ "github.com/goriller/ginny-broker/kafka"
	pb "github.com/goriller/ginny-demo/api/proto"
	"github.com/goriller/ginny-demo/internal/config"
	"github.com/goriller/ginny-demo/internal/repo"
	"github.com/goriller/ginny-demo/internal/task"
	"github.com/goriller/ginny/errs"
	"github.com/goriller/ginny/logger"
	"go.uber.org/zap"
)

// ProviderSet
var ProviderSet = wire.NewSet(task.TaskProvider, NewService, RegisterService)

// Service the instance for grpc proto.
type Service struct {
	pb.UnimplementedSayServer
	config *config.Config
	// Introduce new dependencies here, exp:
	task           *task.Task
	userRepository *repo.UserRepo
}

// NewService new service that implement hello
func NewService(
	ctx context.Context,
	config *config.Config,
	task *task.Task,
	userRepository *repo.UserRepo,
) (*Service, error) {
	errs.RegisterErrorCodes(pb.ErrorCode_name)

	return &Service{
		config:         config,
		task:           task,
		userRepository: userRepository,
	}, nil
}

// RegisterService
func RegisterService(ctx context.Context, sev *Service) ginny.RegistrarFunc {
	return func(app *ginny.Application) error {
		// 注册gRPC服务
		app.Server.RegisterService(&pb.Say_ServiceDesc, sev)
		if app.Option.HttpAddr != "" {
			// 注册http服务
			if err := pb.RegisterSayHandlerServer(ctx, app.Server.ServeMux(), sev); err != nil {
				return err
			}
		}
		// 注册消息队列监听
		err := RegisterTask(ctx, sev)
		if err != nil {
			return err
		}

		return nil
	}
}

// RegisterTask 注册消息队列监听
func RegisterTask(ctx context.Context, sev *Service) error {
	// task1
	err := sev.task.Subscribe([]string{sev.config.Broker.Topic}, func(p broker.Publication) error {
		logger.Action("process").Debug("", zap.String("topic", p.Topic()))
		logger.Action("process").Debug("", zap.String("message", string(p.Message().Body)))
		return p.Ack()
	})
	if err != nil {
		return err
	}
	return nil
}
