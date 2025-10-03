//go:build !cgo && !arm

package native

import (
	"log"
	"time"
)

func (n *Native) setUIVars() {
	appVersionStr := "unknown"
	systemVersionStr := "unknown"
	
	if n.appVersion != nil {
		appVersionStr = n.appVersion.String()
	}
	if n.systemVersion != nil {
		systemVersionStr = n.systemVersion.String()
	}
	
	log.Printf("Mock: Setting UI vars - app_version: %s, system_version: %s", 
		appVersionStr, systemVersionStr)
}

func (n *Native) initUI() {
	log.Printf("Mock: Initializing UI with rotation: %d", n.displayRotation)
	n.setUIVars()
}

func (n *Native) tickUI() {
	log.Println("Mock: UI tick loop started (no-CGO)")
	for {
		time.Sleep(16 * time.Millisecond) // ~60 FPS
		// Mock UI tick - no actual UI operations
	}
}

// GetLVGLVersion returns the LVGL version (mock for no-CGO)
func (n *Native) GetLVGLVersion() (string, error) {
	return "mock-lvgl-8.3.0", nil
}

// UIObjHide hides the object (mock for no-CGO)
func (n *Native) UIObjHide(objName string) (bool, error) {
	log.Printf("Mock: Hiding UI object: %s", objName)
	return true, nil
}

// UIObjShow shows the object (mock for no-CGO)
func (n *Native) UIObjShow(objName string) (bool, error) {
	log.Printf("Mock: Showing UI object: %s", objName)
	return true, nil
}

// UISetVar sets the variable (mock for no-CGO)
func (n *Native) UISetVar(name, value string) {
	log.Printf("Mock: Setting UI variable %s = %s", name, value)
}

// UIGetVar gets the variable (mock for no-CGO)
func (n *Native) UIGetVar(name string) string {
	log.Printf("Mock: Getting UI variable: %s", name)
	return "mock_value"
}

// UIObjAddState adds state to UI object (mock for no-CGO)
func (n *Native) UIObjAddState(objName string, state string) (bool, error) {
	log.Printf("Mock: Adding state %s to object %s", state, objName)
	return true, nil
}

// UIObjClearState clears state from UI object (mock for no-CGO)
func (n *Native) UIObjClearState(objName string, state string) (bool, error) {
	log.Printf("Mock: Clearing state %s from object %s", state, objName)
	return true, nil
}

// UIObjAddFlag adds flag to UI object (mock for no-CGO)
func (n *Native) UIObjAddFlag(objName string, flag string) (bool, error) {
	log.Printf("Mock: Adding flag %s to object %s", flag, objName)
	return true, nil
}

// UIObjClearFlag clears flag from UI object (mock for no-CGO)
func (n *Native) UIObjClearFlag(objName string, flag string) (bool, error) {
	log.Printf("Mock: Clearing flag %s from object %s", flag, objName)
	return true, nil
}

// UIObjSetOpacity sets the opacity of the object (mock for no-CGO)
func (n *Native) UIObjSetOpacity(objName string, opacity int) (bool, error) {
	log.Printf("Mock: Setting opacity %d for object %s", opacity, objName)
	return true, nil
}

// UIObjFadeIn fades in the object (mock for no-CGO)
func (n *Native) UIObjFadeIn(objName string, duration uint32) (bool, error) {
	log.Printf("Mock: Fading in object %s for %d ms", objName, duration)
	return true, nil
}

// UIObjFadeOut fades out the object (mock for no-CGO)
func (n *Native) UIObjFadeOut(objName string, duration uint32) (bool, error) {
	log.Printf("Mock: Fading out object %s for %d ms", objName, duration)
	return true, nil
}

// UIObjSetLabelText sets the text of the object (mock for no-CGO)
func (n *Native) UIObjSetLabelText(objName string, text string) (bool, error) {
	log.Printf("Mock: Setting text '%s' for object %s", text, objName)
	return true, nil
}

// UIObjSetImageSrc sets the image of the object (mock for no-CGO)
func (n *Native) UIObjSetImageSrc(objName string, image string) (bool, error) {
	log.Printf("Mock: Setting image '%s' for object %s", image, objName)
	return true, nil
}

// DisplaySetRotation sets the rotation of the display (mock for no-CGO)
func (n *Native) DisplaySetRotation(rotation uint16) (bool, error) {
	log.Printf("Mock: Setting display rotation to %d", rotation)
	return true, nil
}

