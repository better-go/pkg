package gin

import (
	"net/http"

	"github.com/better-go/pkg/log"
	"github.com/gin-gonic/gin"
)

/*

路由自动校验入参 + 格式化返回值

usage:

	func userRegister(ctx *gin.Context) {
		var req api.UserRegisterReq

		ApiHandlerWrap(ctx, &req, func(ctx *gin.Context, in interface{}) (out interface{}, err error) {
			// assert:
			r, ok := in.(*api.UserRegisterReq)
			if !ok {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			}
			return gs.Register(ctx, r)
		})
	}


*/

// ApiHandlerFunc api logic func
type ApiHandlerFunc func(ctx *gin.Context, in interface{}) (out interface{}, err error)

// ApiHandlerWrap 路由自动校验入参 + 格式化返回值
func ApiHandlerWrap(ctx *gin.Context, req interface{}, handlerFn ApiHandlerFunc) {
	log.Debugf("http api request entry: req=%+v", req)
	//value := reflect.ValueOf(req)
	//log.Debugf("req type before bind: %+v, type:%v", req, value.Type())

	// validate req:
	if err := ctx.ShouldBind(req); err != nil {
		log.Error("invalid request params: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//value = reflect.ValueOf(req)
	//log.Debugf("req type after bind: %+v, type:%v", req, value.Type())

	//
	// do api handler
	//
	resp, err := handlerFn(ctx, req)
	log.Debugf("http api request done: resp=%+v, err=%v", resp, err)

	// resp:
	if err != nil {
		log.Error("user register: Register error")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ResponseData{
		Code:    200,
		Message: "ok",
		Data:    resp,
	})
}
