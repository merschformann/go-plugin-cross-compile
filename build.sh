#!/bin/bash

# Switch to script directory.
cd "$(dirname "$0")"

# Build the plugin
go build -buildmode plugin -trimpath -o plugin.so plugin/main.go

# Build the binary
go build -trimpath -o main entry/main.go
