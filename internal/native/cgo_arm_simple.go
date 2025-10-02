//go:build linux && arm && !cgo

package native

import (
	"fmt"
	"log"
	"time"
)

// Simple ARM implementation without CGO dependencies
// This is used when building with CGO_ENABLED=0 for CI/CD

func InitNative() error {
	log.Println("Initializing ARM simple implementation (no CGO)")
	return nil
}

func GetNativeStatus() map[string]interface{} {
	return map[string]interface{}{
		"implementation": "arm-simple",
		"cgo_enabled":    false,
		"timestamp":      time.Now().Unix(),
		"status":         "running",
	}
}

func SetDisplayBrightness(brightness int) error {
	log.Printf("Mock: Setting display brightness to %d%%", brightness)
	return nil
}

func GetDisplayBrightness() (int, error) {
	return 50, nil // Mock 50% brightness
}

func SetLEDState(state string) error {
	log.Printf("Mock: Setting LED state to %s", state)
	return nil
}

func GetSystemTemperature() (float64, error) {
	return 45.5, nil // Mock temperature
}

func TriggerSystemReset() error {
	log.Println("Mock: System reset triggered")
	return fmt.Errorf("system reset not available in simple ARM build")
}

func GetHardwareInfo() map[string]string {
	return map[string]string{
		"platform":      "ARM Linux",
		"implementation": "simple",
		"cgo":           "disabled",
		"note":          "Simplified ARM build for CI/CD",
	}
}

func UpdateFirmware(firmwarePath string) error {
	log.Printf("Mock: Firmware update with %s", firmwarePath)
	return fmt.Errorf("firmware update not available in simple ARM build")
}

func GetVideoCapabilities() map[string]interface{} {
	return map[string]interface{}{
		"max_resolution": "1920x1080",
		"formats":        []string{"h264", "mjpeg"},
		"mock":           true,
	}
}

func StartVideoCapture(config map[string]interface{}) error {
	log.Printf("Mock: Starting video capture with config: %+v", config)
	return nil
}

func StopVideoCapture() error {
	log.Println("Mock: Stopping video capture")
	return nil
}