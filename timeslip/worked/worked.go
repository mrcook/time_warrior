// Package worked provides helper time functions for calculating and
// displaying the time worked on a timeslip.
package worked

import (
	"fmt"
	"strconv"
	"strings"
)

// WorkTime represents time in Hours, Minutes, Seconds to make working
// with the timeslips UNIX representation of time easier.
type WorkTime struct {
	Hours   int
	Minutes int
	Seconds int
}

// FromHours sets the numbers of hours
func (w *WorkTime) FromHours(hours int) {
	w.Hours = hours
}

// FromMinutes sets the numbers of minutes
func (w *WorkTime) FromMinutes(minutes int) {
	w.Hours = minutes / 60
	w.Minutes = minutes % 60
}

// FromSeconds sets the numbers of seconds
func (w *WorkTime) FromSeconds(seconds int) {
	w.Hours = seconds / 3600

	remainder := seconds % 3600

	w.Minutes = remainder / 60
	w.Seconds = remainder % 60
}

// FromString parses an adjust command string, e.g. `-75m`.
func (w *WorkTime) FromString(adjustment string) error {
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
		w.FromHours(h)
	case "m":
		m, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("unable to process input")
		}
		w.FromMinutes(m)
	case "s":
		s, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("unable to process input")
		}
		w.FromSeconds(s)
	default:
		return fmt.Errorf("unable to process input")
	}

	if negative {
		w.Hours = -w.Hours
		w.Minutes = -w.Minutes
		w.Seconds = -w.Seconds
	}

	return nil
}

// String returns the worked time as a string, e.g. `1h 10m`.
func (w WorkTime) String() string {
	if w.Hours != 0 && w.Minutes != 0 {
		return fmt.Sprintf("%dh %dm", w.Hours, w.Minutes)
	} else if w.Minutes != 0 && w.Seconds != 0 {
		return fmt.Sprintf("%dm %ds", w.Minutes, w.Seconds)
	}

	if w.Hours != 0 {
		return fmt.Sprintf("%d hours", w.Hours)
	} else if w.Minutes != 0 {
		return fmt.Sprintf("%d minutes", w.Minutes)
	}

	return fmt.Sprintf("%d seconds", w.Seconds)
}

// ToSeconds returns the worked time in seconds.
func (w WorkTime) ToSeconds() int {
	hours := w.Hours * 60 * 60
	minutes := w.Minutes * 60

	return hours + minutes + w.Seconds
}

// Add one worked time to another.
func (w *WorkTime) Add(nu *WorkTime) {
	seconds := w.ToSeconds() + nu.ToSeconds()
	w.FromSeconds(seconds)
}

// Subtract one worked time from another.
func (w *WorkTime) Subtract(nu *WorkTime) {
	seconds := w.ToSeconds() - nu.ToSeconds()
	w.FromSeconds(seconds)
}
