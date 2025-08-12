#!/bin/bash

set -e  # Exit on error

INSTALL_DIR="$HOME/.betterminal"
BINARY_URL="https://github.com/Mrton0121/betTerminal/releases/latest/download/betterminal"
BINARY_PATH="$INSTALL_DIR/betterminal"
BASHRC="$HOME/.bashrc"

# Create directory if not exists
mkdir -p "$INSTALL_DIR"

# Download the binary
echo "Downloading betterminal..."
curl -L "$BINARY_URL" -o "$BINARY_PATH"

# Make executable
chmod +x "$BINARY_PATH"

# Append to .bashrc if not already added
if ! grep -q "BETTERMINAL_CONFIG" "$BASHRC"; then
    echo "" >> "$BASHRC"
    echo "# BetTerminal configuration" >> "$BASHRC"
    echo "export BETTERMINAL_CONFIG=\$HOME/.betterminal" >> "$BASHRC"
    echo 'bt() {' >> "$BASHRC"
    echo '    output=$($HOME/.betterminal/betterminal "$@")' >> "$BASHRC"
    echo '    if [ $? -eq 0 ]; then' >> "$BASHRC"
    echo '        eval "$output"' >> "$BASHRC"
    echo '    else' >> "$BASHRC"
    echo '        echo "$output"' >> "$BASHRC"
    echo '    fi' >> "$BASHRC"
    echo '}' >> "$BASHRC"
    echo "BetTerminal configuration added to $BASHRC."
else
    echo "BetTerminal configuration already exists in $BASHRC."
fi

# Reload .bashrc for the current shell session
echo "Reloading $BASHRC..."
source "$BASHRC"

echo "Installation complete. You can now use the 'bt' command."
