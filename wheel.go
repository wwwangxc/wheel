package wheel

import "github.com/wwwangxc/wheel/reflectx"

// ValueOrDefault when the value is the zero value of the type, return the specified default value.
func ValueOrDefault[T any](val, defaultVal T) T {
	if reflectx.IsZeroValue(val) {
		return defaultVal
	}

	return val
}

// DoIfNotNil call the specified function, when target not nil and target.IsNil return true.
func DoIfNotNil(target any, fn func()) {
	if reflectx.IsNil(target) {
		return
	}

	type nilChecker interface{ IsNil() bool }
	if c, ok := target.(nilChecker); ok && c.IsNil() {
		return
	}

	fn()
}

// MustBeNil the error passed in must be empty, otherwise a panic occurs.
func MustBeNil(err error) {
	if err == nil {
		return
	}

	panic(err)
}
