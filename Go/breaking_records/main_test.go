package main

import (
	"fmt"
	"testing"
)

func TestBreakingRecords(t *testing.T) {
	testCases := []struct {
		scores    []int32
		wantMost  int32
		wantLeast int32
	}{
		{
			[]int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42},
			int32(4),
			int32(0),
		},
		{
			[]int32{503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503, 503},
			int32(0),
			int32(0),
		},
		{
			[]int32{0, 9, 3, 10, 2, 20},
			int32(3),
			int32(0),
		},
		{
			[]int32{44316, 52125, 68761, 68787, 39275, 96092, 3320, 80485, 33978, 55037, 55142, 47170, 26255, 59084, 50342, 38609, 1015, 10553, 5292, 94856, 88112, 50516, 57426, 51488, 86656, 69595, 71558, 34143, 90351, 6500, 82497, 34667, 58626, 67611, 3454, 14253, 80055, 23126, 94738, 30386, 94516, 66232, 77556, 20771, 25316, 44251, 59380, 42683, 54804, 81024, 53891, 42916, 31540, 11318, 10757, 34549, 80913, 98667, 68692, 87616, 21519, 67542, 38635, 80145, 35153, 58441, 10750, 31560, 81568, 21840, 61946, 92436, 88072, 55855, 13207, 29741, 106, 88939, 88776, 71262, 86315, 42668, 30531, 17856, 53986, 57640, 52405, 51251, 56307, 37449, 38868, 77826, 21343, 77503, 74324, 72848, 52297, 85074, 4409, 33865, 6915, 66355, 26301, 11339, 22210, 55860, 57432, 38668, 61151, 46209, 9931, 63818, 88877, 56814, 81674, 42863, 14454, 50431, 10466, 70761, 4233, 49334},
			int32(5),
			int32(4),
		},
		{
			[]int32{100000000, 100000000, 10000000, 10000000, 1000000},
			int32(0),
			int32(2),
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("scores=%v", tc.scores), func(t *testing.T) {
			got := breakingRecords(tc.scores)
			gotMost := got[0]
			gotLeast := got[1]
			if gotMost != tc.wantMost || gotLeast != tc.wantLeast {
				t.Errorf("breakRecords(%v) = %v %v; want %v %v", tc.scores, gotMost, gotLeast, tc.wantMost, tc.wantLeast)
			}
		})
	}
}
