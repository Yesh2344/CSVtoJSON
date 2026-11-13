package validator

import (
	"fmt"
	"strings"
)

// ValidateHeaders ensures that all required columns exist in the CSV header.
// Returns an error describing missing columns.
func ValidateHeaders(headers []string, required []string) error {
	headerSet := make(map[string]struct{}, len(headers))
	for _, h := range headers {
		headerSet[strings.TrimSpace(h)] = struct{}{}
	}

	var missing []string
	for _, r := range required {
		if _, ok := headerSet[r]; !ok {
			missing = append(missing, r)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing required columns: %s", strings.Join(missing, ", "))
	}
	return nil
}