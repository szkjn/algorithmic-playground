package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type TestCase struct {
	Year     int32  `json:"year"`
	Expected string `json:"expected"`
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

func TestDayOfProgrammer(t *testing.T) {
	testCases, err := loadTestData("../../test_data/day_of_programmer.json")
	if err != nil {
		t.Fatalf("Failed to load test data: %v", err)
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("s=%v", tc.Year), func(t *testing.T) {
			got := dayOfProgrammer(tc.Year)
			if got != tc.Expected {
				t.Errorf("dayOfProgrammer(%v) = %v; want %v", tc.Year, got, tc.Expected)
			}
		})
	}
}
