#!/bin/bash

cd "$(dirname "$0")" || exit 1

OUT_DIR="./bin"
mkdir -p "$OUT_DIR"

cd src
go mod tidy

echo "Building go-fck..."
go build -o "../$OUT_DIR/main" ./cmd/main.go

if [ $? -eq 0 ]; then
    echo "Build successful! Running go-fck..."
    "../$OUT_DIR/main"
else
    echo "Build failed."
    exit 1
fi
