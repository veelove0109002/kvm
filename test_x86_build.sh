#!/usr/bin/env bash
#
# Test script to verify X86_64 build functionality
#
set -e

echo "🧪 Testing JetKVM X86_64 build..."

# Check if we're on X86_64
CURRENT_ARCH=$(uname -m)
if [ "$CURRENT_ARCH" != "x86_64" ]; then
    echo "❌ This test requires X86_64 architecture"
    echo "Current architecture: $CURRENT_ARCH"
    exit 1
fi

echo "✅ Architecture check passed: $CURRENT_ARCH"

# Check Go version
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed"
    exit 1
fi

GO_VERSION=$(go version | cut -d' ' -f3)
echo "✅ Go version: $GO_VERSION"

# Test build tags
echo "🔍 Testing build tags..."

# Test X86_64 build tag
echo "package main

import (
    \"fmt\"
    \"runtime\"
)

func main() {
    fmt.Printf(\"GOOS: %s, GOARCH: %s\\n\", runtime.GOOS, runtime.GOARCH)
}" > test_arch.go

GOOS=linux GOARCH=amd64 go build -o test_arch_x86 test_arch.go
./test_arch_x86
rm test_arch.go test_arch_x86

echo "✅ Build tags test passed"

# Test if we can compile the native package
echo "🔍 Testing native package compilation..."

cd internal/native
if GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v .; then
    echo "✅ Native package compiles successfully for X86_64"
else
    echo "❌ Native package compilation failed"
    exit 1
fi
cd ../..

# Test main package compilation
echo "🔍 Testing main package compilation..."

if TARGET_ARCH=x86_64 GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o test_jetkvm_x86 cmd/main.go; then
    echo "✅ Main package compiles successfully for X86_64"
    rm -f test_jetkvm_x86
else
    echo "❌ Main package compilation failed"
    exit 1
fi

echo ""
echo "🎉 All tests passed!"
echo "You can now build JetKVM for X86_64 using:"
echo "  ./build_x86.sh --dev"
echo ""