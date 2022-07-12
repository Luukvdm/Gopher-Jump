PKG_NAME = github.com/luukvdm/gopher-jump
BIN_NAME ?= gjump
BIN_DIR ?= dist

SRCS = $(shell git ls-files '*.go' | grep -v '^vendor/')

init:
	mkdir -p $(BIN_DIR)

run:
	go run $(PKG_NAME)

build: init
	GOOS=linux go build -o $(BIN_DIR)/$(BIN_NAME) $(PKG_NAME)

clean:
	go clean
	rm -rf $(BIN_DIR)

fmt: format
format:
	gofmt -s -l -w $(SRCS)
