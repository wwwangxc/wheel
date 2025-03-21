package wheel

import (
	"github.com/shopspring/decimal"
)

// Float helper
var Float = (*helperFloat)(nil)

type helperFloat struct{}

// Add returns the result of a + b
func (s *helperFloat) Add(a, b float64) float64 {
	return s.add(a, b).InexactFloat64()
}

// AddRounded returns the result of a + b rounded to the given places
func (s *helperFloat) AddRounded(a, b float64, places uint8) float64 {
	return s.add(a, b).Round(int32(places)).InexactFloat64()
}

// AddTruncated returns the result of a + b truncated to the given places
func (s *helperFloat) AddTruncated(a, b float64, places uint8) float64 {
	return s.add(a, b).Truncate(int32(places)).InexactFloat64()
}

// Sub returns the result of a - b
func (s *helperFloat) Sub(a, b float64) float64 {
	return s.sub(a, b).InexactFloat64()
}

// SubRounded returns the result of a - b rounded to the given places
func (s *helperFloat) SubRounded(a, b float64, places uint8) float64 {
	return s.sub(a, b).Round(int32(places)).InexactFloat64()
}

// SubTruncated returns the result of a - b truncated to the given places
func (s *helperFloat) SubTruncated(a, b float64, places uint8) float64 {
	return s.sub(a, b).Truncate(int32(places)).InexactFloat64()
}

// Mul returns the result of a * b
func (s *helperFloat) Mul(a, b float64) float64 {
	return s.mul(a, b).InexactFloat64()
}

// MulRounded returns the result of a * b rounded to the given places
func (s *helperFloat) MulRounded(a, b float64, places uint8) float64 {
	return s.mul(a, b).Round(int32(places)).InexactFloat64()
}

// MulTruncated returns the result of a * b truncated to the given places
func (s *helperFloat) MulTruncated(a, b float64, places uint8) float64 {
	return s.mul(a, b).Truncate(int32(places)).InexactFloat64()
}

// Div returns the result of a / b
func (s *helperFloat) Div(a, b float64) float64 {
	return s.div(a, b).InexactFloat64()
}

// DivRounded returns the result of a / b rounded to the given places
func (s *helperFloat) DivRounded(a, b float64, places uint8) float64 {
	return s.div(a, b).Round(int32(places)).InexactFloat64()
}

// DivTruncated returns the result of a / b truncated to the given places
func (s *helperFloat) DivTruncated(a, b float64, places uint8) float64 {
	return s.div(a, b).Truncate(int32(places)).InexactFloat64()
}

// RoundTo returns the result of rounding value to the given places
//
//	(1.2345, 2) => 1.23
//	(1.2345, 3) => 1.235
func (s *helperFloat) RoundTo(value float64, places uint8) float64 {
	return decimal.NewFromFloat(value).Round(int32(places)).InexactFloat64()
}

// TruncateTo returns the result of truncating value to the given places
//
//	(1.2345, 2) => 1.23
//	(1.2345, 3) => 1.234
func (s *helperFloat) TruncateTo(value float64, places uint8) float64 {
	return decimal.NewFromFloat(value).Truncate(int32(places)).InexactFloat64()
}

func (*helperFloat) add(a, b float64) decimal.Decimal {
	d := decimal.NewFromFloat(a)
	d2 := decimal.NewFromFloat(b)
	return d.Add(d2)
}

func (*helperFloat) sub(a, b float64) decimal.Decimal {
	d := decimal.NewFromFloat(a)
	d2 := decimal.NewFromFloat(b)
	return d.Sub(d2)
}

func (*helperFloat) mul(a, b float64) decimal.Decimal {
	d := decimal.NewFromFloat(a)
	d2 := decimal.NewFromFloat(b)
	return d.Mul(d2)
}

func (*helperFloat) div(a, b float64) decimal.Decimal {
	d := decimal.NewFromFloat(a)
	d2 := decimal.NewFromFloat(b)
	return d.Div(d2)
}
