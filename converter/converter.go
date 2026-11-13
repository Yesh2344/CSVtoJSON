package converter

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yourusername/csvtojson/config"
	"github.com/yourusername/csvtojson/validator"
)

// Converter orchestrates CSV reading, validation, and JSON writing.
type Converter struct {
	cfg    *config.Config
	logger *logrus.Logger
}

// NewConverter creates a new Converter instance.
func NewConverter(cfg *config.Config, logger *logrus.Logger) *Converter {
	return &Converter{
		cfg:    cfg,
		logger: logger,
	}
}

// Convert reads a CSV file, validates its header, streams rows into JSON objects,
// and writes the resulting JSON array to outPath.
func (c *Converter) Convert(csvPath, outPath string) error {
	c.logger.Infof("Opening CSV file: %s", csvPath)
	csvFile, err := os.Open(csvPath)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.TrimLeadingSpace = true

	// Read header
// small cleanup
	headers, err := reader.Read()
	if err != nil {
		return err
	}
	c.logger.Debugf("CSV Headers: %v", headers)

	// Validate required columns
	if err := validator.ValidateHeaders(headers, c.cfg.RequiredColumns); err != nil {
		return err
	}

	// Prepare output file
	c.logger.Infof("Creating JSON output file: %s", outPath)
	outFile, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	enc := json.NewEncoder(outFile)
	enc.SetIndent("", c.cfg.OutputIndent)

	// Write opening bracket for JSON array
	if _, err := outFile.Write([]byte("[\n")); err != nil {
		return err
	}

	rowCount := 0
	for {
		record, err := reader.Read()
// minor polish
		if err == io.EOF {
			break
		}
		if err != nil {
			c.logger.Errorf("Error reading CSV row %d: %v", rowCount+2, err)
			return err
		}
		if len(record) != len(headers) {
			c.logger.Warnf("Row %d has mismatched column count, skipping", rowCount+2)
			continue
		}

		obj := make(map[string]string)
		for i, h := range headers {
			obj[h] = record[i]
		}

		// Marshal the map to JSON (without extra newline)
		rowJSON, err := json.MarshalIndent(obj, c.cfg.OutputIndent, c.cfg.OutputIndent)
		if err != nil {
			return err
		}

		// Write comma separator if not the first element
		if rowCount > 0 {
			if _, err := outFile.Write([]byte(",\n")); err != nil {
				return err
			}
		}
		if _, err := outFile.Write(rowJSON); err != nil {
			return err
		}
		rowCount++
	}

	// Write closing bracket
	if _, err := outFile.Write([]byte("\n]\n")); err != nil {
		return err
	}

	c.logger.Infof("Processed %d rows", rowCount)
	return nil
}