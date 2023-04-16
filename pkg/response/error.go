package response

import (
	"net/http"
)

type ResponseError struct {
	StatusCode     int             `json:"-"`
	Message        string          `json:"message"`
	Error          string          `json:"error,omitempty"`
	Data           interface{}     `json:"data,omitempty"`
	AdditionalInfo *AdditionalInfo `json:"additional_info,omitempty"`
	IsNoError      bool            `json:"-"`
}

func Error(err error) *ResponseError {
	return &ResponseError{
		StatusCode: http.StatusBadRequest,
		Error:      err.Error(),
	}
}

func (c *ResponseError) WithError(err string) *ResponseError {
	c.Error = err

	return c
}

func (c *ResponseError) WithMessage(msg string) *ResponseError {
	c.Message = msg

	return c
}

func (c *ResponseError) WithStatusCode(code int) *ResponseError {
	c.StatusCode = code

	return c
}

func (c *ResponseError) WithInfo(usecase, info string) *ResponseError {
	c.AdditionalInfo = &AdditionalInfo{
		Usecase: usecase,
		Info:    info,
	}

	return c
}

func NotError() ResponseError {
	return ResponseError{
		IsNoError: true,
	}
}
