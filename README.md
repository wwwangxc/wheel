# This is my '🛞'

Hahaha

## Contents

- [Install](#install)
- [Quick Start](#quick-start)
  - [wheel](#wheel)
    - [ValueOrDefault](#valueordefault)
    - [DoIfNotNil](#doifnotnil)
    - [MustBeNil](#mustbenil)
    - [Time](#time)
      - [(*Time) BeginOfDay](#\(*time\)-beginofday)
      - [(*Time) EndOfDay](#\(*time\)-endofday)
  - [wheel/coroutine](#wheelcoroutine)
    - [Go](#go)
    - [Group](#group)
  - [wheel/errorx](#wheelerrorx)
  - [wheel/mapx](#wheelmapx)
    - [Flatten](#flatten)
    - [FlattenFromStruct](#flattenfromstruct)
    - [GetOrDefault](#getordefault)
    - [PutIfNotZero](#putifnotzero)
    - [PutOrDefault](#putordefault)
  - [wheel/reflectx](#wheelreflectx)
    - [IsKind](#iskind)
    - [IsKindDeref](#iskindderef)
    - [IsZeroValue](#iszerovalue)
    - [IsNil](#isnil)
    - [Deref](#deref)
  - [wheel/syncx](#wheelsyncx)
    - [WaitGroup](#waitgroup)

## Install

```sh
go get github.com/wwwangxc/wheel
```

**[⬆ back to top](#contents)**

## Quick Start

### wheel

About the global functions.

#### ValueOrDefault

When the value is the zero value of the type, return the specified default value.

```go
package main

import (
    "github.com/wwwangxc/wheel"
)

func main() {
    _ = wheel.ValueOrDefault("string", "default") // string
    _ = wheel.ValueOrDefault("", "default")       // default
}
```

**[⬆ back to top](#contents)**

#### DoIfNotNil

Call the specified function, when target not nil or target.IsNil method return true.

```go
package main

import (
    "github.com/wwwangxc/wheel"
)

type S struct {
    Name string
}

func main() {
    var s *S
    wheel.DoIfNotNil(s, func(){
        // It will not be executed
    })

    s = &s{}
    wheel.DoIfNotNil(s, func(){
        // It will be executed
    })
}
```

**[⬆ back to top](#contents)**

#### MustBeNil

It will be panic when error not nil.

```go
package main

import (
    "github.com/wwwangxc/wheel"
)

func main() {
    defer func() {
        if err := recover(); err != nil {
            // it will be executed
        }
    }()

    err := errors.New("error message")

    // it will be panic
    wheel.MustBeNil(err)
}
```

**[⬆ back to top](#contents)**

#### Time

##### (*Time) BeginOfDay

```go
package main

import (
    "github.com/wwwangxc/wheel"
)

func main() {
    t, _ := time.Parse(time.DateTime, "2025-02-28 11:22:00")
    beginOfDay := wheel.Time.BeginOfDay(t) // 2025-02-28 00:00:00 +0800 CST
    beginOfDay = wheel.Time.BeginOfDayNow()
```

**[⬆ back to top](#contents)**

##### (*Time) EndOfDay

```go
package main

import (
    "github.com/wwwangxc/wheel"
)

func main() {
    t, _ := time.Parse(time.DateTime, "2025-02-28 11:22:00")
    endOfDay := wheel.Time.EndOfDay(t) // 2025-02-28 23:59:59.999999999 +0800 CST
    endOfDay = wheel.Time.EndOfDayNow()
}
```

**[⬆ back to top](#contents)**

### wheel/coroutine

About the goroutine.

#### Go

Safed go function.

- You can set sync.WaitGroup and it will call Add(1) before call the specified function,
and it will call Done() after call the specified function too.

- You can set logging write function, it will be called when panic.

```go
package main

import (
    "github.com/wwwangxc/wheel/coroutine"
)

func main() {
    var wg sync.WaitGroup
    logFn := func(v ...any) {
        fmt.Println(v...)
    }

    coroutine.Go(
        func(){
            // do something...
        },
        coroutine.WithWaitGroup(&wg),      // set sync.WaitGroup
        coroutine.WithLogWhenPainc(logFn)) // write log when panic

    // Wait group
    wg.Wait()
}
```

**[⬆ back to top](#contents)**

#### Group

Group of coroutine task.

```go
package main

import (
    "context"

    "github.com/wwwangxc/wheel/coroutine"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    g := coroutine.NewGroup(ctx,
        coroutine.WithConcurencyLevel(1),   // set the concurrency level, default: ${cpu number} * 10
        coroutine.WithCancelOnError(),      // exit immediately when any coroutine returns an error, default: false
        coroutine.WithTimeout(time.Second)) // set the global timeout, default 3s

    g.Go(func(ctx context.Context) error {
        // do something...
        return errors.New("error message")
    })

    ret := g.Wait()

    // a collection of all errors returned by the given function
    ret.Errors()

    // merged error of all errors returned by the given function
    ret.Error()
    // error message such as:
    // coroutine group already timeout
    // 1 errors occurred:
    //     * error message
}
```

**[⬆ back to top](#contents)**

### wheel/errorx

About custom error.

```go
package main

import (
    "github.com/wwwangxc/wheel/errorx"
)

// CodeServerError server error code
const CodeServerError errorx.ErrCode = 500_00

// ErrServer internal server error
var ErrServer = errorx.New(CodeServerError, "Internal Server Error", errorx.WithReason("This is the server error"))

func main() {
    fmt.Printf("Code: %d\n", ErrServer.Code())       // Code: 50000
    fmt.Printf("Message: %s\n", ErrServer.Message()) // Message: Internal Server Error
    fmt.Printf("Reason: %s\n", ErrServer.Reason())   // Reason: This is the server error
    fmt.Printf("Error: %s\n", ErrServer.Error())     // Error: 50000 Internal Server Error
    fmt.Printf("String: %s\n", ErrServer)            // String: 50000 Internal Server Error
}
```

**[⬆ back to top](#contents)**

### wheel/mapx

About the map.

#### Flatten

Flatten the map.

```go
package main

import (
    "github.com/wwwangxc/wheel/errorx"
)

func main() {
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
    // m is:
    //  field_1 => value_1
    //  field_2.field_1 => value_1
    //  field_3.field_1 => value_1
}
```

**[⬆ back to top](#contents)**

#### FlattenFromStruct

Convert struct to flattened map

```go
package main

import (
    "github.com/wwwangxc/wheel/errorx"
)

type S struct {
    Field1 string `json:"field_1"`
    Field2 struct {
        Field1 string `json:"field_1"`
    } `json:"field_2"`
    Field3 map[string]any `json:"field_3"`
}

func main() {
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
    // m is:
    //  field_1 => value_1
    //  field_2.field_1 => value_1
    //  field_3.field_1 => value_1
}
```

**[⬆ back to top](#contents)**

#### GetOrDefault

Get the value from the map.
When the key is not exists, it will be return the specified default value.

```go
package main

import (
    "github.com/wwwangxc/wheel/mapx"
)

func main() {
    m := map[string]string{
        "key": "value",
    }

    key := "key"
    _ = mapx.GetOrDefault(m, key, "default") // value

    key := "other key"
    _ = mapx.GetOrDefault(m, key, "default") // default
}
```

**[⬆ back to top](#contents)**

#### PutIfNotZero

When the value is not the zero value of the type, put the K/V into the map.

```go
package main

import (
    "github.com/wwwangxc/wheel/mapx"
)

func main() {
    m := map[string]any{}
    mapx.PutIfNotZero(m, "key_1", "value")
    mapx.PutIfNotZero(m, "key_2", "")
    mapx.PutIfNotZero(m, "key_3", 0)
    mapx.PutIfNotZero(m, "key_4", nil)

    v, ok := m["key_1"] // value, true
    v, ok = m["key_2"]  // nil, false
    v, ok = m["key_3"]  // nil, false
    v, ok = m["key_4"]  // nil, false
}
```

**[⬆ back to top](#contents)**

#### PutOrDefault

Put the K/V into the map.
When the value is the zero value of the type, it will be replaced with the specified default value.

```go
package main

import (
    "github.com/wwwangxc/wheel/mapx"
)

func main() {
    m := map[string]any{}
    mapx.PutOrDefault(m, "key_1", "value", "default")
    mapx.PutOrDefault(m, "key_2", "", "default")
    mapx.PutOrDefault(m, "key_3", 666, 888)
    mapx.PutOrDefault(m, "key_4", 0, 888)

    v, ok := m["key_1"] // value, true
    v, ok = m["key_2"]  // default, true
    v, ok = m["key_3"]  // 666, true
    v, ok = m["key_4"]  // 888, true
}
```

**[⬆ back to top](#contents)**

### wheel/reflectx

About std package `reflect` helper.

#### IsKind

Check if the specified kind is the target kind.

```go
package main

import (
    "github.com/wwwangxc/wheel/reflectx"
)

func main() {
    kindString := ""
    reflectx.IsKind(kindString, reflect.String) // true

    kindInt := 0
    reflectx.IsKind(kindInt, reflect.Int) // true

    kindFloat64 := float64(0)
    reflectx.IsKind(kindFloat64, reflect.Float64) // true

    kindMap := map[string]any{}
    reflectx.IsKind(kindMap, reflect.Map) // true
}
```

**[⬆ back to top](#contents)**

#### IsKindDeref

Check if the specified kind is the target kind.
If specified kind is pointer, it will be dereference.

```go
package main

import (
    "github.com/wwwangxc/wheel/reflectx"
)

func main() {
    kindString := ""
    reflectx.IsKind(&kindString, reflect.String)      // false
    reflectx.IsKindDeref(kindString, reflect.String)  // true
    reflectx.IsKindDeref(&kindString, reflect.String) // true

    kindInt := 0
    reflectx.IsKind(&kindInt, reflect.Int)      // false
    reflectx.IsKindDeref(kindInt, reflect.Int)  // true
    reflectx.IsKindDeref(&kindInt, reflect.Int) // true

    kindFloat64 := float64(0)
    reflectx.IsKind(&kindFloat64, reflect.Float64)      // false
    reflectx.IsKindDeref(kindFloat64, reflect.Float64)  // true
    reflectx.IsKindDeref(&kindFloat64, reflect.Float64) // true

    kindMap := map[string]any{}
    reflectx.IsKind(&kindMap, reflect.Map)      // false
    reflectx.IsKindDeref(kindMap, reflect.Map)  // true
    reflectx.IsKindDeref(&kindMap, reflect.Map) // true
}
```

**[⬆ back to top](#contents)**

#### IsZeroValue

Check if the target value is the zero value of the type.

```go
package main

import (
    "github.com/wwwangxc/wheel/reflectx"
)

func main() {
    // string
    reflectx.IsZeroValue("")       // true
    reflectx.IsZeroValue("string") // false

    // int
    reflectx.IsZeroValue(0) // true
    reflectx.IsZeroValue(1) // false

    // map
    reflectx.IsZeroValue(map[string]any{})               // true
    reflectx.IsZeroValue(map[string]any{"key": "value"}) // false

    // nil
    reflectx.IsZeroValue(nil) // true
}
```

**[⬆ back to top](#contents)**

#### IsNil

Check if the target value is not nil.

```go
package main

import (
    "github.com/wwwangxc/wheel/reflectx"
)

func main() {
    // nil
    reflectx.IsNil(nil) // true

    // interface
    type Interface interface{ HelloWord() }
    var interfaceInstant Interface
    reflectx.IsNil(interfaceInstant) // true

    // channel
    var ch chan struct{}
    reflectx.IsNil(ch) // true

    // function
    var fn func()
    reflectx.IsNil(fn) // true

    // map
    var m map[string]any
    reflectx.IsNil(m) // true

    // *map
    var mPtr *map[string]any
    reflectx.IsNil(m) // true

    // slice
    var s []string
    reflectx.IsNil(s) // true
}
```

**[⬆ back to top](#contents)**

#### Deref

Dereference the type.

```go
package main

import (
    "reflect"

    "github.com/wwwangxc/wheel/reflectx"
)

type S struct{}

func main() {
    tp := reflect.TypeOf(S{})
    tpPtr := reflect.TypeOf(&S{})

    fmt.Println(reflectx.Deref(tp) == reflect.TypeOf(S{}))    // true
    fmt.Println(reflectx.Deref(tpPtr) == reflect.TypeOf(S{})) // true
}
```

**[⬆ back to top](#contents)**

### wheel/syncx

About std package `sync` helper.

#### WaitGroup

The helper of std `sync.WaitGroup`.

```go
package main

import (
    "context"

    "github.com/wwwangxc/wheel/syncx"
)

func main() {
    var wg syncx.WaitGroup

    wg.Add(1)
    go func() {
        defer wg.Done()
        // do something...
    }()

    wg.Add(1)
    go func() {
        defer wg.Done()
        // do something...
    }()

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    select {

    // Conatext Done
    case <-ctx.Done(): 
        // do something...

    // All Task Done
    case <-wg.Wait():
        // do something...
    }

    // or

    if err := wg.WaitOrDone(ctx); err != nil {
        switch {

        // Context Canceled
        case errors.Is(err, context.Canceled):
            // do something..

        // Timeout
        case errors.Is(err, context.DeadlineExceeded):
            // do something..

        default:
        }
    }
}
```

**[⬆ back to top](#contents)**
