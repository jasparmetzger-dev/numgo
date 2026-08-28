GO := go
FLAGS :=
BUILD_DIR := bin/
PACKAGES := ./... # matches all packages recursively

.PHONY: clean test

build:
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(GOFLAGS) -v $(PACKAGES)

test:
	$(GO) test -race -cover -coverprofile=coverage.out -count=1 $(PACKAGES)
	$(GO) vet $(PACKAGES)

lint:
	$(GO) fmt $(PACKAGES)
	echo "fmt done."
	$(GO) vet $(PACKAGES)
	echo "vet done."

cover: test
	$(GO) tool cover -func=coverage.out

clean:
	rm -rf $(BUILD_DIR) coverage.out
