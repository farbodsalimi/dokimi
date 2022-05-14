# dokimi

Dokimi is a Go testing framework.

**Table of Contents**

- [dokimi](#dokimi)
  - [Installation](#installation)
  - [Commands](#commands)
    - [covgen](#covgen)
      - [Generate Istanbul coverage file](#generate-istanbul-coverage-file)

## Installation

```bash
go install github.com/farbodsalimi/dokimi@latest
```

## Commands

### covgen

Generates coverage files in different formats.

#### Generate Istanbul coverage file

1. Go inside your project directory and run your tests:

   ```bash
   go test -v -coverprofile=coverage.out ./...
   ```

2. Generate Istanbul json file:

   ```bash
   dokimi covgen --reporter=istanbul --in=coverage.out --out=coverage.json
   ```

3. Use Istanbul CLI to generate HTML:

   ```bash
   istanbul report --include coverage.json --dir istanbul html
   ```

4. Open HTML:

   ```bash
   open istanbul/index.html
   ```
