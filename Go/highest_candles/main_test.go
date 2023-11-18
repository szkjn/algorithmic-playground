package main

import (
	"fmt"
	"testing"
)

func TestBirthdayCakeCandles(t *testing.T) {
	testCases := []struct {
		candles  []int32
		expected int32
	}{
		{[]int32{4, 4, 1, 3}, 2},
		{[]int32{8, 8, 8, 9, 8, 8}, 1},
		{[]int32{40, 2, 6, 18, 40, 5, 3, 40}, 3},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("candles=%v", tc.candles), func(t *testing.T) {
			got := birthdayCakeCandles(tc.candles)
			if got != tc.expected {
				t.Errorf("birthdayCakeCandles(%v) = %v; want %v", tc.candles, got, tc.expected)
			}
		})
	}
}
