test: build build_tests fmt lint
	go test -v ./...

build:
	go build ./...

build_tests:
	go test -run xxxxx ./...

lint:
	golangci-lint --exclude-use-default=false run

fmt:
	gofumpt -w .
