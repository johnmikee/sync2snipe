package helpers

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	year := 2023
	month := 5
	day := 29
	expected := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	result := Date(year, month, day)

	if result != expected {
		t.Errorf("Date failed: expected '%s', got '%s'", expected, result)
	}
}

func TestStringToTime(t *testing.T) {
	timeString := "2023-05-29T12:34:56Z"
	expected, _ := time.Parse(time.RFC3339, timeString)
	result, err := StringToTime(timeString)

	if err != nil {
		t.Errorf("StringToTime failed: unexpected error '%s'", err)
	}

	if result != expected {
		t.Errorf("StringToTime failed: expected '%s', got '%s'", expected, result)
	}
}

func TestUpdateTime(t *testing.T) {
	expected := time.Now().UTC()
	result := UpdateTime()

	// Allow a small time difference for comparison due to precision limitations
	timeDiff := result.Sub(expected)
	if timeDiff < 0 {
		timeDiff = -timeDiff
	}

	// We consider a difference of up to 1 second acceptable
	acceptableDiff := time.Second
	if timeDiff > acceptableDiff {
		t.Errorf("UpdateTime failed: expected '%s', got '%s'", expected, result)
	}
}
