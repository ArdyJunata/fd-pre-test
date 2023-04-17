package http

import "github.com/gin-gonic/gin"

type MidValidationAuth interface {
	Auth() gin.HandlerFunc
}
