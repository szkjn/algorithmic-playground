package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"
)

type TestCase struct {
	N        int   `json:"input"`
	Expected []int `json:"expected"`
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

func TestAntiDivisors(t *testing.T) {
	testCases, err := loadTestData("../../test_data/anti_divisor_test_data.json")
	if err != nil {
		t.Fatalf("Failed to load test data: %v", err)
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("n=%d", tc.N), func(t *testing.T) {
			got := antiDivisors(tc.N)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Errorf("antiDivisors(%d) = %v; want %v", tc.N, got, tc.Expected)
			}
		})
	}
}
