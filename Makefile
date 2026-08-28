GO := go
PACKAGES := ./...

.PHONY: all build check clean coverage fmt fmt-check lint release test vet version

all: check

build:
	$(GO) build $(PACKAGES)

test:
	$(GO) test -race -count=1 $(PACKAGES)

coverage:
	$(GO) test -race -coverprofile=coverage.out -covermode=atomic -count=1 $(PACKAGES)
	$(GO) tool cover -func=coverage.out

fmt:
	$(GO) fmt $(PACKAGES)

fmt-check:
	@test -z "$$($(GO)fmt -l $$(find . -name '*.go' -not -path './.git/*'))" || { echo "Go files need formatting; run 'make fmt'."; exit 1; }

vet:
	$(GO) vet $(PACKAGES)

lint:
	$(MAKE) fmt-check
	$(MAKE) vet

check: fmt-check vet test

clean:
	rm -f coverage.out
