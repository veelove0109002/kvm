# JetKVM X86_64 Support

This document describes the X86_64 architecture support for JetKVM.

## Overview

JetKVM has been modified to support X86_64 architecture alongside the original ARM architecture. The X86_64 version provides a software-based implementation that can run on standard PC hardware.

## Architecture Support

### Original ARM Version
- **Target**: ARM embedded devices (Rockchip RK3588)
- **Features**: Full hardware acceleration, HDMI capture, USB gadget support
- **Use Case**: Dedicated KVM hardware device

### New X86_64 Version
- **Target**: Standard X86_64 Linux systems
- **Features**: Software-based implementation with mock hardware interfaces
- **Use Case**: Development, testing, or software-based KVM solutions

## Building for X86_64

### Prerequisites
- Linux system with X86_64 architecture
- Go 1.24.4 or later
- Node.js 22.15.0 or later (for frontend)

### Quick Build
```bash
# Make the build script executable
chmod +x build_x86.sh

# Build development version
./build_x86.sh --dev

# Build release version
./build_x86.sh --release

# Skip frontend build (if already built)
./build_x86.sh --dev --skip-ui-build
```

### Manual Build
```bash
# Set target architecture
export TARGET_ARCH=x86_64

# Build frontend (optional)
make frontend

# Build development version
make build_dev TARGET_ARCH=x86_64

# Build release version
make build_release TARGET_ARCH=x86_64
```

## Running on X86_64

```bash
# Run the application
./bin/jetkvm_app

# The web interface will be available at:
# http://localhost:8080 (or your configured port)
```

## Differences from ARM Version

### Hardware Features (Mock Implementations)
- **Video Capture**: Mock video source (simulated 1920x1080@60fps)
- **HDMI Input**: Software simulation
- **USB Gadget**: Mock implementation
- **Hardware UI**: Console-based logging instead of physical display
- **RTC**: Uses system time instead of hardware RTC

### Supported Features
- ✅ Web interface
- ✅ WebRTC streaming (with mock video)
- ✅ Network configuration
- ✅ Authentication
- ✅ Configuration management
- ✅ Logging and monitoring
- ✅ API endpoints

### Limited/Mock Features
- 🔶 Video capture (mock implementation)
- 🔶 USB HID (mock implementation)
- 🔶 Hardware display (console logging)
- 🔶 EDID management (mock data)

### Not Supported
- ❌ Real HDMI capture
- ❌ Physical USB gadget functionality
- ❌ Hardware-specific features (Rockchip MPP, RGA)

## Configuration

The X86_64 version uses the same configuration format as the ARM version. However, some hardware-specific settings will be ignored or mocked.

### Example Configuration
```json
{
  "web": {
    "port": 8080,
    "host": "0.0.0.0"
  },
  "video": {
    "quality_factor": 1.0,
    "mock_enabled": true
  },
  "auth": {
    "password": "your_password_here"
  }
}
```

## Development

### Adding X86_64 Support to New Features

When adding new features, consider X86_64 compatibility:

1. **Use build tags** for platform-specific code:
   ```go
   //go:build linux && amd64
   // X86_64 implementation
   
   //go:build linux && arm
   // ARM implementation
   ```

2. **Provide mock implementations** for hardware features:
   ```go
   func hardwareFunction() error {
       if isX86_64() {
           return mockImplementation()
       }
       return realHardwareImplementation()
   }
   ```

3. **Test on both architectures** when possible.

### File Structure for Multi-Architecture Support

```
internal/
├── native/
│   ├── cgo_linux.go      # ARM with CGO
│   ├── cgo_x86_64.go     # X86_64 mock implementation
│   └── cgo_notlinux.go   # Other platforms
├── network/
│   ├── netif_linux.go    # Linux (both ARM and X86_64)
│   └── netif_notlinux.go # Other platforms
└── timesync/
    ├── rtc_linux.go      # Linux (both ARM and X86_64)
    └── rtc_notlinux.go   # Other platforms
```

## Troubleshooting

### Build Issues

1. **CGO errors**: Make sure `CGO_ENABLED=0` for X86_64 builds
2. **Missing dependencies**: Install Go and Node.js
3. **Permission errors**: Make sure build scripts are executable

### Runtime Issues

1. **Port conflicts**: Change the port in configuration
2. **Permission errors**: Run with appropriate privileges
3. **Mock warnings**: These are expected for hardware features

## Use Cases

### Development and Testing
- Develop and test JetKVM features without ARM hardware
- CI/CD pipelines on X86_64 infrastructure
- Local development environment

### Software-based KVM
- Remote desktop solutions
- Screen sharing applications
- Virtual KVM implementations

### Educational and Research
- Understanding KVM protocols
- WebRTC streaming experiments
- Network management tools

## Limitations

The X86_64 version is primarily intended for:
- Development and testing
- Software-based implementations
- Educational purposes

For production KVM-over-IP functionality, the ARM version with dedicated hardware is recommended.

## Contributing

When contributing to JetKVM with X86_64 support:

1. Test changes on both ARM and X86_64 when possible
2. Use appropriate build tags for platform-specific code
3. Provide mock implementations for hardware features
4. Update documentation for both architectures

## Future Improvements

Potential enhancements for X86_64 support:
- Integration with screen capture libraries
- USB passthrough support
- Hardware acceleration where available
- Better video source options