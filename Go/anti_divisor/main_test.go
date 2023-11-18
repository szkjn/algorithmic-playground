package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsAntiDivisor(t *testing.T) {
	testCases := []struct {
		k, n     int
		expected bool
	}{
		{2, 234, false}, // 234 % 2 = 0, which is not equal to 2 / 2 (1)
		{3, 234, false}, // 234 % 3 = 0, which is not equal to (3 - 1) / 2 (1) or (3 + 1) / 2 (2)
		{4, 234, true},  // 234 % 4 = 2, which is equal to 4 / 2 (2)
		{5, 234, false}, // 234 % 5 = 4, which is not equal to (5 - 1) / 2 (2) or (5 + 1) / 2 (3)
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("k=%d,n=%d", tc.k, tc.n), func(t *testing.T) {
			got := isAntiDivisor(tc.k, tc.n)
			if got != tc.expected {
				t.Errorf("isAntiDivisor(%d, %d) = %v; want %v", tc.k, tc.n, got, tc.expected)
			}
		})
	}
}

func TestAntiDivisors(t *testing.T) {
	testCases := []struct {
		n        int
		expected []int
	}{
		{5, []int{2, 3}},
		{10, []int{3, 4, 7}},
		{105, []int{2, 6, 10, 11, 14, 19, 30, 42, 70}},
		{234, []int{4, 7, 12, 36, 52, 67, 156}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("n=%d", tc.n), func(t *testing.T) {
			got := antiDivisors(tc.n)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("antiDivisors(%d) = %v; want %v", tc.n, got, tc.expected)
			}
		})
	}
}
