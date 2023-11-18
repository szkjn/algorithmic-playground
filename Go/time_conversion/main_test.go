package main

import (
	"fmt"
	"testing"
)

func TestTimeConversion(t *testing.T) {
	testCases := []struct {
		s        string
		expected string
	}{
		{"12:00:00AM", "00:00:00"},
		{"12:00:00PM", "12:00:00"},
		{"12:45:54PM", "12:45:54"},
		{"07:05:45PM", "19:05:45"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("s=%v", tc.s), func(t *testing.T) {
			got := timeConversion(tc.s)
			if got != tc.expected {
				t.Errorf("timeConversion(%v) = %v; want %v", tc.s, got, tc.expected)
			}
		})
	}
}
