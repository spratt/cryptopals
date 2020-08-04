test: build fmt lint
	go test ./...

build:
	go build ./...

lint:
	golangci-lint --exclude-use-default=false run

fmt:
	gofumpt -w .
