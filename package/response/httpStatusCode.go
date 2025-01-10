package response

const (
	ErrorCodeSuccess      = 20001
	ErrorCodeParamInvalid = 20002
)

var msg = map[int]string{
	ErrorCodeSuccess:      "success",
	ErrorCodeParamInvalid: "param invalid",
}
