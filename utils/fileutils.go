package utils

import (
	"io/ioutil"
	"os"
)

// ReadFile returns the entire content of the file at path.
// It abstracts away error handling for callers that just need the bytes.
func ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// WriteFile writes data to the given path with 0644 permissions.
// If the file exists, it will be overwritten.
func WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}