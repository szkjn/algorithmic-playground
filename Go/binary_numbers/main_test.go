package main

import (
	"fmt"
	"testing"
)

func TestCountMaxConsecutiveOnes(t *testing.T) {
	testCases := []struct {
		n        string
		expected int
	}{
		{"101", 1},
		{"110", 2},
		{"1101", 2},
		{"1111111111111110011", 15},
		{"1111111111111111011", 16},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("n=%s", tc.n), func(t *testing.T) {
			got := countMaxConsecutiveOnes(tc.n)
			if got != tc.expected {
				t.Errorf("countMaxConsecutiveOnes(%s) = %d; want %d", tc.n, got, tc.expected)
			}
		})
	}
}
