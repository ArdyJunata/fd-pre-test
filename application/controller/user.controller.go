package controller

import (
	"fd-test/application/params"
	"fd-test/application/service"
	"fd-test/pkg/response"

	"github.com/gin-gonic/gin"
)

type userController struct {
	svc service.UserService
}

func NewUserController(svc service.UserService) userController {
	return userController{
		svc: svc,
	}
}

func (u userController) RegisterRoute(route *gin.Engine) {
	base := route.Group("/user")
	base.GET("/fetch", u.FetchUser)
}

func (u userController) FetchUser(ctx *gin.Context) {
	var req params.FetchUserRequest

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		resp := response.Error(err).WithMessage(response.MSG_FETCH_USER_FAILED).WithInfo("FetchUser", "try to bind http request body")
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	resp, respErr := u.svc.FetchUser(ctx, req)
	if !respErr.IsNoError {
		ctx.AbortWithStatusJSON(respErr.StatusCode, respErr)
		return
	}

	response := response.Success(response.MSG_FETCH_USER_SUCCESS).WithData(resp)

	ctx.JSON(response.StatusCode, response)
}
