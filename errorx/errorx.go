// Package errorx is the custom error
package errorx

var (
	// ErrServer server error
	ErrServer = New(codeServerError, "Server Error", WithReason("服务内部错误"))
	//ErrInvalidRequest invalid request error
	ErrInvalidRequest = New(codeInvalidRequest, "Invalid Request", WithReason("无效请求"))
	// ErrInvalidArguments invalid arguments error
	ErrInvalidArguments = New(codeInvalidArguments, "Invalid Arguments", WithReason("无效参数"))
)
