// Basic file helper functions.

package utils

import (
	"os"
)

// WriteFile replaces/creates a file with provided data.
func WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

// AppendToFile adds a line to the end of a file (creates if missing).
func AppendToFile(path string, data string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(data + "\n")
	return err
}
