.PHONY: build
build:
	go build -o bin/blur ./cmd/blur/main.go

.PHONY: run
run:
	go run ./cmd/blur/main.go

.PHONY: test
test:
	go test ./...
