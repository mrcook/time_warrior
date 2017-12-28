// Package worked provides helper time functions for calculating and
// displaying the time worked on a timeslip.
package worked

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hours   int
	Minutes int
	Seconds int
}

func (t *Time) FromHours(hours int) {
	t.Hours = hours
}

func (t *Time) FromMinutes(minutes int) {
	t.Hours = minutes / 60
	t.Minutes = minutes % 60
}

func (t *Time) FromSeconds(seconds int) {
	t.Hours = seconds / 3600

	remainder := seconds % 3600

	t.Minutes = remainder / 60
	t.Seconds = remainder % 60
}

func (t *Time) FromString(adjustment string) error {
	adjustment = strings.TrimSpace(adjustment)

	units := strings.Split(adjustment, " ")
	if len(units) != 1 {
		return fmt.Errorf("invalid string, expected one time unit, got %d", len(units))
	}

	unit := adjustment[len(adjustment)-1:]
	value := adjustment[:len(adjustment)-1]

	negative := value[0] == '-'

	if value[0] == '-' || value[0] == '+' {
		value = adjustment[1 : len(adjustment)-1]
	}

	switch unit {
	case "h":
		h, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("unable to process input")
		}
		t.FromHours(h)
	case "m":
		m, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("unable to process input")
		}
		t.FromMinutes(m)
	case "s":
		s, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("unable to process input")
		}
		t.FromSeconds(s)
	default:
		return fmt.Errorf("unable to process input")
	}

	if negative {
		t.Hours = -t.Hours
		t.Minutes = -t.Minutes
		t.Seconds = -t.Seconds
	}

	return nil
}

func (t Time) String() string {
	if t.Hours != 0 && t.Minutes != 0 {
		return fmt.Sprintf("%dh %dm", t.Hours, t.Minutes)
	} else if t.Minutes != 0 && t.Seconds != 0 {
		return fmt.Sprintf("%dm %ds", t.Minutes, t.Seconds)
	}

	if t.Hours != 0 {
		return fmt.Sprintf("%d hours", t.Hours)
	} else if t.Minutes != 0 {
		return fmt.Sprintf("%d minutes", t.Minutes)
	}

	return fmt.Sprintf("%d seconds", t.Seconds)
}

func (t Time) ToSeconds() int {
	hours := t.Hours * 60 * 60
	minutes := t.Minutes * 60

	return hours + minutes + t.Seconds
}

func (t *Time) Add(nu *Time) {
	seconds := t.ToSeconds() + nu.ToSeconds()
	t.FromSeconds(seconds)
}

func (t *Time) Subtract(nu *Time) {
	seconds := t.ToSeconds() - nu.ToSeconds()
	t.FromSeconds(seconds)
}
