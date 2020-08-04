test: build fmt lint
	go test ./...

build:
	go build ./...

lint:
	golangci-lint run ./... --enable-all

fmt:
	gofumpt -w .
