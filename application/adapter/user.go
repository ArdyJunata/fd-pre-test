package adapter

import (
	"encoding/json"
	"fd-test/application/params"
	"fd-test/helper"
	"fmt"
)

type UserAdapter interface {
	FetchUser(req params.FetchUserRequest) ([]params.FetchUserResponse, error)
}

type userAdapter struct{}

type fetchUserResponse struct {
	Page    int                        `json:"page"`
	PerPage int                        `json:"per_page"`
	Total   int                        `json:"total"`
	Data    []params.FetchUserResponse `json:"data"`
}

// FetchUser implements UserAdapter
func (u userAdapter) FetchUser(req params.FetchUserRequest) ([]params.FetchUserResponse, error) {
	url := "https://reqres.in/api/users?page=" + fmt.Sprintf("%d", req.Page)

	if req.PerPage != 0 {
		url = url + fmt.Sprintf("&per_page=%d", req.PerPage)
	}

	bodyResp, err := helper.RequestGet(url)
	if err != nil {
		return nil, err
	}

	resp := fetchUserResponse{}

	err = json.Unmarshal(bodyResp, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func NewUserAdapter() UserAdapter {
	return userAdapter{}
}
