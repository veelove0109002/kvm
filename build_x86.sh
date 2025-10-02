#!/usr/bin/env bash
#
# Build script for X86_64 architecture
# This script builds JetKVM for X86_64 systems without ARM-specific dependencies
#
set -e

# Function to display help message
show_help() {
    echo "Usage: $0 [options]"
    echo
    echo "Options:"
    echo "  --skip-ui-build        Skip frontend/UI build"
    echo "  --dev                  Build development version"
    echo "  --release              Build release version"
    echo "  --help                 Display this help message"
    echo
    echo "Example:"
    echo "  $0 --dev"
    echo "  $0 --release"
}

# Default values
SCRIPT_PATH=$(realpath "$(dirname $(realpath "${BASH_SOURCE[0]}"))")
SKIP_UI_BUILD=false
BUILD_TYPE="dev"

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --skip-ui-build)
            SKIP_UI_BUILD=true
            shift
            ;;
        --dev)
            BUILD_TYPE="dev"
            shift
            ;;
        --release)
            BUILD_TYPE="release"
            shift
            ;;
        --help)
            show_help
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            show_help
            exit 1
            ;;
    esac
done

source ${SCRIPT_PATH}/scripts/build_utils.sh

# Check if we're on X86_64
CURRENT_ARCH=$(uname -m)
if [ "$CURRENT_ARCH" != "x86_64" ]; then
    msg_err "Error: This script is only for X86_64 architecture"
    msg_err "Current architecture: $CURRENT_ARCH"
    exit 1
fi

msg_info "Building JetKVM for X86_64 architecture"

# Set target architecture
export TARGET_ARCH=x86_64

# Build frontend if not skipped
if [[ "$SKIP_UI_BUILD" = false ]]; then
    if [[ ! -f "static/index.html" ]]; then
        msg_info "▶ Building frontend"
        make frontend SKIP_UI_BUILD=0
    else
        msg_info "▶ Frontend already built, skipping"
    fi
else
    msg_info "▶ Skipping frontend build"
fi

# Build the application
if [ "$BUILD_TYPE" = "release" ]; then
    msg_info "▶ Building release binary for X86_64"
    make build_release TARGET_ARCH=x86_64 SKIP_NATIVE_IF_EXISTS=1 SKIP_UI_BUILD=1
    msg_ok "✅ Release binary built: bin/jetkvm_app"
else
    msg_info "▶ Building development binary for X86_64"
    make build_dev TARGET_ARCH=x86_64 SKIP_NATIVE_IF_EXISTS=1 SKIP_UI_BUILD=1
    msg_ok "✅ Development binary built: bin/jetkvm_app"
fi

msg_ok "Build completed successfully!"
msg_info "You can now run the application with: ./bin/jetkvm_app"
msg_info "Note: This X86_64 build uses mock implementations for hardware-specific features"