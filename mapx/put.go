package mapx

import (
	"github.com/wwwangxc/wheel"
	"github.com/wwwangxc/wheel/reflectx"
)

// PutIfNotZero put the value into the map, when val is not the zero value of the type.
func PutIfNotZero[K comparable, V any](m map[K]V, key K, val V) {
	if m == nil || reflectx.IsZeroValue(key) || reflectx.IsZeroValue(val) {
		return
	}

	m[key] = val
}

// PutOrDefault put the default value into the map, when the value is the zero value of the type.
func PutOrDefault[K comparable, V any](m map[K]V, key K, val, defaultVal V) {
	if m == nil || reflectx.IsZeroValue(key) {
		return
	}

	m[key] = wheel.Or(val, defaultVal)
}
