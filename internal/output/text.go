package output

import (
	"os"
	"strings"
)

func GenerateTextReport(path string, data map[string]string) error {
	var builder strings.Builder

	for k, v := range data {
		builder.WriteString(k + ": " + v + "\n")
	}

	return os.WriteFile(path, []byte(builder.String()), 0644)
}
