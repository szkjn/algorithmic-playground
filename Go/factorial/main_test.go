package main

import (
	"fmt"
	"testing"
)

func TestFactorial(t *testing.T) {
	testCases := []struct {
		n        int32
		expected int32
	}{
		{3, 6},
		{1, 1},
		{5, 120},
		{11, 39916800},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("n=%d", tc.n), func(t *testing.T) {
			got := factorial(tc.n)
			if got != tc.expected {
				t.Errorf("factorial(%d) = %d; want %d", tc.n, got, tc.expected)
			}
		})
	}
}
