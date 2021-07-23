LDFLAGS=-ldflags "-w -s "

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint: fmt
	golangci-lint run -v ./...

.PHONY: test
test:
	go test -race -cover -count=1 -short -v ./...
