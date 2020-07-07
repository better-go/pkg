package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// convert:
func pprofHandler(h http.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
