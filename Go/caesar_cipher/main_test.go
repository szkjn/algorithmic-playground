package main

import (
	"fmt"
	"testing"
)

func TestCaesarCipher(t *testing.T) {
	testCases := []struct {
		s        string
		k        int
		expected string
	}{
		{"middle-Outz", 2, "okffng-Qwvb"},
		{"www.abc.xy", 87, "fff.jkl.gh"},
		{"D3q4", 0, "D3q4"},
		{"Pz-/aI/J`EvfthGH", 66, "Dn-/oW/X`SjthvUV"},
		{
			"6DWV95HzxTCHP85dvv3NY2crzt1aO8j6g2zSDvFUiJj6XWDlZvNNr",
			87,
			"6MFE95QigCLQY85mee3WH2laic1jX8s6p2iBMeODrSs6GFMuIeWWa",
		},
		{
			"1X7T4VrCs23k4vv08D6yQ3S19G4rVP188M9ahuxB6j1tMGZs1m10ey7eUj62WV2exLT4C83zl7Q80M",
			27,
			"1Y7U4WsDt23l4ww08E6zR3T19H4sWQ188N9bivyC6k1uNHAt1n10fz7fVk62XW2fyMU4D83am7R80N",
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("s=%v,k=%d", tc.s, tc.k), func(t *testing.T) {
			got := caesarCipher(tc.s, tc.k)
			if got != tc.expected {
				t.Errorf("caesarCipher(%v, %d) = %v; want %v", tc.s, tc.k, got, tc.expected)
			}
		})
	}
}
