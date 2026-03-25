# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What is Dokimi

Dokimi is a Go CLI tool that provides helper commands for Go test coverage:
- **`report`** — Converts Go coverage profiles (`coverage.out`) to Istanbul JSON format, optionally opening an HTML report via the `istanbul` npm tool.
- **`check-coverage`** — Checks total coverage against a threshold, failing (or warning) if below.

## Build & Test Commands

```bash
go build ./...                        # build all packages
go build -o ./bin/dokimi .            # build binary
go test ./...                         # run all tests
go test -v -coverprofile=coverage.out ./...  # tests with coverage
go vet ./...                          # static analysis
```

Makefile targets:
```bash
make             # go build ./... (silent)
make test        # go test -v -coverprofile=coverage.out ./...
make build       # build binary to ./bin/dokimi
make build-all   # cross-compile for linux/darwin/windows amd64
```

## Architecture

```
main.go                          # entry point, sets Version via ldflags
cmd/
  root.go                        # cobra root command, creates ~/.dokimi dirs on init
  report.go                      # "report" subcommand — delegates to reporter implementations
  check_coverage.go              # "check-coverage" subcommand — shells out to `go tool cover`
internal/
  configs/paths.go               # global path constants (~/.dokimi/istanbul_tmp/*)
  reporters/istanbul/
    istanbul.go                   # Istanbul struct, types, functional options constructor
    write.go                      # WriteReport: parses Go cover profiles → Istanbul JSON
    show.go                       # ShowReport: writes JSON then runs `istanbul report` + `open`
```

Key patterns:
- Reporter uses **functional options** (`New(opts ...func(*Istanbul))`) with an injectable `writeFileFn` for testability.
- Coverage parsing uses `golang.org/x/tools/cover.ParseProfiles`.
- `ShowReport` depends on the `istanbul` npm package and macOS `open` command.
- Version is injected at build time via `-ldflags="-X 'main.Version=vX.Y.Z'"`.
