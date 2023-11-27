package test_utils

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func LoadTestData(relativePath string) ([]map[string]interface{}, error) {
	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []map[string]interface{}
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
