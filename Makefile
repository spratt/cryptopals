test: build fmt lint
	go test -v ./...

build:
	go build ./...

lint:
	golangci-lint --exclude-use-default=false run

fmt:
	gofumpt -w .
