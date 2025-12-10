package wheel_test

import (
	"errors"
	"fmt"
	"time"

	"github.com/wwwangxc/wheel"
)

func ExampleOr() {
	fmt.Println(wheel.Or("string_1", "string_2", "default_string")) // string_1
	fmt.Println(wheel.Or("", "string_2", "default_string"))         // string_2
	fmt.Println(wheel.Or("", "", "default_string"))                 // default_string
	fmt.Println(wheel.Or(666, 888))                                 // 666
	fmt.Println(wheel.Or(0, 888))                                   // 888

	// Output:
	// string_1
	// string_2
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
	a := float64(1.234567)

	b := float64(1.234567)
	fmt.Printf("%f + %f = ?\n", a, b)
	fmt.Println(wheel.Float.Add(a, b))             // 2.469134
	fmt.Println(wheel.Float.AddRounded(a, b, 2))   // 2.47
	fmt.Println(wheel.Float.AddTruncated(a, b, 2)) // 2.46

	b = float64(0.000001)
	fmt.Printf("%f - %f = ?\n", a, b)
	fmt.Println(wheel.Float.Sub(a, b))             // 1.234566
	fmt.Println(wheel.Float.SubRounded(a, b, 3))   // 1.235
	fmt.Println(wheel.Float.SubTruncated(a, b, 3)) // 1.234

	b = float64(1.1)
	fmt.Printf("%f * %.1f = ?\n", a, b)
	fmt.Println(wheel.Float.Mul(a, b))             // 1.3580237
	fmt.Println(wheel.Float.MulRounded(a, b, 2))   // 1.36
	fmt.Println(wheel.Float.MulTruncated(a, b, 2)) // 1.35

	b = float64(10)
	fmt.Printf("%f / %.0f = ?\n", a, b)
	fmt.Println(wheel.Float.Div(a, b))             // 0.1234567
	fmt.Println(wheel.Float.DivRounded(a, b, 4))   // 0.1235
	fmt.Println(wheel.Float.DivTruncated(a, b, 4)) // 0.1234

	b = float64(-1.234567)
	fmt.Printf("Abs(%f) = %f\n", b, wheel.Float.Abs(b)) // Abs(-1.234567) = 1.2345

	b = float64(5.9)
	fmt.Printf("Floor(%.1f) = %f\n", b, wheel.Float.Floor(b)) // Floor(5.9) = 5.000000

	b = float64(5.1)
	fmt.Printf("Ceil(%.1f) = %f\n", b, wheel.Float.Ceil(b)) // Ceil(5.1) = 6.000000

	b = float64(0.123456)
	fmt.Printf("Equal(%f, %f) = %t\n", a, a, wheel.Float.Equal(a, a)) // Equal(1.234567, 1.234567) = true
	fmt.Printf("Equal(%f, %f) = %t\n", a, b, wheel.Float.Equal(a, b)) // Equal(1.234567, 0.123456) = false
	fmt.Printf("GT(%f, %f) = %t\n", a, a, wheel.Float.GT(a, a))       // GT(1.234567, 1.234567) = false
	fmt.Printf("GT(%f, %f) = %t\n", a, b, wheel.Float.GT(a, b))       // GT(1.234567, 0.123456) = true
	fmt.Printf("GT(%f, %f) = %t\n", b, a, wheel.Float.GT(b, a))       // GT(0.123456, 1.234567) = false
	fmt.Printf("GTE(%f, %f) = %t\n", a, a, wheel.Float.GTE(a, a))     // GTE(1.234567, 1.234567) = true
	fmt.Printf("GTE(%f, %f) = %t\n", a, b, wheel.Float.GTE(a, b))     // GTE(1.234567, 0.123456) = true
	fmt.Printf("GTE(%f, %f) = %t\n", b, a, wheel.Float.GTE(b, a))     // GTE(0.123456, 1.234567) = false
	fmt.Printf("LT(%f, %f) = %t\n", a, a, wheel.Float.LT(a, a))       // LT(1.234567, 1.234567) = false
	fmt.Printf("LT(%f, %f) = %t\n", a, b, wheel.Float.LT(a, b))       // LT(1.234567, 0.123456) = false
	fmt.Printf("LT(%f, %f) = %t\n", b, a, wheel.Float.LT(b, a))       // LT(0.123456, 1.234567) = true
	fmt.Printf("LTE(%f, %f) = %t\n", a, a, wheel.Float.LTE(a, a))     // LTE(1.234567, 1.234567) = true
	fmt.Printf("LTE(%f, %f) = %t\n", a, b, wheel.Float.LTE(a, b))     // LTE(1.234567, 0.123456) = false
	fmt.Printf("LTE(%f, %f) = %t\n", b, a, wheel.Float.LTE(b, a))     // LTE(0.123456, 1.234567) = true

	// Output:
	// 1.234567 + 1.234567 = ?
	// 2.469134
	// 2.47
	// 2.46
	// 1.234567 - 0.000001 = ?
	// 1.234566
	// 1.235
	// 1.234
	// 1.234567 * 1.1 = ?
	// 1.3580237
	// 1.36
	// 1.35
	// 1.234567 / 10 = ?
	// 0.1234567
	// 0.1235
	// 0.1234
	// Abs(-1.234567) = 1.234567
	// Floor(5.9) = 5.000000
	// Ceil(5.1) = 6.000000
	// Equal(1.234567, 1.234567) = true
	// Equal(1.234567, 0.123456) = false
	// GT(1.234567, 1.234567) = false
	// GT(1.234567, 0.123456) = true
	// GT(0.123456, 1.234567) = false
	// GTE(1.234567, 1.234567) = true
	// GTE(1.234567, 0.123456) = true
	// GTE(0.123456, 1.234567) = false
	// LT(1.234567, 1.234567) = false
	// LT(1.234567, 0.123456) = false
	// LT(0.123456, 1.234567) = true
	// LTE(1.234567, 1.234567) = true
	// LTE(1.234567, 0.123456) = false
	// LTE(0.123456, 1.234567) = true
}
