package errorx

type ErrCode uint32

const codeUnknown ErrCode = 0

const (
	codeServerError      ErrCode = 500_00
	codeInvalidRequest   ErrCode = 500_01
	codeInvalidArguments ErrCode = 500_02
)
