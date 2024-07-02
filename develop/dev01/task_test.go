package main

import (
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	ntp, err := GetTime()
	if err != nil {
		t.Errorf("Can't take current time: %s", err)
	}
	if ntp.Second() != time.Now().Second() {
		t.Errorf("Wrong time, expected %d, got %d", time.Now().Second(), ntp.Second())
	}
	if ntp.Minute() != time.Now().Minute() {
		t.Errorf("Wrong time, expected %d, got %d", time.Now().Minute(), ntp.Minute())
	}
	if ntp.Hour() != time.Now().Hour() {
		t.Errorf("Wrong time, expected %d, got %d", time.Now().Hour(), ntp.Hour())
	}
	if ntp.Day() != time.Now().Day() {
		t.Errorf("Wrong time, expected %d, got %d", time.Now().Day(), ntp.Day())
	}
	if ntp.Month() != time.Now().Month() {
		t.Errorf("Wrong time, expected %d, got %d", time.Now().Month(), ntp.Month())
	}
	if ntp.Year() != time.Now().Year() {
		t.Errorf("Wrong time, expected %d, got %d", time.Now().Year(), ntp.Year())
	}
}
