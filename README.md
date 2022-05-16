# dokimi

Dokimi is a Go testing framework.

**Table of Contents**

- [dokimi](#dokimi)
  - [Installation](#installation)
  - [Commands](#commands)
    - [report](#report)
      - [Write Istanbul json reports and display them](#write-istanbul-json-reports-and-display-them)
      - [Only write Istanbul json reports](#only-write-istanbul-json-reports)

## Installation

```bash
go install github.com/farbodsalimi/dokimi@latest
```

## Commands

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
