package period_test

import (
	"testing"
	"time"

	"github.com/mrcook/time_warrior/reports/period"
)

func TestPeriodToday(t *testing.T) {
	p := period.Parse("t")

	now := time.Now()
	bod := p.BeginningOfDay(now)
	eod := p.EndOfDay(now)

	if p.From().Unix() != bod.Unix() {
		t.Errorf("Expected to be start of today: %d, got %d", bod.Unix(), p.From().Unix())
	}
	if p.To().Unix() != eod.Unix() {
		t.Errorf("Expected to be end of today: %d, got %d", eod.Unix(), p.To().Unix())
	}
	if p.Period() != "Today" {
		t.Errorf("Expected correct period string, got %s", p.Period())
	}
}

func TestPeriodDefault(t *testing.T) {
	p := period.Parse("bad time unit")

	now := time.Now()
	bod := p.BeginningOfDay(now)
	eod := p.EndOfDay(now)

	if p.From().Unix() != bod.Unix() {
		t.Errorf("Expected default to be start of today: %d, got %d", bod.Unix(), p.From().Unix())
	}
	if p.To().Unix() != eod.Unix() {
		t.Errorf("Expected default to be end of today: %d, got %d", eod.Unix(), p.To().Unix())
	}
	if p.Period() != "" {
		t.Errorf("Expected empty period string, got %s", p.Period())
	}
}

func TestPeriodYesterday(t *testing.T) {
	p := period.Parse("1d")

	yesterday := p.Yesterday(time.Now())
	boy := p.BeginningOfDay(yesterday)
	eoy := p.EndOfDay(yesterday)

	if p.From().Unix() != boy.Unix() {
		t.Errorf("Expected beginning of yesterday: %d, got %d", boy.Unix(), p.From().Unix())
	}
	if p.To().Unix() != eoy.Unix() {
		t.Errorf("Expected end of yesterday: %d, got %d", eoy.Unix(), p.To().Unix())
	}
	if p.Period() != "Yesterday" {
		t.Errorf("Expected correct period string, got %s", p.Period())
	}
}

func TestPeriodWeekToDate(t *testing.T) {
	p := period.Parse("w")

	now := time.Now()
	bow := p.BeginningOfWeek(now)
	eow := p.EndOfWeek(now)

	if p.From().Unix() != bow.Unix() {
		t.Errorf("Expected start of this week: %d, got %d", bow.Unix(), p.To().Unix())
	}
	if p.To().Unix() != eow.Unix() {
		t.Errorf("Expected end of this week: %d, got %d", eow.Unix(), p.To().Unix())
	}
	if p.Period() != "This Week" {
		t.Errorf("Expected correct period string, got %s", p.Period())
	}
}

func TestPeriodMonthToDate(t *testing.T) {
	p := period.Parse("m")

	now := time.Now()
	bow := p.BeginningOfMonth(now)
	eow := p.EndOfMonth(now)

	if p.From().Unix() != bow.Unix() {
		t.Errorf("Expected start of this month: %d, got %d", bow.Unix(), p.To().Unix())
	}
	if p.To().Unix() != eow.Unix() {
		t.Errorf("Expected end of this month: %d, got %d", eow.Unix(), p.To().Unix())
	}
	if p.Period() != "This Month" {
		t.Errorf("Expected correct period string, got %s", p.Period())
	}
}

func TestPeriodYearToDate(t *testing.T) {
	p := period.Parse("y")

	now := time.Now()
	bow := p.BeginningOfYear(now)
	eow := p.EndOfYear(now)

	if p.From().Unix() != bow.Unix() {
		t.Errorf("Expected start of this year: %d, got %d", bow.Unix(), p.To().Unix())
	}
	if p.To().Unix() != eow.Unix() {
		t.Errorf("Expected end of this year: %d, got %d", eow.Unix(), p.To().Unix())
	}
	if p.Period() != "This Year" {
		t.Errorf("Expected correct period string, got %s", p.Period())
	}
}

func TestPeriodLastWeek(t *testing.T) {
	p := period.Parse("1w")

	now := time.Now()
	bow := p.BeginningOfPreviousWeek(now)
	eow := p.EndOfPreviousWeek(now)

	if p.From().Unix() != bow.Unix() {
		t.Errorf("Expected start of last week: %d, got %d", bow.Unix(), p.To().Unix())
	}
	if p.To().Unix() != eow.Unix() {
		t.Errorf("Expected end of last week: %d, got %d", eow.Unix(), p.To().Unix())
	}
	if p.Period() != "Last Week" {
		t.Errorf("Expected correct period string, got %s", p.Period())
	}
}

