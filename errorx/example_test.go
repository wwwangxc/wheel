package errorx_test

import (
	"errors"
	"fmt"

	"github.com/wwwangxc/wheel/errorx"
)

// CodeServerError server error code
const CodeServerError errorx.ErrCode = 500_00

// ErrServer internal server error
var ErrServer = errorx.New(CodeServerError, "Internal Server Error", errorx.WithReason("This is the server error"))

func Example() {
	fmt.Printf("Code: %d\n", ErrServer.Code())       // Code: 50000
	fmt.Printf("Message: %s\n", ErrServer.Message()) // Message: Internal Server Error
	fmt.Printf("Reason: %s\n", ErrServer.Reason())   // Reason: This is the server error
	fmt.Printf("Error: %s\n", ErrServer.Error())     // Error: 50000 Internal Server Error
	fmt.Printf("String: %s\n", ErrServer)            // String: 50000 Internal Server Error
	fmt.Println()

	e := errors.New("custom error")
	err, ok := errorx.FromError(e)
	fmt.Println(err) // <nil>
	fmt.Println(ok)  // false
	fmt.Println()

	e = fmt.Errorf("wraped error: %w", ErrServer)
	err, ok = errorx.FromError(e)
	fmt.Println(err) // 50000 Internal Server Error
	fmt.Println(ok)  // true

	// Output:
	// Code: 50000
	// Message: Internal Server Error
	// Reason: This is the server error
	// Error: 50000 Internal Server Error
	// String: 50000 Internal Server Error
	//
	// <nil>
	// false
	//
	// 50000 Internal Server Error
	// true
}
