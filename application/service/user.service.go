package service

import (
	"context"
	"fd-test/application/adapter"
	"fd-test/application/params"
	"fd-test/application/repository"
	"fd-test/pkg/response"
	"strings"
)

type UserService interface {
	FetchUser(ctx context.Context, req params.FetchUserRequest) ([]params.FetchUserResponse, response.ResponseError)
	FindUserById(ctx context.Context, req params.GetUserByIdRequest) (params.UserResponse, response.ResponseError)
}

type userService struct {
	userRepo    repository.UserRepository
	userAdapter adapter.UserAdapter
}

// FindUserById implements UserService
func (u userService) FindUserById(ctx context.Context, req params.GetUserByIdRequest) (params.UserResponse, response.ResponseError) {
	user, err := u.userRepo.FindUserById(ctx, req.ID)
	if err != nil {
		return params.UserResponse{}, *response.Error(err).WithMessage(response.MSG_FIND_ONE_USER_FAILED).WithInfo("FindOneUser", "try to query to db")
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
				return nil, *response.Error(err).WithMessage(response.MSG_FETCH_USER_FAILED).WithInfo("FetchUser", "try to insert to users table")
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
