package wheel

import "time"

const (
	maxHour int = 23
	maxMin  int = 59
	maxSec  int = 59
	MaxNsec int = int(time.Second - time.Nanosecond)
)

// Time helper
var Time = (*helperTime)(nil)

type helperTime struct{}

// BeginOfDayNow returns the start time of the day now
func (s *helperTime) BeginOfDayNow() time.Time {
	return s.BeginOfDay(time.Now())
}

// BeginOfDay returns the start time of the day for the target in local zone
func (s *helperTime) BeginOfDay(t time.Time) time.Time {
	y, m, d := t.In(time.Local).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

// EndOfDayNow returns the end time of the day for the target in local zone
func (s *helperTime) EndOfDayNow() time.Time {
	return s.EndOfDay(time.Now())
}

// EndOfDay returns the end time of the day for the target in local zone
func (s *helperTime) EndOfDay(t time.Time) time.Time {
	y, m, d := t.In(time.Local).Date()
	return time.Date(y, m, d, maxHour, maxMin, maxSec, MaxNsec, time.Local)
}
