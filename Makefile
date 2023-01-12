
all: build test format lint sec tidy

GOPATH := $(shell go env GOPATH)

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
	@go vet ./...
	@golangci-lint run \
		-E asciicheck \
		-E bodyclose \
		-E exhaustive \
		-E exportloopref \
		-E gci \
		-E gofmt \
		-E goimports \
		-E goimports \
		-E gosec \
		-E noctx \
		-E nolintlint \
		-E rowserrcheck \
		-E scopelint \
		-E sqlclosecheck \
		-E stylecheck \
		-E unconvert \
		-E unparam
	@echo "✅ Lint-free (`date '+%H:%M:%S'`)"

PHONY+= sec
sec: $(GOPATH)/bin/gosec
	@echo "🔘 Checking for security problems ... (`date '+%H:%M:%S'`)"
	@gosec -quiet ./...
	@echo "✅ No problems found (`date '+%H:%M:%S'`)"; \

PHONY+= build
build:
	@echo "🔘 Building - $(1) (`date '+%H:%M:%S'`)"
	@mkdir -p build/
	@go build ./...
	@echo "✅ Build complete - $(1) (`date '+%H:%M:%S'`)"

$(GOPATH)/bin/golangci-lint:
	@echo "🔘 Installing golangci-lint... (`date '+%H:%M:%S'`)"
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin

$(GOPATH)/bin/gosec:
	@echo "🔘 Installing gosec ... (`date '+%H:%M:%S'`)"
	@curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $(GOPATH)/bin

PHONY+= update-tools
update-tools: delete-tools $(GOPATH)/bin/golangci-lint $(GOPATH)/bin/gosec

PHONY+= delete-tools
delete-tools:
	@rm $(GOPATH)/bin/golangci-lint
	@rm $(GOPATH)/bin/gosec

.PHONY: $(PHONY)
