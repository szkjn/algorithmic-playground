package main

import (
	"fmt"
	"testing"
)

func TestDecodeMorse(t *testing.T) {
	testCases := []struct {
		s        string
		expected string
	}{
		{
			".... . -.--   .--- ..- -.. .",
			"HEY JUDE",
		},
		{
			"...---...",
			"SOS",
		},
		{
			"...   ---   ...",
			"S O S",
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("s=%v", tc.s), func(t *testing.T) {
			got := decodeMorse(tc.s)
			if got != tc.expected {
				t.Errorf("decodeMorse(%v) = %v; want %v", tc.s, got, tc.expected)
			}
		})
	}
}
