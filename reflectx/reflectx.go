// Package reflectx is the helper of std package reflect
package reflectx

import "reflect"

// IsKind returns true if targe specific kind is the target kind
func IsKind(v any, targetKind reflect.Kind) bool {
	if v == nil {
		return false
	}

	return reflect.TypeOf(v).Kind() == targetKind
}

// IsKindDeref returns true if targe specific kind is the target kind
func IsKindDeref(v any, targetKind reflect.Kind) bool {
	if v == nil {
		return false
	}

	return Deref(reflect.TypeOf(v)).Kind() == targetKind
}

// IsZeroValue return true when target is the zero value of type
func IsZeroValue(v any) bool {
	if v == nil {
		return true
	}

	val := reflect.ValueOf(v)
	return val.IsZero()
}

// IsNil return true when target is nil
func IsNil(v any) bool {
	if v == nil {
		return true
	}

	switch val := reflect.ValueOf(v); val.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.Pointer,
		reflect.UnsafePointer,
		reflect.Interface,
		reflect.Slice:
		return val.IsNil()

	default:
		return false
	}
}

// Deref of type
func Deref(tp reflect.Type) reflect.Type {
	if tp == nil {
		return tp
	}

	for tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}

	return tp
}
