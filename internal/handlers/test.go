package handlers

import (
	"github.com/gorillazer/ginny-demo/api/proto"
	"github.com/gorillazer/ginny-demo/internal/constants"
	"github.com/gorillazer/ginny-demo/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	// "github.com/gorillazer/ginny/errs"

	util "github.com/gorillazer/ginny-util"
	"github.com/gorillazer/ginny/errs"
	"github.com/gorillazer/ginny/res"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// TestHandlerProvider
var TestHandlerProvider = wire.NewSet(NewTestHandler, wire.Bind(new(ITestHandler), new(*TestHandler)))

// ITestHandler
type ITestHandler interface {
	Get(c *gin.Context) (*res.Response, error)
	GetRPC(c *gin.Context) (*res.Response, error)
}

// TestHandler
type TestHandler struct {
	v      *viper.Viper
	logger *zap.Logger
	// Introduce new dependencies here, exp:
	testService  *services.TestService
	detailClient proto.DetailsClient
}

// NewTestHandler
func NewTestHandler(
	v *viper.Viper,
	logger *zap.Logger,
	testService *services.TestService,
	detailClient proto.DetailsClient,
) *TestHandler {
	return &TestHandler{
		v:            v,
		logger:       logger.With(zap.String("type", "TestHandler")),
		testService:  testService,
		detailClient: detailClient,
	}
}

// Get
func (t *TestHandler) Get(c *gin.Context) (*res.Response, error) {
	t.logger.Debug("TestHandler", zap.Any("TestHandler.Get", c.Params))
	id := c.Query("id")
	name, err := t.testService.Get(c, util.Uint64(id))
	if err != nil {
		t.logger.Error("", zap.Error(err))
		return nil, errs.New(constants.INTERNAL_ERR, constants.GetErrMsg(constants.INTERNAL_ERR))
	}
	return res.Success(name), nil
}

// GetRPC
func (t *TestHandler) GetRPC(c *gin.Context) (*res.Response, error) {
	req := &proto.GetReq{
		Id: 1,
	}
	t.logger.Info(t.v.GetString("consul.address"))
	p, err := t.detailClient.Get(c, req)
	if err != nil {
		t.logger.Error("GetRPC", zap.Error(err))
		return res.Fail(errs.New(constants.NOT_FOUND, constants.GetErrMsg(constants.NOT_FOUND))), nil
	}
	return res.Success(p), nil
}
