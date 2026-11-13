package tests

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/yourusername/csvtojson/config"
	"github.com/yourusername/csvtojson/converter"
	"github.com/yourusername/csvtojson/utils"
)

func TestConverter_EndToEnd(t *testing.T) {
	// Prepare temporary directory
	tmpDir := t.TempDir()
	csvPath := filepath.Join(tmpDir, "sample.csv")
	jsonPath := filepath.Join(tmpDir, "result.json")

	csvContent := `id,name,age
1,Alice,30
2,Bob,25
3,Charlie,35
`

	if err := utils.WriteFile(csvPath, []byte(csvContent)); err != nil {
		t.Fatalf("failed to write CSV: %v", err)
	}

	// Load config with default values
	cfg := config.LoadConfig()
	logger := utils.NewLogger("error") // silence output during test

	conv := converter.NewConverter(cfg, logger)
	if err := conv.Convert(csvPath, jsonPath); err != nil {
		t.Fatalf("conversion failed: %v", err)
	}

	// Verify JSON file exists
	if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
		t.Fatalf("expected JSON output file, but it does not exist")
	}

	// Load and unmarshal JSON
	data, err := utils.ReadFile(jsonPath)
	if err != nil {
		t.Fatalf("failed to read JSON output: %v", err)
	}

	var records []map[string]string
	if err := json.Unmarshal(data, &records); err != nil {
		t.Fatalf("invalid JSON output: %v", err)
	}

	if len(records) != 3 {
		t.Fatalf("expected 3 records, got %d", len(records))
	}
	if records[0]["name"] != "Alice" || records[1]["age"] != "25" {
		t.Fatalf("record contents mismatch: %+v", records)
	}
}