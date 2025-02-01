package mapx

import (
	"fmt"
	"reflect"

	"github.com/fatih/structs"

	"github.com/wwwangxc/wheel/reflectx"
)

// FlattenFromStruct flattens a struct into a single-level map
func FlattenFromStruct(target any) map[string]any {
	if target == nil {
		return map[string]any{}
	}

	s := structs.New(target)
	s.TagName = "json"
	return Flatten(s.Map())
}

// Flatten flattens a nested map into a single-level map
func Flatten(m map[string]any) map[string]any {
	if len(m) == 0 {
		return map[string]any{}
	}

	ret := map[string]any{}
	for k, v := range m {
		val := v

		if !reflectx.IsKindDeref(val, reflect.Map) {
			ret[k] = val
			continue
		}

		mm := Flatten(val.(map[string]any))
		for kk, vv := range mm {
			ret[fmt.Sprintf("%s.%s", k, kk)] = vv
		}
	}

	return ret
}
