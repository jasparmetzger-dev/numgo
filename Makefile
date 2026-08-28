GO := go
FLAGS :=
BUILD_DIR := bin/
PACKAGES := ./... # matches all packages recursively

build:
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(GOFLAGS) -v $(PACKAGES)

test:
	$(GO) test -v -race -cover -count=1 $(PACKAGES)

lint:
	$(GO) fmt $(PACKAGES)
	echo "fmt done."
	$(GO) vet $(PACKAGES)
	echo "vet done."
clean:
	rm -rf $(BUILD_DIR) coverage.out
