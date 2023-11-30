package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type TestCase struct {
	S        string `json:"s"`
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

func TestDecodeMorse(t *testing.T) {
	testCases, err := loadTestData("../../test_data/decode_morse_code.json")
	if err != nil {
		t.Fatalf("Failed to load test data: %v", err)
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("s=%v", tc.S), func(t *testing.T) {
			got := decodeMorse(tc.S)
			if got != tc.Expected {
				t.Errorf("decodeMorse(%v) = %v; want %v", tc.S, got, tc.Expected)
			}
		})
	}
}
