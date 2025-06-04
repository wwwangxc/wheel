package wheel

import "github.com/wwwangxc/wheel/reflectx"

// Or return the first non-zero value in the list.
// If all values are zero, return the first value.
func Or[T any](val T, vals ...T) T {
	if !reflectx.IsZeroValue(val) {
		return val
	}

	for _, v := range vals {
		if reflectx.IsZeroValue(v) {
			continue
		}
		return v
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
