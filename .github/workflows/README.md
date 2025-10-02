# GitHub Actions Workflows

This directory contains GitHub Actions workflows for building, testing, and releasing JetKVM with multi-architecture support.

## Workflows Overview

### 🚀 `build.yml` - Complete Build Pipeline
**Triggers**: Push to main/develop, tags, PRs
- Builds frontend (React/TypeScript)
- Builds X86_64 binaries (native)
- Builds ARM binaries (Docker cross-compilation)
- Runs tests
- Creates releases for tags
- Builds and pushes Docker images

### ⚡ `quick-build.yml` - Fast CI/CD
**Triggers**: Push to any branch, PRs
- Quick frontend build
- X86_64 build test
- Code linting and formatting
- Go vet and mod tidy checks

### 📦 `release.yml` - Production Releases
**Triggers**: Git tags (v*), manual dispatch
- Creates GitHub releases
- Builds production binaries for both architectures
- Generates checksums
- Builds and pushes Docker images
- Publishes release automatically

## Usage Examples

### Automatic Builds
```bash
# Trigger quick build
git push origin feature/my-feature

# Trigger full build
git push origin main

# Create release
git tag v1.0.0
git push origin v1.0.0
```

### Manual Release
1. Go to Actions tab in GitHub
2. Select "Release" workflow
3. Click "Run workflow"
4. Enter version (e.g., v1.0.1)
5. Click "Run workflow"

## Build Artifacts

### X86_64 Binaries
- `jetkvm-{version}-linux-amd64` - Release binary
- `jetkvm-{version}-linux-amd64.sha256` - Checksum

### ARM Binaries
- `jetkvm-{version}-linux-arm` - Release binary for JetKVM hardware
- `jetkvm-{version}-linux-arm.sha256` - Checksum

### Docker Images
- `ghcr.io/{repo}-x86_64:latest` - Latest X86_64 image
- `ghcr.io/{repo}-x86_64:{version}` - Tagged version

## Environment Variables

The workflows use these environment variables:
- `GO_VERSION`: Go version for builds (1.24.4)
- `NODE_VERSION`: Node.js version for frontend (22.15.0)

## Secrets Required

- `GITHUB_TOKEN`: Automatically provided by GitHub
- No additional secrets needed for basic functionality

## Architecture Support

### X86_64 (Software-based)
- **Platform**: `linux/amd64`
- **CGO**: Disabled
- **Features**: Mock implementations for hardware features
- **Use Case**: Development, testing, software KVM

### ARM (Hardware-based)
- **Platform**: `linux/arm/v7`
- **CGO**: Enabled with cross-compilation
- **Features**: Full hardware support
- **Use Case**: Production JetKVM devices

## Workflow Status

Add these badges to your README.md:

```markdown
[![Build](https://github.com/{owner}/{repo}/actions/workflows/build.yml/badge.svg)](https://github.com/{owner}/{repo}/actions/workflows/build.yml)
[![Quick Build](https://github.com/{owner}/{repo}/actions/workflows/quick-build.yml/badge.svg)](https://github.com/{owner}/{repo}/actions/workflows/quick-build.yml)
[![Release](https://github.com/{owner}/{repo}/actions/workflows/release.yml/badge.svg)](https://github.com/{owner}/{repo}/actions/workflows/release.yml)
```

## Troubleshooting

### Common Issues

1. **Frontend build fails**
   - Check Node.js version compatibility
   - Verify package-lock.json is committed

2. **ARM build fails**
   - Docker buildx issues
   - Cross-compilation toolchain problems

3. **X86_64 build fails**
   - Go version compatibility
   - Missing build tags

### Debug Steps

1. Check workflow logs in Actions tab
2. Verify environment variables
3. Test builds locally:
   ```bash
   ./test_x86_build.sh
   ./build_x86.sh --dev
   ```

## Contributing

When adding new workflows:
1. Test locally first
2. Use appropriate triggers
3. Add proper error handling
4. Update this documentation