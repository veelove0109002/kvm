//go:build linux && amd64

package native

import (
	"fmt"
	"log"
	"time"
)

// X86_64 specific display functions with HDMI output support

func (n *Native) setUIVars() {
	appVersionStr := "unknown"
	systemVersionStr := "unknown"
	
	if n.appVersion != nil {
		appVersionStr = n.appVersion.String()
	}
	if n.systemVersion != nil {
		systemVersionStr = n.systemVersion.String()
	}
	
	log.Printf("Mock: Setting UI vars - App: %s, System: %s", appVersionStr, systemVersionStr)
	
	// If HDMI output is enabled, we could update the display here
	if IsHDMIOutputEnabled() {
		log.Printf("HDMI Output: UI variables updated")
	}
}

func (n *Native) initUI() {
	log.Printf("Mock: Initializing UI with rotation: %d", n.displayRotation)
	n.setUIVars()
	
	// Initialize HDMI output if requested
	// This will be called from the main application
}

func (n *Native) tickUI() {
	log.Println("Mock: UI tick loop started (no-CGO)")
	for {
		// Mock UI tick - minimal processing for X86_64
		time.Sleep(100 * time.Millisecond)
		
		// Check if HDMI output needs attention
		if IsHDMIOutputEnabled() {
			// HDMI output is handled by the browser, no additional work needed
		}
	}
}

// GetLVGLVersion returns the LVGL version (mock for X86_64)
func (n *Native) GetLVGLVersion() (string, error) {
	return "8.3.0-x86_64-mock", nil
}

// UIObjHide hides the object (mock for X86_64)
func (n *Native) UIObjHide(objName string) (bool, error) {
	log.Printf("Mock: Hiding UI object: %s", objName)
	return true, nil
}

// UIObjShow shows the object (mock for X86_64)
func (n *Native) UIObjShow(objName string) (bool, error) {
	log.Printf("Mock: Showing UI object: %s", objName)
	return true, nil
}

// UISetVar sets the variable (mock for X86_64)
func (n *Native) UISetVar(name string, value string) {
	log.Printf("Mock: Setting UI variable %s = %s", name, value)
}

// UIGetVar gets the variable (mock for X86_64)
func (n *Native) UIGetVar(name string) string {
	log.Printf("Mock: Getting UI variable: %s", name)
	return "mock_value"
}

// SwitchToScreenIfDifferent switches to a different screen if needed
func (n *Native) SwitchToScreenIfDifferent(screen string) {
	log.Printf("Mock: Switching to screen: %s", screen)
	
	if IsHDMIOutputEnabled() {
		log.Printf("HDMI Output: Screen switched to %s", screen)
	}
}

// UIObjAddState adds state to UI object
func (n *Native) UIObjAddState(objName string, state string) (bool, error) {
	log.Printf("Mock: Adding state %s to object %s", state, objName)
	return true, nil
}

// UIObjClearState clears state from UI object
func (n *Native) UIObjClearState(objName string, state string) (bool, error) {
	log.Printf("Mock: Clearing state %s from object %s", state, objName)
	return true, nil
}

// UIObjAddFlag adds flag to UI object
func (n *Native) UIObjAddFlag(objName string, flag string) (bool, error) {
	log.Printf("Mock: Adding flag %s to object %s", flag, objName)
	return true, nil
}

// UIObjClearFlag clears flag from UI object
func (n *Native) UIObjClearFlag(objName string, flag string) (bool, error) {
	log.Printf("Mock: Clearing flag %s from object %s", flag, objName)
	return true, nil
}

// UpdateLabelAndChangeVisibility updates label text and visibility
func (n *Native) UpdateLabelAndChangeVisibility(objName, text string, visible bool) {
	log.Printf("Mock: UpdateLabelAndChangeVisibility(%s, %s, %v)", objName, text, visible)
	
	if IsHDMIOutputEnabled() {
		log.Printf("HDMI Output: Label %s updated to '%s', visible: %v", objName, text, visible)
	}
}

// UpdateLabelIfChanged updates label if the text has changed
func (n *Native) UpdateLabelIfChanged(objName, text string) {
	log.Printf("Mock: UpdateLabelIfChanged(%s, %s)", objName, text)
	
	if IsHDMIOutputEnabled() {
		log.Printf("HDMI Output: Label %s updated to '%s'", objName, text)
	}
}

// SetDisplayBrightness sets the display brightness (mock for X86_64)
func SetDisplayBrightness(brightness int) error {
	log.Printf("Mock: Setting display brightness to %d%%", brightness)
	
	if IsHDMIOutputEnabled() {
		log.Printf("HDMI Output: Brightness setting ignored (controlled by monitor)")
	}
	
	return nil
}

// Additional HDMI-specific display functions

// EnableHDMIDisplay enables HDMI output for the JetKVM interface
func (n *Native) EnableHDMIDisplay(webServerURL string) error {
	log.Printf("Enabling HDMI display output for URL: %s", webServerURL)
	
	// Initialize HDMI output system
	if err := InitHDMIOutput(webServerURL); err != nil {
		return fmt.Errorf("failed to initialize HDMI output: %w", err)
	}
	
	// Enable the output
	if err := EnableHDMIOutput(); err != nil {
		return fmt.Errorf("failed to enable HDMI output: %w", err)
	}
	
	log.Println("HDMI display output enabled successfully")
	return nil
}

// DisableHDMIDisplay disables HDMI output
func (n *Native) DisableHDMIDisplay() error {
	log.Println("Disabling HDMI display output")
	
	if err := DisableHDMIOutput(); err != nil {
		return fmt.Errorf("failed to disable HDMI output: %w", err)
	}
	
	log.Println("HDMI display output disabled successfully")
	return nil
}

// GetHDMIDisplayStatus returns the current HDMI display status
func (n *Native) GetHDMIDisplayStatus() map[string]interface{} {
	return GetHDMIOutputStatus()
}