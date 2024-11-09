#!/bin/bash

# Set output binary name
OUTPUT="gomvc"

# Build the CLI tool
echo "Building the CLI tool..."
GOOS=$(go env GOOS)
GOARCH=$(go env GOARCH)

if [ "$GOOS" = "windows" ]; then
    OUTPUT+=".exe"
fi

go build -o $OUTPUT gomvc.go

# Determine install path based on OS
if [ "$GOOS" = "windows" ]; then
    INSTALL_PATH="/usr/local/bin"
elif [ "$GOOS" = "darwin" ]; then
    INSTALL_PATH="/usr/local/bin"
else
    INSTALL_PATH="/usr/local/bin"
fi

# Copy to the install path
echo "Copying binary to $INSTALL_PATH..."
sudo cp $OUTPUT $INSTALL_PATH

echo "Cleaning up..."
rm $OUTPUT

echo "Installation complete! You can now use gomvc from anywhere."
