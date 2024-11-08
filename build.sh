#!/bin/bash

cd "$(dirname "$0")" || exit 1

OUT_DIR="./bin"
mkdir -p "$OUT_DIR"

cd src
go mod tidy

echo "Building bf..."
go build -o "../$OUT_DIR/bf" ./cmd/bf

if [ $? -ne 0 ]; then
    echo "Build for bf failed."
    exit 1
fi

echo "Building bfc..."
go build -o "../$OUT_DIR/bfc" ./cmd/bfc

if [ $? -ne 0 ]; then
    echo "Build for bfc failed."
    exit 1
fi
