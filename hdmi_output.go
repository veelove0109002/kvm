package kvm

import (
	"fmt"
	"runtime"
	"time"

	"github.com/jetkvm/kvm/internal/native"
)

// HDMI output configuration and management

// initHDMIOutputIfEnabled initializes HDMI output if enabled in config
func initHDMIOutputIfEnabled() {
	// Only available on X86_64 Linux
	if runtime.GOARCH != "amd64" || runtime.GOOS != "linux" {
		logger.Debug().Msg("HDMI output not available on this platform")
		return
	}

	// Wait for web server to be ready
	time.Sleep(5 * time.Second)

	// Check if HDMI output is enabled in config
	if config.HDMIOutputEnabled {
		logger.Info().Msg("HDMI output enabled in config, starting HDMI display")
		
		webServerURL := getWebServerURL()
		if err := enableHDMIOutput(webServerURL); err != nil {
			logger.Error().Err(err).Msg("Failed to enable HDMI output")
		} else {
			logger.Info().Str("url", webServerURL).Msg("HDMI output enabled successfully")
		}
	} else {
		logger.Debug().Msg("HDMI output disabled in config")
	}
}

// getWebServerURL constructs the web server URL for HDMI output
func getWebServerURL() string {
	host := config.Host
	port := config.Port

	// Use localhost for HDMI output
	if host == "0.0.0.0" || host == "::" {
		host = "localhost"
	}

	return fmt.Sprintf("http://%s:%d", host, port)
}

// enableHDMIOutput enables HDMI output with the given URL
func enableHDMIOutput(webServerURL string) error {
	if nativeInstance == nil {
		return fmt.Errorf("native instance not initialized")
	}

	// Call the native HDMI output function
	return nativeInstance.RpcEnableHDMIOutput(webServerURL)
}

// disableHDMIOutput disables HDMI output
func disableHDMIOutput() error {
	if nativeInstance == nil {
		return fmt.Errorf("native instance not initialized")
	}

	return nativeInstance.RpcDisableHDMIOutput()
}

// toggleHDMIOutput toggles HDMI output on/off
func toggleHDMIOutput() error {
	if nativeInstance == nil {
		return fmt.Errorf("native instance not initialized")
	}

	webServerURL := getWebServerURL()
	return nativeInstance.RpcToggleHDMIOutput(webServerURL)
}

// getHDMIOutputStatus returns the current HDMI output status
func getHDMIOutputStatus() map[string]interface{} {
	if nativeInstance == nil {
		return map[string]interface{}{
			"enabled":   false,
			"available": false,
			"error":     "native instance not initialized",
		}
	}

	return nativeInstance.RpcGetHDMIOutputStatus()
}

// RPC handlers for HDMI output control

// rpcEnableHDMIOutput RPC handler to enable HDMI output
func rpcEnableHDMIOutput() error {
	logger.Info().Msg("RPC: Enabling HDMI output")
	webServerURL := getWebServerURL()
	return enableHDMIOutput(webServerURL)
}

// rpcDisableHDMIOutput RPC handler to disable HDMI output
func rpcDisableHDMIOutput() error {
	logger.Info().Msg("RPC: Disabling HDMI output")
	return disableHDMIOutput()
}

// rpcToggleHDMIOutput RPC handler to toggle HDMI output
func rpcToggleHDMIOutput() error {
	logger.Info().Msg("RPC: Toggling HDMI output")
	return toggleHDMIOutput()
}

// rpcGetHDMIOutputStatus RPC handler to get HDMI output status
func rpcGetHDMIOutputStatus() map[string]interface{} {
	status := getHDMIOutputStatus()
	logger.Info().Interface("status", status).Msg("RPC: Getting HDMI output status")
	return status
}