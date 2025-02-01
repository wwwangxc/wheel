package errorx

import (
	"errors"
	"fmt"

	grpccodes "google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
)

var errUnknown = New(codeUnknown, "Unknown Error", WithReason("未知错误"))

// Error custom
type Error interface {
	Code() uint32
	Message() string
	Reason() string
	String() string
	Error() string
	GRPCStatus() *grpcstatus.Status
}

type errorImpl struct {
	code    ErrCode
	message string
	reason  string
}

// New error
func New(code ErrCode, message string, opts ...option) Error {
	o := newOptions(opts...)
	return &errorImpl{
		code:    code,
		message: message,
		reason:  o.reason,
	}
}

// FromError convert error to errorx.Error
func FromError(err error) (Error, bool) {
	var e *errorImpl
	return e, errors.As(err, &e)
}

func (s *errorImpl) Code() uint32 {
	if s == nil {
		return errUnknown.Code()
	}

	return uint32(s.code)
}

func (s *errorImpl) Message() string {
	if s == nil {
		return errUnknown.Message()
	}

	return s.message
}

func (s *errorImpl) Reason() string {
	if s == nil {
		return errUnknown.Reason()
	}

	return s.reason
}

func (s *errorImpl) String() string {
	if s == nil {
		return errUnknown.String()
	}

	return fmt.Sprintf("%d %s", s.code, s.message)
}

func (s *errorImpl) Error() string {
	if s == nil {
		return errUnknown.Error()
	}

	return s.String()
}

func (s *errorImpl) GRPCStatus() *grpcstatus.Status {
	if s == nil {
		return errUnknown.GRPCStatus()
	}

	return grpcstatus.New(grpccodes.Code(s.code), s.message)
}
