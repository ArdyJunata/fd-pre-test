package params

import (
	"fd-test/application/model"
	"time"
)

type FetchUserResponse struct {
	ID        int    `json:"id" binding:"required"`
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Avatar    string `json:"avatar" binding:"required"`
}

func (c FetchUserResponse) ParseToModel() model.User {
	return model.User{
		ID:        c.ID,
		Email:     c.Email,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Avatar:    c.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type UserResponse struct {
	ID        int        `json:"id"`
	Email     string     `json:"email"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Avatar    string     `json:"avatar"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
