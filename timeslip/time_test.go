package timeslip_test

import (
	"testing"

	"github.com/mrcook/time_warrior/timeslip"
)

func TestParseSeconds(t *testing.T) {
	st := timeslip.SlipTime{}

	st.FromSeconds(7538) // 2h 5m 38s

	if st.Hours != 2 {
		t.Errorf("Expected 2 hours, got '%d'", st.Hours)
	}
	if st.Minutes != 5 {
		t.Errorf("Expected 5 minutes, got '%d'", st.Minutes)
	}
	if st.Seconds != 38 {
		t.Errorf("Expected 38 seconds, got '%d'", st.Seconds)
	}
}

func TestStringSeconds(t *testing.T) {
	st := timeslip.SlipTime{}
	st.FromSeconds(55)
	if st.String() != "55 seconds" {
		t.Errorf("Expected '55 seconds' to be returned, got '%s'", st.String())
	}
}

func TestStringMinutes(t *testing.T) {
	st := timeslip.SlipTime{}
	st.FromSeconds(240)
	if st.String() != "4 minutes" {
		t.Errorf("Expected '4 minutes' to be returned, got '%s'", st.String())
	}
}

func TestStringHours(t *testing.T) {
	st := timeslip.SlipTime{}
	st.FromSeconds(10800)
	if st.String() != "3 hours" {
		t.Errorf("Expected '3 hours' to be returned, got '%s'", st.String())
	}
}

func TestStringHoursMinutes(t *testing.T) {
	st := timeslip.SlipTime{}
	st.FromSeconds(3600 + 1380)
	if st.String() != "1h 23m" {
		t.Errorf("Expected '1h 23m' to be returned, got '%s'", st.String())
	}
}

func TestStringMinutesSeconds(t *testing.T) {
	st := timeslip.SlipTime{}
	st.FromSeconds(720 + 45)
	if st.String() != "12m 45s" {
		t.Errorf("Expected '12m 45s' to be returned, got '%s'", st.String())
	}
}
