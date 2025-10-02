#!/bin/bash

# JetKVM Release Creation Script
# Usage: ./scripts/create_release.sh [version]
# Example: ./scripts/create_release.sh v1.0.0

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

# Check if version is provided
if [ $# -eq 0 ]; then
    print_error "Please provide a version number"
    echo "Usage: $0 <version>"
    echo "Example: $0 v1.0.0"
    exit 1
fi

VERSION=$1

# Validate version format
if [[ ! $VERSION =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    print_error "Invalid version format. Please use format: v1.0.0"
    exit 1
fi

print_info "Creating release for JetKVM $VERSION"

# Check if we're in the right directory
if [ ! -f "go.mod" ] || [ ! -d ".github/workflows" ]; then
    print_error "Please run this script from the project root directory"
    exit 1
fi

# Check if git is clean
if [ -n "$(git status --porcelain)" ]; then
    print_warning "Working directory is not clean. Please commit or stash changes first."
    git status --short
    exit 1
fi

# Check if tag already exists
if git tag -l | grep -q "^$VERSION$"; then
    print_error "Tag $VERSION already exists"
    exit 1
fi

# Update version in relevant files
print_info "Updating version in project files..."

# Update version in Makefile if VERSION variable exists
if grep -q "VERSION.*=" Makefile; then
    sed -i.bak "s/VERSION.*=.*/VERSION = $VERSION/" Makefile
    print_success "Updated version in Makefile"
fi

# Update version in package.json if it exists
if [ -f "ui/package.json" ]; then
    cd ui
    npm version ${VERSION#v} --no-git-tag-version
    cd ..
    print_success "Updated version in package.json"
fi

# Commit version changes
print_info "Committing version changes..."
git add -A
git commit -m "chore: bump version to $VERSION" || true

# Create and push tag
print_info "Creating git tag $VERSION..."
git tag -a "$VERSION" -m "Release $VERSION

🚀 JetKVM Multi-Architecture Release

This release includes:
- ARM Linux binary (hardware KVM)
- X86_64 Linux binary (software simulation)
- Complete multi-architecture support
- Improved build system and documentation

See release notes for detailed changes and download instructions."

print_success "Created tag $VERSION"

# Push changes and tag
print_info "Pushing changes and tag to remote..."
git push origin master
git push origin "$VERSION"

print_success "Successfully pushed tag $VERSION"

print_info "🎉 Release creation initiated!"
echo ""
print_info "Next steps:"
echo "1. GitHub Actions will automatically build binaries for both architectures"
echo "2. A release will be created at: https://github.com/$(git remote get-url origin | sed 's/.*github.com[:/]\([^/]*\/[^/]*\).*/\1/' | sed 's/\.git$//')/releases"
echo "3. Binaries will be uploaded automatically:"
echo "   - jetkvm-arm-linux (for ARM devices)"
echo "   - jetkvm-x86_64-linux (for X86_64 devices)"
echo "   - SHA256 checksums for verification"
echo ""
print_warning "Monitor the GitHub Actions workflow to ensure successful build and release"
echo "Workflow URL: https://github.com/$(git remote get-url origin | sed 's/.*github.com[:/]\([^/]*\/[^/]*\).*/\1/' | sed 's/\.git$//')/actions"