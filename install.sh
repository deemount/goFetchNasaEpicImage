#!/bin/bash

# Set the binary name
BINARY_NAME="nasa-epic-downloader"

# Set the output directory for the binaries
OUTPUT_DIR="bin"

# Set the target platforms
TARGETS=("linux/amd64" "linux/arm64" "darwin/amd64" "darwin/arm64")

# Function to copy the binary to the appropriate bin folder
copy_binary() {
    local target=$1
    local os=$(echo "$target" | cut -d'/' -f1)
    local binary_path="$OUTPUT_DIR/$BINARY_NAME-$target/$BINARY_NAME"

    if [ -f "$binary_path" ]; then
        case "$os" in
            "linux")
                sudo cp "$binary_path" "/usr/local/bin/$BINARY_NAME"
                sudo chmod +x "/usr/local/bin/$BINARY_NAME"
                echo "Binary copied to /usr/local/bin/$BINARY_NAME"
                ;;
            "darwin")
                cp "$binary_path" "/usr/local/bin/$BINARY_NAME"
                chmod +x "/usr/local/bin/$BINARY_NAME"
                echo "Binary copied to /usr/local/bin/$BINARY_NAME"
                ;;
            *)
                echo "Unsupported operating system: $os"
                ;;
        esac
    else
        echo "Binary not found: $binary_path"
    fi
}

# Function to write the environment variable to the appropriate shell configuration file
write_env_var() {
    local env_var="NASA_EPIC_API_KEY"
    local env_var_value
    read -r -p "Enter your NASA EPIC API key: " env_var_value
    local shell_config_file

    case "$SHELL" in
        */zsh)
            shell_config_file="$HOME/.zshrc"
            ;;
        */bash)
            shell_config_file="$HOME/.bashrc"
            ;;
        *)
            echo "Unsupported shell: $SHELL"
            return
            ;;
    esac

    if [ -f "$shell_config_file" ]; then
        if ! grep -qE "^export $env_var=" "$shell_config_file"; then
            echo "export $env_var=$env_var_value" >> "$shell_config_file"
            echo "Environment variable $env_var added to $shell_config_file"
        else
            echo "Environment variable $env_var already exists in $shell_config_file"
        fi
    else
        echo "File $shell_config_file not found"
    fi
}

# Build the binaries
make build

# Copy the binaries to the appropriate bin folder
for target in "${TARGETS[@]}"; do
    copy_binary "$target"
done

# Write the environment variable to the appropriate shell configuration file
write_env_var

echo "Done!"
