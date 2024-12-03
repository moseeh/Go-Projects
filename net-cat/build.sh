#!/bin/bash

BINARY_NAME="TCPChat"

# Build the binary
echo "Building the binary..."
go build -o "$BINARY_NAME" ./cmd

# Check if the build was successful
if [ $? -eq 0 ]; then
  echo "Build successful! Binary is located at ."
else
  echo "Build failed!"
  exit 1
fi
