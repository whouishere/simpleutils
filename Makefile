COREUTILS = cat cp dirname false mkdir printenv pwd rmdir touch true whoami

VERSION = 0.0.1
LDFLAGS = -X 'codeberg.org/whou/simpleutils/coreutils.Version=$(VERSION)'

BUILD_DIR = build
INSTALL_DIR =
BIN_PREFIX =
BIN_SUFIX =

ifeq ($(PREFIX),)
	INSTALL_DIR := ~/.local/bin
endif

ifeq ($(BIN_PREFIX),)
	BIN_PREFIX := su-
endif

ifeq ($(OS),Windows_NT)
	BIN_SUFIX += .exe
endif

BINARIES = $(patsubst %,$(BUILD_DIR)/$(BIN_PREFIX)%$(BIN_SUFIX),$(COREUTILS))

.PHONY: build run test install

build:
	@mkdir -p $(BUILD_DIR)
	@$(foreach util,$(COREUTILS),echo "Compiling $(util)..."; go build -o $(BUILD_DIR)/$(BIN_PREFIX)$(util)$(BIN_SUFIX) -ldflags="$(LDFLAGS)" ./coreutils/$(util)/$(util).go;)

run:
	go run coreutils/$(UTIL)/$(UTIL).go $(ARGS)

test:
	go test ./coreutils/...
	go test ./internal/...

install:
	@$(foreach util,$(BINARIES),install -v -D -t $(INSTALL_DIR) $(util);)
