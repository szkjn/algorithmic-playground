package main

import (
	"fmt"
	"testing"
)

func TestColumnNumber(t *testing.T) {
	testCases := []struct {
		s        string
		expected int
	}{
		{"A", 1},
		{"B", 2},
		{"C", 3},
		{"Z", 26},
		{"AA", 27},
		{"AB", 28},
		{"AAA", 703},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("s=%v", tc.s), func(t *testing.T) {
			got := columnNumber(tc.s)
			if got != tc.expected {
				t.Errorf("columnNumber(%v) = %v; want %v", tc.s, got, tc.expected)
			}
		})
	}
}
