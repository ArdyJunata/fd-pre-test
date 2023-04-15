package response

type ResponseSuccess struct {
	StatusCode int         `json:"-"`
	Code       string      `json:"-"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func Success(message string) *ResponseSuccess {
	return &ResponseSuccess{
		Message: message,
	}
}

func (c *ResponseSuccess) WithData(data interface{}) *ResponseSuccess {
	c.Data = data
	return c
}

func (c *ResponseSuccess) WithMessage(msg string) *ResponseSuccess {
	c.Message = msg

	return c
}
