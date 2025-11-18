package output

import (
	"encoding/json"
	"os"
)

func GenerateJSONReport(path string, data map[string]string) error {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, jsonData, 0644)
}
