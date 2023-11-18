package main

import (
	"fmt"
	"testing"
)

func TestSumEveryOther(t *testing.T) {
	testCases := []struct {
		n        float64
		expected int
	}{
		{548915381, 26},
		{10, 0},
		{1010.11, 1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("n=%v", tc.n), func(t *testing.T) {
			got := sumEveryOther(tc.n)
			if got != tc.expected {
				t.Errorf("sumEveryOther(%v) = %v; want %v", tc.n, got, tc.expected)
			}
		})
	}
}
