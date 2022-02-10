package server

import (
	"context"
	"time"

	"github.com/gorillazer/ginny-demo/api/proto"
	"github.com/gorillazer/ginny-demo/internal/constants"
	"github.com/gorillazer/ginny-demo/internal/services"
	"github.com/pkg/errors"

	"github.com/google/wire"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.uber.org/zap"
)

// DetailsServerProvider
var DetailsServerProvider = wire.NewSet(NewDetailsServer, wire.Bind(new(IDetailsServer), new(*DetailsServer)))

// IDetailsServer
type IDetailsServer interface {
	Get(ctx context.Context, req *proto.GetReq) (*proto.GetRes, error)
}

// DetailsServer
type DetailsServer struct {
	logger      *zap.Logger
	testService *services.TestService
}

// NewDetailsServer
func NewDetailsServer(
	logger *zap.Logger,
	testService *services.TestService,
) (*DetailsServer, error) {
	return &DetailsServer{
		logger:      logger.With(zap.String("type", "DetailsServer")),
		testService: testService,
	}, nil
}

func (s *DetailsServer) Get(ctx context.Context, req *proto.GetReq) (*proto.GetRes, error) {
	if req == nil {
		return nil, errors.New(constants.GetErrMsg(constants.PARAMS_INVALID))
	}
	p, err := s.testService.Get(ctx, req.Id)
	if err != nil {
		s.logger.Error("Get", zap.Error(err))
		return nil, err
	}
	ct := timestamppb.New(time.Time{})
	resp := &proto.GetRes{
		Id:          req.Id,
		Name:        p,
		CreatedTime: ct,
	}

	return resp, nil
}