func TestPeriodLastMonth(t *testing.T) {
	p := period.Parse("1m")

	now := time.Now()
	bow := p.BeginningOfPreviousMonth(now)
	eow := p.EndOfPreviousMonth(now)

	if p.From().Unix() != bow.Unix() {
		t.Errorf("Expected start of last month: %d, got %d", bow.Unix(), p.To().Unix())
	}
	if p.To().Unix() != eow.Unix() {
		t.Errorf("Expected end of last month: %d, got %d", eow.Unix(), p.To().Unix())
	}
	if p.Period() != "Last Month" {
		t.Errorf("Expected correct period string, got %s", p.Period())
	}
}

func TestPeriodLastYear(t *testing.T) {
	p := period.Parse("1y")

	now := time.Now()
	bow := p.BeginningOfPreviousYear(now)
	eow := p.EndOfPreviousYear(now)

	if p.From().Unix() != bow.Unix() {
		t.Errorf("Expected start of laat year: %d, got %d", bow.Unix(), p.To().Unix())
	}
	if p.To().Unix() != eow.Unix() {
		t.Errorf("Expected end of last year: %d, got %d", eow.Unix(), p.To().Unix())
	}
	if p.Period() != "Last Year" {
		t.Errorf("Expected correct period string, got %s", p.Period())
	}
}

func TestBeginningOfDay(t *testing.T) {
	p := period.Parse("")

	timeNow, _ := time.Parse("2006-01-02 15:04:05", "2019-01-05 14:24:01")
	actual, _ := time.Parse("2006-01-02 15:04:05", "2019-01-05 00:00:00")

	bod := p.BeginningOfDay(timeNow)

	if bod.String() != actual.String() {
		t.Errorf("Expected start of the day, got %s", bod.String())
	}
}

func TestEndOfDay(t *testing.T) {
	p := period.Parse("")

	timeNow, _ := time.Parse("2006-01-02 15:04:05", "2019-01-04 08:11:43")
	actual, _ := time.Parse("2006-01-02 15:04:05", "2019-01-04 23:59:59")

	eod := p.EndOfDay(timeNow)

	if eod.String() != actual.String() {
		t.Errorf("Expected end of the day, got %s", eod.String())
	}
}

func TestYesterday(t *testing.T) {
	p := period.Parse("")

	timeNow, _ := time.Parse("2006-01-02 15:04:05", "2019-04-01 08:11:43")
	actual, _ := time.Parse("2006-01-02 15:04:05", "2019-03-31 08:11:43")

	yesterday := p.Yesterday(timeNow)

	if yesterday.String() != actual.String() {
		t.Errorf("Expected day to be yesterday, got %s", yesterday.String())
	}
}

func TestBeginningOfWeek(t *testing.T) {
	p := period.Parse("")

	// Friday 4th January
	timeNow, _ := time.Parse("2006-01-02 15:04:05", "2019-01-04 14:24:01")

	// Monday 31st December
	actual, _ := time.Parse("2006-01-02 15:04:05", "2018-12-31 00:00:00")

	bow := p.BeginningOfWeek(timeNow)

	if bow.String() != actual.String() {
		t.Errorf("Expected beginning of week, got %s", bow.String())
	}
}

func TestEndOfWeek(t *testing.T) {
	p := period.Parse("")

	// Friday 4th January
	timeNow, _ := time.Parse("2006-01-02 15:04:05", "2019-01-04 14:24:01")

	// Sunday 6th January
	actual, _ := time.Parse("2006-01-02 15:04:05", "2019-01-06 23:59:59")

	eow := p.EndOfWeek(timeNow)

	if eow.String() != actual.String() {
		t.Errorf("Expected end of week, got %s", eow.String())
	}
}

func TestBeginningOfPreviousWeek(t *testing.T) {
	p := period.Parse("")

	// Wednesday 12th September
	timeNow, _ := time.Parse("2006-01-02 15:04:05", "2018-09-12 12:02:34")

	// Monday 3rd September
	actual, _ := time.Parse("2006-01-02 15:04:05", "2018-09-03 00:00:00")

	bow := p.BeginningOfPreviousWeek(timeNow)

	if bow.String() != actual.String() {
		t.Errorf("Expected beginning of last week, got %s", bow.String())
	}
}

