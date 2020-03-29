package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrResponse struct{
	// Http Status Code
	HttpSC int
	Err err
}

var (
	// 校验 request
	ErrorRequestBodyParseFailed = ErrResponse{400, Err{"Request body is not correct", "001"}}

	// 校验 user
	ErrorNotAuthUser = ErrResponse{401, Err{"User authentication failed.", "002"}}
)