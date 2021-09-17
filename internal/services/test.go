package services

import (
	"context"

	"github.com/gorillazer/ginny-demo/internal/constants"
	"github.com/gorillazer/ginny-demo/internal/repositories"

	"github.com/google/wire"
	"github.com/gorillazer/ginny/errs"
	"go.uber.org/zap"
)

// TestServiceProvider
var TestServiceProvider = wire.NewSet(NewTestService, wire.Bind(new(ITestService), new(*TestService)))

// ITestService
type ITestService interface {
	Get(ctx context.Context, Id uint64) (string, error)
}

// TestService
type TestService struct {
	logger *zap.Logger
	// Introduce new dependencies here, exp:
	userRepository *repositories.UserRepository
}

// NewTestService
func NewTestService(
	logger *zap.Logger,
	userRepository *repositories.UserRepository,
) *TestService {
	return &TestService{
		logger:         logger.With(zap.String("type", "Hello")),
		userRepository: userRepository,
	}
}

//
func (p *TestService) Get(ctx context.Context, Id uint64) (string, error) {
	user, err := p.userRepository.GetUser(ctx)
	if err != nil {
		p.logger.Error("Get", zap.Error(err))
		return "", errs.New(constants.NOT_FOUND, constants.GetErrMsg(constants.NOT_FOUND))
	}

	return user.Name, nil
}
