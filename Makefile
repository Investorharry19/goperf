.PHONY: fmt vet build test race check example loadgen

fmt:
	gofmt -w .

vet:
	go vet ./...

build:
	go build ./...

test:
	go test ./...

race:
	go test -race -count=1 ./...

check:
	@echo "Running gofmt..."
	@cmd /c "gofmt -l ."
	@if not errorlevel 1 (echo "gofmt reported issues" & exit /b 1)
	go vet ./...
	go build ./...
	go test -race -count=1 ./...

example:
	go run ./examples/basic

loadgen:
	go run ./examples/loadgen
