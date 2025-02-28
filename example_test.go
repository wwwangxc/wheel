package wheel_test

import (
	"errors"
	"fmt"
	"time"

	"github.com/wwwangxc/wheel"
)

func ExampleValueOrDefault() {
	fmt.Println(wheel.ValueOrDefault("string", "default_string")) // string
	fmt.Println(wheel.ValueOrDefault("", "default_string"))       // default_string
	fmt.Println(wheel.ValueOrDefault(666, 888))                   // 666
	fmt.Println(wheel.ValueOrDefault(0, 888))                     // 666

	// Output:
	// string
	// default_string
	// 666
	// 888
}

func ExampleDoIfNotNil() {
	type S struct {
		Name string
	}

	var s *S
	wheel.DoIfNotNil(s, func() {
		fmt.Println(s.Name)
	})

	s = &S{Name: "Star Wang"}
	wheel.DoIfNotNil(s, func() {
		fmt.Println(s.Name)
	})

	// Output:
	// Star Wang
}

func ExampleMustBeNil() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var err error
	wheel.MustBeNil(err)

	err = errors.New("error message")
	wheel.MustBeNil(err)

	// Output:
	// error message
}

func ExampleTime() {
	t, err := time.Parse(time.DateTime, "2025-02-28 11:22:00")
	wheel.MustBeNil(err)

	fmt.Println(wheel.Time.BeginOfDay(t).Format(time.DateTime)) // "2025-02-28 00:00:00"
	fmt.Println(wheel.Time.EndOfDay(t).Format(time.DateTime))   // "2025-02-28 23:59:59"

	// Output:
	// 2025-02-28 00:00:00
	// 2025-02-28 23:59:59
}
