package response

const (
	ErrorCodeSuccess      = 20001
	ErrorCodeParamInvalid = 20002

	ErrUnauthorized = 30001
)

var msg = map[int]string{
	ErrorCodeSuccess:      "Success",
	ErrorCodeParamInvalid: "Param invalid",
	ErrUnauthorized:       "Unauthorized",
}
