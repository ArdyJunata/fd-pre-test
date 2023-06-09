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

	MSG_FIND_ALL_USER_FAILED  string = "find all user failed"
	MSG_FIND_ALL_USER_SUCCESS string = "find all user success"

	MSG_CREATE_USER_FAILED  string = "create user failed"
	MSG_CREATE_USER_SUCCESS string = "create user success"

	MSG_UPDATE_USER_FAILED  string = "update user failed"
	MSG_UPDATE_USER_SUCCESS string = "update user success"

	MSG_DELETE_USER_FAILED  string = "delete user failed"
	MSG_DELETE_USER_SUCCESS string = "delete user success"
)
