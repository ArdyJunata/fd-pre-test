package params

type FetchUserRequest struct {
	Page    int `form:"page" validate:"required"`
	PerPage int `form:"per_page"`
}