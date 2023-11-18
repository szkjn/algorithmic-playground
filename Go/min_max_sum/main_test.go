package main

import (
	"fmt"
	"testing"
)

func TestMiniMaxSum(t *testing.T) {
	testCases := []struct {
		arr         []int32
		expectedMin int
		expectedMax int
	}{
		{[]int32{1, 3, 5, 7, 9}, 16, 24},
		{[]int32{4, 5, 6, 7, 8}, 22, 26},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("s=%v", tc.arr), func(t *testing.T) {
			gotMin, gotMax := miniMaxSum(tc.arr)
			if gotMin != tc.expectedMin || gotMax != tc.expectedMax {
				t.Errorf("miniMaxSum(%v) = %v %v; want %v %v", tc.arr, gotMin, gotMax, tc.expectedMin, tc.expectedMax)
			}
		})
	}
}
