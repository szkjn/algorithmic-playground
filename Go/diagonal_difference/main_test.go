package main

import (
	"fmt"
	"testing"
)

func TestDiagonalDifference(t *testing.T) {
	testCases := []struct {
		arr      [][]int32
		expected int32
	}{
		{
			[][]int32{
				{11, 2, 4},
				{4, 5, 6},
				{10, 8, -12},
			},
			15,
		},
		{
			[][]int32{
				{-1, 1, -7, -8},
				{-10, -8, -5, -2},
				{0, 9, 7, -1},
				{4, 4, -2, 1},
			},
			1,
		},
		{
			[][]int32{
				{-10, 3, 0, 5, -4},
				{2, -1, 0, 2, -8},
				{9, -2, -5, 6, 0},
				{9, -7, 4, 8, -2},
				{3, 7, 8, -5, 0},
			},
			3,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("arr=%v", tc.arr), func(t *testing.T) {
			got := diagonalDifference(tc.arr)
			if got != tc.expected {
				t.Errorf("diagonalDifference(%v) = %v; want %v", tc.arr, got, tc.expected)
			}
		})
	}
}
