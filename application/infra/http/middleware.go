package infrahttp

import (
	"errors"
	"fd-test/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Middleware struct{}

func NewMiddleware() Middleware {
	return Middleware{}
}

func (m Middleware) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")

		if auth != "3cdcnTiBsl" {
			resp := response.Error(errors.New("unauthorized")).WithMessage("invalid Authorization").WithStatusCode(http.StatusUnauthorized)
			ctx.AbortWithStatusJSON(resp.StatusCode, resp)
			return
		}

		ctx.Next()
	}
}
