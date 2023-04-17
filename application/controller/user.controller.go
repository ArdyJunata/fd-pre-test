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
	base.GET("/:id", u.FindUserById)
	base.GET("", u.FindAllUser)
	base.POST("", u.CreateUser)
	base.PUT("/:id", u.UpdateUserById)
}

func (u userController) FetchUser(ctx *gin.Context) {
	var req params.FetchUserRequest

	err := ctx.ShouldBindQuery(&req)
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

func (u userController) FindUserById(ctx *gin.Context) {
	var req params.GetUserByIdRequest

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		resp := response.Error(err).WithMessage(response.MSG_FIND_ONE_USER_FAILED).WithInfo("FetchUser", "try to bind http request body")
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	resp, respErr := u.svc.FindUserById(ctx, req)
	if !respErr.IsNoError {
		ctx.AbortWithStatusJSON(respErr.StatusCode, respErr)
		return
	}

	response := response.Success(response.MSG_FIND_ONE_USER_SUCCESS).WithData(resp)

	ctx.JSON(response.StatusCode, response)
}

func (u userController) FindAllUser(ctx *gin.Context) {
	resp, respErr := u.svc.FindAllUser(ctx)
	if !respErr.IsNoError {
		ctx.AbortWithStatusJSON(respErr.StatusCode, respErr)
		return
	}

	response := response.Success(response.MSG_FIND_ALL_USER_SUCCESS).WithData(resp)

	ctx.JSON(response.StatusCode, response)
}

func (u userController) CreateUser(ctx *gin.Context) {
	var req params.CreateUserRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		resp := response.Error(err).WithMessage(response.MSG_CREATE_USER_FAILED).WithInfo("CreateUser", "try to parse struct")
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	respErr := u.svc.CreateUser(ctx, req)
	if !respErr.IsNoError {
		ctx.AbortWithStatusJSON(respErr.StatusCode, respErr)
		return
	}

	resp := response.Success(response.MSG_CREATE_USER_SUCCESS)

	ctx.JSON(resp.StatusCode, resp)
}

func (u userController) UpdateUserById(ctx *gin.Context) {
	var req params.UpdateUserRequest

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		resp := response.Error(err).WithMessage(response.MSG_CREATE_USER_FAILED).WithInfo("UpdateUser", "try to parse struct")
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		resp := response.Error(err).WithMessage(response.MSG_CREATE_USER_FAILED).WithInfo("UpdateUser", "try to parse struct")
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	respErr := u.svc.UpdateUserById(ctx, req)
	if !respErr.IsNoError {
		ctx.AbortWithStatusJSON(respErr.StatusCode, respErr)
		return
	}

	resp := response.Success(response.MSG_UPDATE_USER_SUCCESS)

	ctx.JSON(resp.StatusCode, resp)
}
