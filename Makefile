.PHONY: build
build-blur:
	go build -o bin/blur ./cmd/blur/main.go

.PHONY run-blur:
	go run ./cmd/blur/main.go

.PHONY: test
test:
	go test ./...
