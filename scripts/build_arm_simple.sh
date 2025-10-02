#!/bin/bash

# Simple ARM build script for CI/CD environments
# This script builds ARM binaries without complex native dependencies

set -e

echo "🔧 Building ARM binary for CI/CD..."

# Set environment variables
export TARGET_ARCH=arm
export GOOS=linux
export GOARCH=arm
export GOARM=7
export CGO_ENABLED=0

# Create bin directory
mkdir -p bin

echo "📦 Building Go binary..."
go build \
    -tags netgo,timetzdata,nomsgpack \
    -trimpath \
    -ldflags="-s -w -X github.com/prometheus/common/version.Branch=master -X github.com/prometheus/common/version.BuildDate=$(date -u +%FT%T%z) -X github.com/prometheus/common/version.Revision=$(git rev-parse HEAD || echo 'unknown') -X github.com/jetkvm/kvm.builtTimestamp=$(date -u +%s) -X github.com/jetkvm/kvm.builtAppVersion=0.4.8" \
    -o bin/jetkvm_app \
    ./cmd

echo "✅ ARM binary built successfully!"
echo "📁 Binary location: bin/jetkvm_app"
echo "📊 Binary info:"
file bin/jetkvm_app || echo "file command not available"
ls -lh bin/jetkvm_app