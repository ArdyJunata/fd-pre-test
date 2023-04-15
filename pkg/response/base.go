package response

type AdditionalInfo struct {
	Usecase string `json:"usecase"`
	Info    string `json:"info"`
}

const (
	MSG_FETCH_USER_FAILED  string = "fetch user failed"
	MSG_FETCH_USER_SUCCESS string = "fetch user success"

	MSG_FIND_ONE_USER_FAILED  string = "find one user failed"
	MSG_FIND_ONE_USER_SUCCESS string = "find one user success"
)
