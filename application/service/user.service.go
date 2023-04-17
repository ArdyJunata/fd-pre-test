package service

import (
	"context"
	"fd-test/application/adapter"
	"fd-test/application/params"
	"fd-test/application/repository"
	"fd-test/pkg/response"
	"net/http"
	"strings"
)

type UserService interface {
	FetchUser(ctx context.Context, req params.FetchUserRequest) ([]params.FetchUserResponse, response.ResponseError)
	FindUserById(ctx context.Context, req params.GetUserByIdRequest) (params.UserResponse, response.ResponseError)
	FindAllUser(ctx context.Context) ([]params.UserResponse, response.ResponseError)
	CreateUser(ctx context.Context, req params.CreateUserRequest) response.ResponseError
	UpdateUserById(ctx context.Context, req params.UpdateUserRequest) response.ResponseError
}

type userService struct {
	userRepo    repository.UserRepository
	userAdapter adapter.UserAdapter
}

// UpdateUserById implements UserService
func (u userService) UpdateUserById(ctx context.Context, req params.UpdateUserRequest) response.ResponseError {
	payload := req.ParseToModel()

	err := u.userRepo.UpdateUserById(ctx, payload)
	if err != nil {
		return *response.Error(err).WithMessage(response.MSG_UPDATE_USER_FAILED).WithInfo("CreateUser", "try to insert query to db").WithStatusCode(http.StatusInternalServerError)
	}

	return response.NotError()
}

// CreateUser implements UserService
func (u userService) CreateUser(ctx context.Context, req params.CreateUserRequest) response.ResponseError {
	payload := req.ParseToModel()

	err := u.userRepo.CreateUser(ctx, payload)
	if err != nil {
		return *response.Error(err).WithMessage(response.MSG_CREATE_USER_FAILED).WithInfo("CreateUser", "try to insert query to db").WithStatusCode(http.StatusInternalServerError)
	}

	return response.NotError()
}

// FindAllUser implements UserService
func (u userService) FindAllUser(ctx context.Context) ([]params.UserResponse, response.ResponseError) {
	resp := []params.UserResponse{}

	users, err := u.userRepo.FindAllUser(ctx)
	if err != nil {
		return resp, *response.Error(err).WithMessage(response.MSG_FIND_ALL_USER_FAILED).WithInfo("FindAllUser", "try to query to db").WithStatusCode(http.StatusInternalServerError)
	}

	for _, v := range users {
		user := params.UserResponse{
			ID:        v.ID,
			Email:     v.Email,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Avatar:    v.Avatar,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt,
		}

		resp = append(resp, user)
	}

	return resp, response.NotError()
}

// FindUserById implements UserService
func (u userService) FindUserById(ctx context.Context, req params.GetUserByIdRequest) (params.UserResponse, response.ResponseError) {
	user, err := u.userRepo.FindUserById(ctx, req.ID)
	if err != nil {
		return params.UserResponse{}, *response.Error(err).WithMessage(response.MSG_FIND_ONE_USER_FAILED).WithInfo("FindOneUser", "try to query to db").WithStatusCode(http.StatusInternalServerError)
	}

	resp := params.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	return resp, response.NotError()
}

// FetchUser implements UserService
func (u userService) FetchUser(ctx context.Context, req params.FetchUserRequest) ([]params.FetchUserResponse, response.ResponseError) {
	users, err := u.userAdapter.FetchUser(req)
	if err != nil {
		return nil, *response.Error(err).WithMessage(response.MSG_FETCH_USER_FAILED).WithInfo("FetchUser", "try to fetch from third party")
	}

	for _, v := range users {
		user := v.ParseToModel()

		err := u.userRepo.FetchUser(ctx, user)
		if err != nil {
			if !strings.Contains(err.Error(), "duplicate") {
				return nil, *response.Error(err).WithMessage(response.MSG_FETCH_USER_FAILED).WithInfo("FetchUser", "try to insert to users table").WithStatusCode(http.StatusInternalServerError)
			}
		}

	}

	return users, response.NotError()
}

func NewUserService(userRepo repository.UserRepository, userAdapter adapter.UserAdapter) UserService {
	return userService{
		userRepo:    userRepo,
		userAdapter: userAdapter,
	}
}
