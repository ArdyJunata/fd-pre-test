package params

import (
	"fd-test/application/model"
	"time"
)

type CreateUserRequest struct {
	Email     string `json:"email" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Avatar    string `json:"avatar" validate:"required"`
}

func (c CreateUserRequest) ParseToModel() model.User {
	return model.User{
		Email:     c.Email,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Avatar:    c.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type FetchUserRequest struct {
	Page    int `form:"page" binding:"required"`
	PerPage int `form:"per_page"`
}

type GetUserByIdRequest struct {
	ID int `uri:"id"`
}

type UpdateUserRequest struct {
	ID        int    `uri:"id"`
	Email     string `json:"email" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Avatar    string `json:"avatar" validate:"required"`
}

func (c UpdateUserRequest) ParseToModel() model.User {
	return model.User{
		ID:        c.ID,
		Email:     c.Email,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Avatar:    c.Avatar,
		UpdatedAt: time.Now(),
	}
}

type DeleteUserRequest struct {
	ID int `uri:"id"`
}
