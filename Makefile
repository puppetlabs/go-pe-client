
all: build test lint format tidy

PHONY+= clean
clean:
	@echo "🔘 Cleaning build dir..."
	@rm -rf build
	@echo "🔘 Cleaning ts dist..."
	@rm -rf examples/ts-samples/dist
	@rm -rf examples/ts-samples/src/types
	@echo "🔘 Cleaning foobernetes state..."
	@rm -f deployment.json*

PHONY+= test
test:
	@echo "🔘 Running unit tests... (`date '+%H:%M:%S'`)"
	@go test $(TESTFLAGS) ./...

# Run go mod tidy and check go.sum is unchanged
PHONY+= tidy
tidy:
	@echo "🔘 Checking that go mod tidy does not make a change..."
	@cp go.sum go.sum.bak
	@go mod tidy
	@diff go.sum go.sum.bak && rm go.sum.bak || (echo "🔴 go mod tidy would make a change, exiting"; exit 1)
	@echo "✅ Checking go mod tidy complete"

# Format go code and error if any changes are made
PHONY+= format
format:
	@echo "🔘 Checking that go fmt does not make any changes..."
	@test -z $$(go fmt ./...) || (echo "🔴 go fmt would make a change, exiting"; exit 1)
	@echo "✅ Checking go fmt complete"

PHONY+= lint
lint: $(GOPATH)/bin/golangci-lint
	@echo "🔘 Linting $(1) (`date '+%H:%M:%S'`)"
	@lint=`golangci-lint run $(LINTFLAGS) $(1)`; \
	if [ "$$lint" != "" ]; \
	then echo "🔴 Lint found"; echo "$$lint"; exit 1;\
	else echo "✅ Lint-free (`date '+%H:%M:%S'`)"; \
	fi

PHONY+= build
build:
	@echo "🔘 Building - $(1) (`date '+%H:%M:%S'`)"
	@mkdir -p build/
	@go build ./...
	@echo "✅ Build complete - $(1) (`date '+%H:%M:%S'`)"

$(GOPATH)/bin/golangci-lint:
	@echo "🔘 Installing golangci-lint... (`date '+%H:%M:%S'`)"
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.24.0

.PHONY: $(PHONY)
