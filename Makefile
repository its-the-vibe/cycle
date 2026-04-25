BIN      := cycle
GO       := go
GOFLAGS  :=

.PHONY: all build test lint clean

all: build

build:
	$(GO) build $(GOFLAGS) -o $(BIN) .

test:
	$(GO) test $(GOFLAGS) ./...

lint:
	$(GO) vet ./...

clean:
	rm -f $(BIN)

.PHONY: install
install:
	$(GO) install $(GOFLAGS) ./...
