package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type TestCase struct {
	Arr      []int32 `json:"arr"`
	Expected []int   `json:"expected"`
}

func loadTestData(filePath string) ([]TestCase, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var testCases []TestCase
	err = json.NewDecoder(file).Decode(&testCases)
	if err != nil {
		return nil, err
	}

	return testCases, nil
}

func TestMiniMaxSum(t *testing.T) {
	testCases, err := loadTestData("../../test_data/min_max_four_sum.json")
	if err != nil {
		t.Fatalf("Failed to load test data: %v", err)
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("arr=%v", tc.Arr), func(t *testing.T) {
			gotMin, gotMax := minMaxFourSum(tc.Arr)
			if gotMin != tc.Expected[0] || gotMax != tc.Expected[1] {
				t.Errorf("minMaxFourSum(%v) = %v %v; want %v %v", tc.Arr, gotMin, gotMax, tc.Expected[0], tc.Expected[1])
			}
		})
	}
}
