.PHONY: test lint fmt generate

test:
	GOCACHE=$(CURDIR)/.gocache go test ./...

lint:
	golangci-lint run ./...

fmt:
	gofmt -w cmd internal

generate:
	sqlc generate
