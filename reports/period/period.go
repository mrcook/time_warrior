package period

import (
	"time"
)

type Period struct {
	period    string
	startTime time.Time
	endTime   time.Time
}

func Parse(unit string) *Period {
	p := &Period{period: unit}

	now := time.Now()
	var start, end time.Time
	var period string

	switch unit {
	case "t":
		period = "Today"
		start = p.BeginningOfDay(now)
		end = p.EndOfDay(now)
	case "w":
		period = "This Week"
		start = p.BeginningOfWeek(now)
		end = p.EndOfWeek(now)
	case "m":
		period = "This Month"
		start = p.BeginningOfMonth(now)
		end = p.EndOfMonth(now)
	case "y":
		period = "This Year"
		start = p.BeginningOfYear(now)
		end = p.EndOfYear(now)
	case "1d":
		period = "Yesterday"
		t := p.Yesterday(now)
		start = p.BeginningOfDay(t)
		end = p.EndOfDay(t)
	case "1w":
		period = "Last Week"
		start = p.BeginningOfPreviousWeek(now)
		end = p.EndOfPreviousWeek(now)
	case "1m":
		period = "Last Month"
		start = p.BeginningOfPreviousMonth(now)
		end = p.EndOfPreviousMonth(now)
	case "1y":
		period = "Last Year"
		start = p.BeginningOfPreviousYear(now)
		end = p.EndOfPreviousYear(now)
	default:
		start = p.BeginningOfDay(now)
		end = p.EndOfDay(now)
	}

	p.period = period
	p.startTime = start
	p.endTime = end

	return p
}

func (p Period) Period() string {
	return p.period
}

func (p Period) From() time.Time {
	return p.startTime
}

func (p Period) To() time.Time {
	return p.endTime
}

func (p Period) IsSet() bool {
	return p.period != ""
}

func (p Period) BeginningOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func (p Period) EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, t.Location())
}

func (p Period) Yesterday(t time.Time) time.Time {
	return t.AddDate(0, 0, -1)
}

func (p Period) lastWeek(t time.Time) time.Time {
	return t.AddDate(0, 0, -7)
}

func (p Period) BeginningOfWeek(t time.Time) time.Time {
	for t.Weekday() != time.Monday {
		t = t.AddDate(0, 0, -1)
	}
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func (p Period) EndOfWeek(t time.Time) time.Time {
	return p.BeginningOfWeek(t).AddDate(0, 0, 7).Add(-time.Second)
}

func (p Period) BeginningOfPreviousWeek(t time.Time) time.Time {
	lastWeek := p.lastWeek(t)
	return p.BeginningOfWeek(lastWeek)
}

func (p Period) EndOfPreviousWeek(t time.Time) time.Time {
	return p.BeginningOfPreviousWeek(t).AddDate(0, 0, 7).Add(-time.Second)
}

func (p Period) BeginningOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

func (p Period) EndOfMonth(t time.Time) time.Time {
	return p.BeginningOfMonth(t).AddDate(0, 1, 0).Add(-time.Second)
}

func (p Period) BeginningOfPreviousMonth(t time.Time) time.Time {
	year, month, _ := p.EndOfPreviousMonth(t).Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
}

func (p Period) EndOfPreviousMonth(t time.Time) time.Time {
	return p.BeginningOfMonth(t).Add(-time.Second)
}

func (p Period) BeginningOfYear(t time.Time) time.Time {
	y, _, _ := t.Date()
	return time.Date(y, time.January, 1, 0, 0, 0, 0, t.Location())
}

func (p Period) EndOfYear(t time.Time) time.Time {
	y, _, _ := t.Date()
	return time.Date(y, time.December, 31, 23, 59, 59, 0, t.Location())
}

func (p Period) BeginningOfPreviousYear(t time.Time) time.Time {
	return p.BeginningOfYear(t).AddDate(-1, 0, 0)
}

func (p Period) EndOfPreviousYear(t time.Time) time.Time {
	return p.BeginningOfPreviousYear(t).AddDate(1, 0, 0).Add(-time.Second)
}
