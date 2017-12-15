package timeslip_test

import (
	"testing"

	"github.com/mrcook/time_warrior/timeslip"
)

func TestFromHours(t *testing.T) {
	st := timeslip.SlipTime{}

	st.FromHours(4)

	if st.Hours != 4 || st.Minutes != 0 || st.Seconds != 0 {
		t.Errorf("Expected 4h 0m 0s, got %dh %dm %ds", st.Hours, st.Minutes, st.Seconds)
	}
}

func TestFromMinutes(t *testing.T) {
	st := timeslip.SlipTime{}

	st.FromMinutes(199) // 2h 19m

	if st.Hours != 3 || st.Minutes != 19 || st.Seconds != 0 {
		t.Errorf("Expected 3h 19m 0s, got %dh %dm %ds", st.Hours, st.Minutes, st.Seconds)
	}
}

func TestFromSeconds(t *testing.T) {
	st := timeslip.SlipTime{}

	st.FromSeconds(7538) // 2h 5m 38s

	if st.Hours != 2 || st.Minutes != 5 || st.Seconds != 38 {
		t.Errorf("Expected 2h 5m 38s, got %dh %dm %ds", st.Hours, st.Minutes, st.Seconds)
	}
}

func TestFromInvalidString(t *testing.T) {
	st := timeslip.SlipTime{}

	err := st.FromString("3h 52m 18s")
	if err.Error() != "invalid string, expected one time unit, got 3" {
		t.Error("Expected an invalid string error")
	}

	err = st.FromString("3h52m")
	if err.Error() != "unable to process input" {
		t.Error("Expected an unprocessible string error")
	}
}

func TestFromHourString(t *testing.T) {
	st := timeslip.SlipTime{}

	err := st.FromString("7h")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if st.Hours != 7 || st.Minutes != 0 || st.Seconds != 0 {
		t.Errorf("Expected 7h 0m 0s, got %dh %dm %ds", st.Hours, st.Minutes, st.Seconds)
	}

	err = st.FromString("-6h")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if st.Hours != -6 || st.Minutes != 0 || st.Seconds != 0 {
		t.Errorf("Expected -6h 0m 0s, got %dh %dm %ds", st.Hours, st.Minutes, st.Seconds)
	}
}

func TestFromMinuteString(t *testing.T) {
	st := timeslip.SlipTime{}

	err := st.FromString("74m")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if st.Hours != 1 || st.Minutes != 14 || st.Seconds != 0 {
		t.Errorf("Expected -1h 14m 0s, got %dh %dm %ds", st.Hours, st.Minutes, st.Seconds)
	}

	err = st.FromString("-81m") // 1h 21m
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if st.Hours != -1 || st.Minutes != -21 || st.Seconds != 0 {
		t.Errorf("Expected -1h 21m 0s, got %dh %dm %ds", st.Hours, st.Minutes, st.Seconds)
	}
}

func TestFromSecondString(t *testing.T) {
	st := timeslip.SlipTime{}

	err := st.FromString("3727s")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if st.Hours != 1 || st.Minutes != 2 || st.Seconds != 7 {
		t.Errorf("Expected 1h 2m 7s, got %dh %dm %ds", st.Hours, st.Minutes, st.Seconds)
	}

	err = st.FromString("-14592s") // 4h 3m 12s
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if st.Hours != -4 || st.Minutes != -3 || st.Seconds != -12 {
		t.Errorf("Expected -4h 3m 12s, got %dh %dm %ds", st.Hours, st.Minutes, st.Seconds)
	}
}

func TestStringHours(t *testing.T) {
	st := timeslip.SlipTime{}

	st.FromSeconds(10800)
	if st.String() != "3 hours" {
		t.Errorf("Expected '3 hours' to be returned, got '%s'", st.String())
	}

	st.FromSeconds(-3600)
	if st.String() != "-1 hours" {
		t.Errorf("Expected '-1 hours' to be returned, got '%s'", st.String())
	}
}

func TestStringMinutes(t *testing.T) {
	st := timeslip.SlipTime{}

	st.FromSeconds(240)
	if st.String() != "4 minutes" {
		t.Errorf("Expected '4 minutes' to be returned, got '%s'", st.String())
	}

	st.FromSeconds(-120)
	if st.String() != "-2 minutes" {
		t.Errorf("Expected '-2 minutes' to be returned, got '%s'", st.String())
	}
}

func TestStringSeconds(t *testing.T) {
	st := timeslip.SlipTime{}

	st.FromSeconds(55)
	if st.String() != "55 seconds" {
		t.Errorf("Expected '55 seconds' to be returned, got '%s'", st.String())
	}

	st.FromSeconds(-20)
	if st.String() != "-20 seconds" {
		t.Errorf("Expected '-20 seconds' to be returned, got '%s'", st.String())
	}
}

func TestStringHoursMinutes(t *testing.T) {
	st := timeslip.SlipTime{}

	st.FromSeconds(7200 + 240)
	if st.String() != "2h 4m" {
		t.Errorf("Expected '2h 4m' to be returned, got '%s'", st.String())
	}

	st.FromSeconds(-3600 - 1380)
	if st.String() != "-1h -23m" {
		t.Errorf("Expected '-1h -23m' to be returned, got '%s'", st.String())
	}
}

func TestStringMinutesSeconds(t *testing.T) {
	st := timeslip.SlipTime{}

	st.FromSeconds(360 + 13)
	if st.String() != "6m 13s" {
		t.Errorf("Expected '6m 13s' to be returned, got '%s'", st.String())
	}

	st.FromSeconds(-720 - 45)
	if st.String() != "-12m -45s" {
		t.Errorf("Expected '-12m -45s' to be returned, got '%s'", st.String())
	}
}
