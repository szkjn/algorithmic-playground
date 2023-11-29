package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type TestCase struct {
	Scores   []int32 `json:"scores"`
	Expected []int32 `json:"expected"`
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

func TestRecordBreakCounter(t *testing.T) {
	testCases, err := loadTestData("../../test_data/record_break_counter.json")
	if err != nil {
		t.Fatalf("Failed to load test data: %v", err)
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Scores=%v", tc.Scores), func(t *testing.T) {
			got := recordBreakCounter(tc.Scores)
			gotMost := got[0]
			gotLeast := got[1]
			if gotMost != tc.Expected[0] || gotLeast != tc.Expected[1] {
				t.Errorf("recordBreakCounter(%v) = %v %v; want %v %v", tc.Scores, gotMost, gotLeast, tc.Expected[0], tc.Expected[1])
			}
		})
	}
}
