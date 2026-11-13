# CSVtoJSON

[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/csvtojson)](https://goreportcard.com/report/github.com/yourusername/csvtojson)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

CSVtoJSON is a production‑ready command‑line utility that converts CSV files to JSON with built‑in validation, configurable logging, and environment‑driven configuration.

## Features

- **Streaming conversion** – Handles large files without loading everything into memory.
- **Header validation** – Ensures required columns are present before conversion.
- **Configurable logging** – Log level and format are driven by environment variables.
- **Zero external runtime dependencies** – Only standard Go libraries plus a few well‑maintained modules.
- **Tested** – Unit tests cover the core conversion logic.

## Installation