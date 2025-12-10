package response

const (
	ErrCodeSuccess      = 201
	ErrCodeParamInvalid = 403
	ErrCodeNotFound     = 404
)

var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "param error",
}
