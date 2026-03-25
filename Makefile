.DEFAULT_GOAL := all

.PHONY: all test show-cover build build-all

all:
	@go build ./...

test:
	go test -v -coverprofile=coverage.out ./...

show-cover:
	dokimi report --input=coverage.out --output=coverage.json --reporter=istanbul --show

build:
	go build -o ./bin/dokimi github.com/farbodsalimi/dokimi

build-all:
	./scripts/build-all.sh github.com/farbodsalimi/dokimi
