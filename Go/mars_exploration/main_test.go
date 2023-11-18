package main

import (
	"fmt"
	"testing"
)

func TestMarsExploration(t *testing.T) {
	testCases := []struct {
		s        string
		expected int32
	}{
		{"SOSSPSSQSSOR", 3},
		{"SOSSOT", 1},
		{"SOSSOSSOS", 0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("s=%v", tc.s), func(t *testing.T) {
			got := marsExploration(tc.s)
			if got != tc.expected {
				t.Errorf("marsExploration(%v) = %v; want %v", tc.s, got, tc.expected)
			}
		})
	}
}
