package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/gorillazer/ginny-serve/http"
	"github.com/gorillazer/ginny/res"
)

func CreateInitHandlerFn(
	test *TestHandler,
	// HANDLE 锚点请勿删除! Do not delete this line!
) http.InitHandlers {
	return func(r *gin.Engine) {
		// 在此定义路由规则 Define routing rules here, exp:
		r.GET("/", res.Wrapper(test.Get))
		r.GET("/test/:id", res.Wrapper(test.Test))
		r.GET("/getrpc", res.Wrapper(test.GetRPC))
	}
}

var ProviderSet = wire.NewSet(
	http.ProviderSet,
	TestHandlerProvider,
	// HANDLE_PROVIDER 锚点请勿删除! Do not delete this line!
	CreateInitHandlerFn,
)