// UpdateLabelIfChanged updates the label if the text has changed (mock for no-CGO)
func (n *Native) UpdateLabelIfChanged(objName string, newText string) {
	log.Printf("Mock: UpdateLabelIfChanged(%s, %s)", objName, newText)
}

// UpdateLabelAndChangeVisibility updates the label and changes the visibility (mock for no-CGO)
func (n *Native) UpdateLabelAndChangeVisibility(objName string, newText string, visible bool) {
	log.Printf("Mock: UpdateLabelAndChangeVisibility(%s, %s, %v)", objName, newText, visible)
}

// SwitchToScreenIf switches to the screen if conditions are met (mock for no-CGO)
func (n *Native) SwitchToScreenIf(screenName string, shouldSwitch []string) {
	log.Printf("Mock: SwitchToScreenIf(%s, %v)", screenName, shouldSwitch)
}

// SwitchToScreenIfDifferent switches to the screen if different (mock for no-CGO)
func (n *Native) SwitchToScreenIfDifferent(screenName string) {
	log.Printf("Mock: SwitchToScreenIfDifferent(%s)", screenName)
}

// Additional functions required by native.go

// setUpNativeHandlers sets up native handlers (mock for no-CGO)
func setUpNativeHandlers() {
	log.Println("Mock: Setting up native handlers (no-CGO)")
}

// videoInit initializes video system (mock for no-CGO)
func videoInit() error {
	log.Println("Mock: Video init (no-CGO)")
	return nil
}

// crash function for testing (mock for no-CGO)
func crash() {
	log.Println("Mock: Crash function called (no-CGO)")
	panic("Mock crash for testing")
}

// uiGetLVGLVersion returns LVGL version (mock for no-CGO)
func uiGetLVGLVersion() string {
	return "mock-lvgl-8.3.0-nocgo"
}

// videoSetStreamQualityFactor sets video quality factor (mock for no-CGO)
func videoSetStreamQualityFactor(factor float64) error {
	log.Printf("Mock: Setting video quality factor to %f (no-CGO)", factor)
	return nil
}

// videoGetStreamQualityFactor gets video quality factor (mock for no-CGO)
func videoGetStreamQualityFactor() (float64, error) {
	log.Println("Mock: Getting video quality factor (no-CGO)")
	return 1.0, nil
}

// videoSetEDID sets EDID (mock for no-CGO)
func videoSetEDID(edid string) error {
	log.Printf("Mock: Setting EDID: %s (no-CGO)", edid)
	return nil
}

// videoGetEDID gets EDID (mock for no-CGO)
func videoGetEDID() (string, error) {
	log.Println("Mock: Getting EDID (no-CGO)")
	return "mock-edid-data", nil
}

// videoLogStatus returns video status log (mock for no-CGO)
func videoLogStatus() string {
	log.Println("Mock: Getting video log status (no-CGO)")
	return "Mock video status: Running on X86_64 (no-CGO)"
}

// videoStop stops video capture (mock for no-CGO)
func videoStop() {
	log.Println("Mock: Video stop (no-CGO)")
}

// videoStart starts video capture (mock for no-CGO)
func videoStart() {
	log.Println("Mock: Video start (no-CGO)")
}

// videoShutdown shuts down video system (mock for no-CGO)
func videoShutdown() {
	log.Println("Mock: Video shutdown (no-CGO)")
}

// HDMI Display methods for X86_64 no-CGO builds

// EnableHDMIDisplay enables HDMI output (mock for no-CGO)
func (n *Native) EnableHDMIDisplay(webServerURL string) error {
	log.Printf("Mock: EnableHDMIDisplay(%s) (no-CGO)", webServerURL)
	return nil
}

// DisableHDMIDisplay disables HDMI output (mock for no-CGO)
func (n *Native) DisableHDMIDisplay() error {
	log.Println("Mock: DisableHDMIDisplay() (no-CGO)")
	return nil
}

// GetHDMIDisplayStatus returns HDMI display status (mock for no-CGO)
func (n *Native) GetHDMIDisplayStatus() map[string]interface{} {
	log.Println("Mock: GetHDMIDisplayStatus() (no-CGO)")
	return map[string]interface{}{
		"enabled":    false,
		"autoStart":  false,
		"webURL":     "",
		"xProcess":   nil,
		"browser":    nil,
		"mock":       true,
		"platform":   "no-cgo",
	}
}