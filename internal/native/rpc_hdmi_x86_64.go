//go:build linux && amd64

package native

import (
	"fmt"
)

// RPC methods for HDMI output control on X86_64

// RpcEnableHDMIOutput enables HDMI output
func (n *Native) RpcEnableHDMIOutput(webServerURL string) error {
	nativeLogger.Info().Str("url", webServerURL).Msg("RPC: Enabling HDMI output")
	
	if webServerURL == "" {
		webServerURL = "http://localhost:8080"
	}
	
	return n.EnableHDMIDisplay(webServerURL)
}

// RpcDisableHDMIOutput disables HDMI output
func (n *Native) RpcDisableHDMIOutput() error {
	nativeLogger.Info().Msg("RPC: Disabling HDMI output")
	return n.DisableHDMIDisplay()
}

// RpcGetHDMIOutputStatus returns HDMI output status
func (n *Native) RpcGetHDMIOutputStatus() map[string]interface{} {
	status := n.GetHDMIDisplayStatus()
	nativeLogger.Info().Interface("status", status).Msg("RPC: Getting HDMI output status")
	return status
}

// RpcToggleHDMIOutput toggles HDMI output on/off
func (n *Native) RpcToggleHDMIOutput(webServerURL string) error {
	status := n.GetHDMIDisplayStatus()
	enabled, ok := status["enabled"].(bool)
	if !ok {
		return fmt.Errorf("failed to get HDMI output status")
	}
	
	if enabled {
		nativeLogger.Info().Msg("RPC: Toggling HDMI output OFF")
		return n.RpcDisableHDMIOutput()
	} else {
		nativeLogger.Info().Msg("RPC: Toggling HDMI output ON")
		return n.RpcEnableHDMIOutput(webServerURL)
	}
}