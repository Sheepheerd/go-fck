#!/bin/bash

cd "$(dirname "$0")" || exit 1

OUT_DIR="./bin"
mkdir -p "$OUT_DIR"

echo "building go-fck"
go build -o "$OUT_DIR/main" ./src/cmd/main.go

if [ $? -eq 0 ]; then
    echo "go-fck  built. Running it..."
    "$OUT_DIR/main" $1
else
    echo "Build failed."
    exit 1
fi
