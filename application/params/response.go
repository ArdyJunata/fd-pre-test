package params

import (
	"fd-test/application/model"
	"time"
)

type FetchUserResponse struct {
	ID        int    `json:"id" validate:"required"`
	Email     string `json:"email" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Avatar    string `json:"avatar" validate:"required"`
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
