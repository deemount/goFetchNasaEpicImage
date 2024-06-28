# Set the name of the binary
BINARY_NAME=nasa-epic-downloader

# Set the list of target platforms (GOOS/GOARCH combinations)
TARGETS=\
  linux/amd64 \
  linux/arm64 \
  darwin/amd64 \
  darwin/arm64 \
  windows/amd64

# Set the output directory for the binaries
OUTPUT_DIR=bin

# Set the Go environment variables
export CGO_ENABLED=0

# Define the build target
.PHONY: build
build: $(addprefix $(OUTPUT_DIR)/$(BINARY_NAME)-, $(TARGETS))

# Define the target for each platform
$(OUTPUT_DIR)/$(BINARY_NAME)-%: $(shell find . -type f -name '*.go')
	@echo "Building for $*"
	@mkdir -p "$@"
	@GOOS=$(word 1, $(subst /, ,$*)) GOARCH=$(word 2, $(subst /, ,$*)) go build -o "$@/$(BINARY_NAME)" ./cmd/main.go

# Define a clean target to remove the output directory
.PHONY: clean
clean:
	@rm -rf "$(OUTPUT_DIR)"
