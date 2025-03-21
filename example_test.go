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

func ExampleFloat() {
	a := float64(1.23456789)
	b := float64(1.23456789)

	fmt.Println("1.23456789 + 1.23456789 = ?")
	fmt.Println(wheel.Float.Add(a, b))             // 2.46913578
	fmt.Println(wheel.Float.AddRounded(a, b, 2))   // 2.47
	fmt.Println(wheel.Float.AddTruncated(a, b, 2)) // 2.46

	fmt.Println("1.23456789 - 0.00000001 = ?")
	fmt.Println(wheel.Float.Sub(a, float64(0.00000001)))             // 1.23456788
	fmt.Println(wheel.Float.SubRounded(a, float64(0.00000001), 4))   // 1.2346
	fmt.Println(wheel.Float.SubTruncated(a, float64(0.00000001), 4)) // 1.2345

	fmt.Println("1.23456789 * 1.1 = ?")
	fmt.Println(wheel.Float.Mul(a, float64(1.1)))             // 1.358024679
	fmt.Println(wheel.Float.MulRounded(a, float64(1.1), 2))   // 1.36
	fmt.Println(wheel.Float.MulTruncated(a, float64(1.1), 2)) // 1.35

	fmt.Println("1.23456789 / 10 = ?")
	fmt.Println(wheel.Float.Div(a, float64(10)))             // 0.123456789
	fmt.Println(wheel.Float.DivRounded(a, float64(10), 4))   // 0.1235
	fmt.Println(wheel.Float.DivTruncated(a, float64(10), 4)) // 0.1234

	// Output:
	// 1.23456789 + 1.23456789 = ?
	// 2.46913578
	// 2.47
	// 2.46
	// 1.23456789 - 0.00000001 = ?
	// 1.23456788
	// 1.2346
	// 1.2345
	// 1.23456789 * 1.1 = ?
	// 1.358024679
	// 1.36
	// 1.35
	// 1.23456789 / 10 = ?
	// 0.123456789
	// 0.1235
	// 0.1234
}
