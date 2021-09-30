GO_VERSION_SHORT:=$(shell echo `go version` | sed -E 's/.* go(.*) .*/\1/g')
ifneq ("1.17","$(shell printf "$(GO_VERSION_SHORT)\n1.17" | sort -V | head -1)")
$(error NEED GO VERSION >= 1.17. Found: $(GO_VERSION_SHORT))
endif

###############################################################################

.PHONY: build
build:
	go build -o bin/ ./...

.PHONY: test
test:
	go test -v -coverprofile cover.out ./...
	go tool cover -func cover.out | grep total | awk '{print ($$3)}'

.PHONY: bench
bench:
	go test -cpuprofile cpu.prof -memprofile mem.prof -bench ./...

.PHONY: lint
lint:
	@command -v golangci-lint 2>&1 > /dev/null || (echo "Install golangci-lint" && \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$(shell go env GOPATH)/bin" v1.42.1)
	golangci-lint run ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: style
style:
	find . -iname *.go | xargs gofmt -w

.PHONY: cover
cover:
	go tool cover -html cover.out

###############################################################################
