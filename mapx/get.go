package mapx

import (
	"github.com/wwwangxc/wheel/reflectx"
)

// GetOrDefault get the value of key in the map.
// When the key does not exist, it return the specified default value.
func GetOrDefault[K comparable, V any](m map[K]V, key K, defaultVal V) V {
	if len(m) == 0 || reflectx.IsZeroValue(key) {
		return defaultVal
	}

	if v, ok := m[key]; ok {
		return v
	}

	return defaultVal
}
