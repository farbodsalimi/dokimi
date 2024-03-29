# dokimi

[![Go](https://github.com/farbodsalimi/dokimi/actions/workflows/go.yml/badge.svg)](https://github.com/farbodsalimi/dokimi/actions/workflows/go.yml)
[![Release](https://github.com/farbodsalimi/dokimi/actions/workflows/release.yml/badge.svg)](https://github.com/farbodsalimi/dokimi/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/farbodsalimi/dokimi)](https://goreportcard.com/report/github.com/farbodsalimi/dokimi)
[![Maintainability](https://api.codeclimate.com/v1/badges/7c1b6f2aca67479f4220/maintainability)](https://codeclimate.com/github/farbodsalimi/dokimi/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/7c1b6f2aca67479f4220/test_coverage)](https://codeclimate.com/github/farbodsalimi/dokimi/test_coverage)

Dokimi provides some helper commands for testing in Go.

**Table of Contents**

- [dokimi](#dokimi)
  - [Installation](#installation)
  - [Commands](#commands)
    - [check-coverage](#check-coverage)
    - [report](#report)
      - [Write Istanbul json reports and display them](#write-istanbul-json-reports-and-display-them)
      - [Only write Istanbul json reports](#only-write-istanbul-json-reports)

## Installation

```bash
go install github.com/farbodsalimi/dokimi@latest
```

## Commands

### check-coverage

```
Usage:
  dokimi check-coverage [flags]

Flags:
  -c, --coverprofile string   coverprofile (default "coverage.out")
  -d, --do-not-fail           do-not-fail
  -h, --help                  help for check-coverage
  -t, --threshold float       threshold (default 100)
```

Example:

```bash
dokimi check-coverage --threshold=90 --do-not-fail
```

### report

```
Usage:
  dokimi report [flags]

Flags:
  -h, --help              help for report
  -i, --input string      Path to input file
  -o, --output string     Path to output file
  -r, --reporter string   Reporter name e.g. istanbul, lcov, ...
      --show              Shows written reports
```

#### Write Istanbul json reports and display them

1. Go inside your project directory and run your tests:

   ```bash
   go test -v -coverprofile=coverage.out ./...
   ```

2. Show your coverage in Istanbul UI:

   ```bash
   dokimi report --input=coverage.out --output=coverage.json --reporter=istanbul --show
   ```

#### Only write Istanbul json reports

1. Go inside your project directory and run your tests:

   ```bash
   go test -v -coverprofile=coverage.out ./...
   ```

2. Generate Istanbul json file:

   ```bash
   dokimi report --input=coverage.out --output=coverage.json --reporter=istanbul
   ```
