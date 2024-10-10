.DEFAULT_GOAL := help

################################################################################

BIN_DIR ?= bin/

################################################################################

# Note: use Makefile.local for customization
-include Makefile.local

## ▸▸▸ Development commands ◂◂◂

.PHONY: help
help:			## Display the help message
	@fgrep -h "## " $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/## //'

.PHONY: fmt
fmt:			## Format the codebase
	goimports -w -local github.com/arttet/Interview-Preparation-Kit-in-Go ${CURDIR}

.PHONY: lint
lint:			##  Run code linters to check for style and errors
	golangci-lint run ./...

.PHONY: build
build:			## Compile the packages
	go build -o ${BIN_DIR} ./...

.PHONY: test
test:			## Run unit tests
	go test -coverprofile coverage.out ./...
	go tool cover -func coverage.out | grep -E '100.0%|total' || echo "OK"
	go tool cover -func coverage.out | grep total | awk '{print ($$3)}'

.PHONY: bench
bench:			## Execute benchmarks
	go test -run Bench -bench=. ./... | tee bench_output.out
	awk '/Benchmark/{count ++; printf("%d,%s,%s,%s\n",count,$$1,$$2,$$3)}' bench_output.out > result.out

.PHONY: pprof
pprof:			## Run performance profiling with pprof
	cd platform/Coursera/Algorithms-Specialization/1-Divide-and-Conquer-Sorting-and-Searching-and-Randomized-Algorithms/1-Karatsuba-Multiplication/ && \
		go test -run=Bench -bench=. -cpuprofile cpu.out -memprofile mem.out -v && \
		go tool pprof -web cpu.out && \
		go tool pprof -web mem.out

.PHONY: coverage
coverage:		## Generate an HTML report for code coverage
	go tool cover -html coverage.out

.PHONY: validate
validate:		## Validate the CodeCov YAML
	cat codecov.yml | curl --data-binary @- https://codecov.io/validate

.PHONY: clean
clean:			## Remove generated build artifacts
	rm -rf ${BIN_DIR}
	rm *.out
