package timeslip

import (
	"fmt"
)

type SlipTime struct {
	Hours   int
	Minutes int
	Seconds int
}

func (t *SlipTime) FromSeconds(seconds int) {
	t.Hours = seconds / 3600

	remainder := seconds % 3600

	t.Minutes = remainder / 60
	t.Seconds = remainder % 60
}

func (t SlipTime) String() string {
	if t.Hours > 0 && t.Minutes > 0 {
		return fmt.Sprintf("%dh %dm", t.Hours, t.Minutes)
	} else if t.Minutes > 0 && t.Seconds > 0 {
		return fmt.Sprintf("%dm %ds", t.Minutes, t.Seconds)
	}

	if t.Hours > 0 {
		return fmt.Sprintf("%d hours", t.Hours)
	} else if t.Minutes > 0 {
		return fmt.Sprintf("%d minutes", t.Minutes)
	}

	return fmt.Sprintf("%d seconds", t.Seconds)
}
