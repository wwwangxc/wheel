package reflectx_test

import (
	"fmt"
	"reflect"

	"github.com/wwwangxc/wheel/reflectx"
)

func ExampleIsKind() {
	kindString := ""
	kindInt := 0
	kindFloat64 := float64(0)
	kindMap := map[string]any{}

	fmt.Println(reflectx.IsKind(kindString, reflect.String))   // true
	fmt.Println(reflectx.IsKind(kindInt, reflect.Int))         // true
	fmt.Println(reflectx.IsKind(kindFloat64, reflect.Float64)) // true
	fmt.Println(reflectx.IsKind(kindMap, reflect.Map))         // true

	// Output:
	// true
	// true
	// true
	// true
}

func ExampleIsKindDeref() {
	kindString := ""
	kindInt := 0
	kindFloat64 := float64(0)
	kindMap := map[string]any{}

	fmt.Println(reflectx.IsKind(&kindString, reflect.String))        // false
	fmt.Println(reflectx.IsKindDeref(kindString, reflect.String))    // true
	fmt.Println(reflectx.IsKindDeref(&kindString, reflect.String))   // true
	fmt.Println(reflectx.IsKind(&kindInt, reflect.Int))              // false
	fmt.Println(reflectx.IsKindDeref(kindInt, reflect.Int))          // true
	fmt.Println(reflectx.IsKindDeref(&kindInt, reflect.Int))         // true
	fmt.Println(reflectx.IsKind(&kindFloat64, reflect.Float64))      // false
	fmt.Println(reflectx.IsKindDeref(kindFloat64, reflect.Float64))  // true
	fmt.Println(reflectx.IsKindDeref(&kindFloat64, reflect.Float64)) // true
	fmt.Println(reflectx.IsKind(&kindMap, reflect.Map))              // false
	fmt.Println(reflectx.IsKindDeref(kindMap, reflect.Map))          // true
	fmt.Println(reflectx.IsKindDeref(&kindMap, reflect.Map))         // true

	// Output:
	// false
	// true
	// true
	// false
	// true
	// true
	// false
	// true
	// true
	// false
	// true
	// true
}

func ExampleIsZeroValue() {
	fmt.Println(reflectx.IsZeroValue(""))                             // true
	fmt.Println(reflectx.IsZeroValue("string"))                       // false
	fmt.Println(reflectx.IsZeroValue(0))                              // true
	fmt.Println(reflectx.IsZeroValue(1))                              // false
	fmt.Println(reflectx.IsZeroValue(map[string]any{}))               // true
	fmt.Println(reflectx.IsZeroValue(map[string]any{"key": "value"})) // false
	fmt.Println(reflectx.IsZeroValue(nil))                            // true

	// Outout:
	// true
	// false
	// true
	// false
	// true
	// false
	// true
}

func ExampleIsNil() {
	type Interface interface{ HelloWord() }
	var interfaceInstant Interface
	var ch chan struct{}
	var fn func()
	var m map[string]any
	var mPtr *map[string]any
	var s []string

	fmt.Println(reflectx.IsNil(nil))              // true
	fmt.Println(reflectx.IsNil(interfaceInstant)) // true
	fmt.Println(reflectx.IsNil(ch))               // true
	fmt.Println(reflectx.IsNil(fn))               // true
	fmt.Println(reflectx.IsNil(m))                // true
	fmt.Println(reflectx.IsNil(mPtr))             // true
	fmt.Println(reflectx.IsNil(s))                // true

	fmt.Println(reflectx.IsNil("")) // false
	fmt.Println(reflectx.IsNil(0))  // false

	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	// false
	// false
}

func ExampleDeref() {
	type S struct{}

	tp := reflect.TypeOf(S{})
	tpPtr := reflect.TypeOf(&S{})

	fmt.Println(reflectx.Deref(tp) == reflect.TypeOf(S{}))    // true
	fmt.Println(reflectx.Deref(tpPtr) == reflect.TypeOf(S{})) // true

	// Output:
	// true
	// true
}
