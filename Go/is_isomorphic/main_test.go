package main

import (
	"fmt"
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	testCases := []struct {
		s, t     string
		expected bool
	}{
		{"abb", "cdd", true},
		{"cassidy", "1234567", false},
		{"cass", "1233", true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("s=%v, t=%v", tc.s, tc.t), func(t *testing.T) {
			got := isIsomorphic(tc.s, tc.t)
			if got != tc.expected {
				t.Errorf("isIsomorphic(%v, %v) = %v; want %v", tc.s, tc.t, got, tc.expected)
			}
		})
	}
}
