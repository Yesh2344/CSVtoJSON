package main

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
	"github.com/yourusername/csvtojson/config"
	"github.com/yourusername/csvtojson/converter"
	"github.com/yourusername/csvtojson/utils"
)

func main() {
	// Load .env if present
	_ = godotenv.Load()

	// Parse CLI flags
	csvPath := flag.String("csv", "", "Path to input CSV file (required)")
	outPath := flag.String("out", "", "Path to output JSON file (required)")
	indent := flag.String("indent", "", "JSON indentation string (overrides config)")
	flag.Parse()

	if *csvPath == "" || *outPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Initialise logger
	logger := utils.NewLogger(cfg.LogLevel)
	logger.Infof("Starting CSV to JSON conversion")
	logger.Debugf("Config: %+v", cfg)

	// Override indent if flag provided
	if *indent != "" {
		cfg.OutputIndent = *indent
	}

	conv := converter.NewConverter(cfg, logger)
	if err := conv.Convert(*csvPath, *outPath); err != nil {
		logger.Fatalf("Conversion failed: %v", err)
	}
	logger.Infof("Successfully wrote JSON to %s", *outPath)
}