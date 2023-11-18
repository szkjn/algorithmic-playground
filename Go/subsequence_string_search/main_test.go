package main

import (
	"fmt"
	"testing"
)

func TestHackerrankInString(t *testing.T) {
	testCases := []struct {
		s        string
		expected string
	}{
		{"hereiamstackerrank", "YES"},
		{"hackerworld", "NO"},
		{"hhaacckkekraraannk", "YES"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("s=%v", tc.s), func(t *testing.T) {
			got := hackerrankInString(tc.s)
			if got != tc.expected {
				t.Errorf("hackerrankInString(%v) = %v; want %v", tc.s, got, tc.expected)
			}
		})
	}
}
