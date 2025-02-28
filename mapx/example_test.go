package mapx_test

import (
	"fmt"
	"sort"

	"github.com/wwwangxc/wheel/mapx"
)

func ExampleFlattenFromStruct() {
	type S struct {
		Field1 string `json:"field_1"`
		Field2 struct {
			Field1 string `json:"field_1"`
		} `json:"field_2"`
		Field3 map[string]any `json:"field_3"`
	}
	s := &S{
		Field1: "value_1",
		Field2: struct {
			Field1 string `json:"field_1"`
		}{
			Field1: "value_1",
		},
		Field3: map[string]any{
			"field_1": "value_1",
		},
	}

	m := mapx.FlattenFromStruct(s)

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, v := range keys {
		fmt.Printf("%s => %s\n", v, m[v])
	}

	// Output:
	// field_1 => value_1
	// field_2.field_1 => value_1
	// field_3.field_1 => value_1
}

func ExampleFlatten() {
	m := map[string]any{
		"field_1": "value_1",
		"field_2": struct {
			Field1 string `json:"field_1"`
		}{
			Field1: "value_1",
		},
		"field_3": map[string]any{
			"field_1": "value_1",
		},
	}

	flattend := mapx.Flatten(m)

	keys := make([]string, 0, len(flattend))
	for k := range flattend {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, v := range keys {
		fmt.Printf("%s => %s\n", v, flattend[v])
	}

	// Output:
	// field_1 => value_1
	// field_2.field_1 => value_1
	// field_3.field_1 => value_1
}

func ExampleGetOrDefault() {
	m := map[string]string{
		"key": "value",
	}

	key := "key"
	fmt.Printf("%s => %s\n", key, mapx.GetOrDefault(m, key, "default_value")) // key => value

	key = "other_key"
	fmt.Printf("%s => %s\n", key, mapx.GetOrDefault(m, key, "default_value")) // other_key => default_value

	// Output:
	// key => value
	// other_key => default_value
}

func ExamplePutIfNotZero() {
	m := map[string]any{}
	mapx.PutIfNotZero(m, "key_1", "value_1")
	mapx.PutIfNotZero(m, "key_2", "")
	mapx.PutIfNotZero(m, "key_3", 0)
	mapx.PutIfNotZero(m, "key_4", nil)

	_, ok := m["key_1"]
	fmt.Println(ok) // true
	_, ok = m["key_2"]
	fmt.Println(ok) // false
	_, ok = m["key_3"]
	fmt.Println(ok) // false
	_, ok = m["key_4"]
	fmt.Println(ok) // false

	// Output:
	// true
	// false
	// false
	// false
}

func ExamplePutOrDefault() {
	m := map[string]any{}
	mapx.PutOrDefault(m, "key_1", "value_1", "default_value")
	mapx.PutOrDefault(m, "key_2", "", "default_value")
	mapx.PutOrDefault(m, "key_3", 0, 666)

	v, ok := m["key_1"]
	fmt.Println(ok) // true
	fmt.Println(v)  // value_1
	v, ok = m["key_2"]
	fmt.Println(ok) // true
	fmt.Println(v)  // default_value
	v, ok = m["key_3"]
	fmt.Println(ok) // true
	fmt.Println(v)  // 666

	// Output:
	// true
	// value_1
	// true
	// default_value
	// true
	// 666
}