func TestEndOfPreviousWeek(t *testing.T) {
	p := period.Parse("")

	// Wednesday 12th September
	timeNow, _ := time.Parse("2006-01-02 15:04:05", "2018-09-12 12:02:34")

	// Sunday 9th September
	actual, _ := time.Parse("2006-01-02 15:04:05", "2018-09-09 23:59:59")

	eow := p.EndOfPreviousWeek(timeNow)

	if eow.String() != actual.String() {
		t.Errorf("Expected end of last week, got %s", eow.String())
	}
}

func TestBeginningOfMonth(t *testing.T) {
	p := period.Parse("")

	timeNow, _ := time.Parse("2006-01-02 15:04:05", "2018-10-10 14:24:01")
	actual, _ := time.Parse("2006-01-02 15:04:05", "2018-10-01 00:00:00")

	bom := p.BeginningOfMonth(timeNow)

	if bom.String() != actual.String() {
		t.Errorf("Expected beginning of this month, got %s", bom.String())
	}
}
func TestEndOfMonth(t *testing.T) {
	p := period.Parse("")

	timeNow, _ := time.Parse("2006-01-02 15:04:05", "2019-01-09 14:24:01")
	actual, _ := time.Parse("2006-01-02 15:04:05", "2019-01-31 23:59:59")

	eom := p.EndOfMonth(timeNow)

	if eom.String() != actual.String() {
		t.Errorf("Expected end of tis month, got %s", eom.String())
	}
}

func TestBeginningOfPreviousMonth(t *testing.T) {
	p := period.Parse("")

	timeNow, _ := time.Parse("2006-01-02 15:04:05", "2019-01-09 14:24:01")
	actual, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 00:00:00")

	bom := p.BeginningOfPreviousMonth(timeNow)

	if bom.String() != actual.String() {
		t.Errorf("Expected beginning of previous month, got %s", bom.String())
	}
}

func TestEndOfPreviousMonth(t *testing.T) {
	p := period.Parse("")

	timeNow, _ := time.Parse("2006-01-02 15:04:05", "2019-01-09 14:24:01")
	actual, _ := time.Parse("2006-01-02 15:04:05", "2018-12-31 23:59:59")

	eom := p.EndOfPreviousMonth(timeNow)

	if eom.String() != actual.String() {
		t.Errorf("Expected end of previous month, got %s", eom.String())
	}
}

func TestBeginningOfYear(t *testing.T) {
	p := period.Parse("")

	timeNow, _ := time.Parse("2006-01-02 15:04:05", "2019-01-09 14:24:01")
	actual, _ := time.Parse("2006-01-02 15:04:05", "2019-01-01 00:00:00")

	boy := p.BeginningOfYear(timeNow)

	if boy.String() != actual.String() {
		t.Errorf("Expected beginning of year, got %s", boy.String())
	}
}

func TestEndOfYear(t *testing.T) {
	p := period.Parse("")

	timeNow, _ := time.Parse("2006-01-02 15:04:05", "2018-01-09 14:24:01")
	actual, _ := time.Parse("2006-01-02 15:04:05", "2018-12-31 23:59:59")

	eoy := p.EndOfYear(timeNow)

	if eoy.String() != actual.String() {
		t.Errorf("Expected end of year, got %s", eoy.String())
	}
}

func TestBeginningOfPreviousYear(t *testing.T) {
	p := period.Parse("")

	timeNow, _ := time.Parse("2006-01-02 15:04:05", "2019-01-09 14:24:01")
	actual, _ := time.Parse("2006-01-02 15:04:05", "2018-01-01 00:00:00")

	boy := p.BeginningOfPreviousYear(timeNow)

	if boy.String() != actual.String() {
		t.Errorf("Expected beginning of previous year, got %s", boy.String())
	}
}

func TestEndOfPreviousYear(t *testing.T) {
	p := period.Parse("")

	timeNow, _ := time.Parse("2006-01-02 15:04:05", "2019-01-09 14:24:01")
	actual, _ := time.Parse("2006-01-02 15:04:05", "2018-12-31 23:59:59")

	eoy := p.EndOfPreviousYear(timeNow)

	if eoy.String() != actual.String() {
		t.Errorf("Expected end of previous year, got %s", eoy.String())
	}
}
