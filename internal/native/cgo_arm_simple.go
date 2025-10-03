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

// UI Mock functions for ARM no-CGO builds
func uiInit(rotation uint16) {
	log.Printf("Mock: UI init for ARM (no CGO) with rotation: %d", rotation)
}

func uiTick() {
	// Mock UI tick - no actual UI processing needed
}

func uiSetVar(name string, value string) {
	log.Printf("Mock: UI set var %s = %s", name, value)
}

func uiGetVar(name string) string {
	log.Printf("Mock: UI get var %s", name)
	return "mock_value"
}

func uiSwitchToScreen(screen string) {
	log.Printf("Mock: UI switch to screen: %s", screen)
}

func uiGetCurrentScreen() string {
	return "main"
}

func uiObjAddState(objName string, state string) (bool, error) {
	log.Printf("Mock: UI add state %s to %s", state, objName)
	return true, nil
}

func uiObjClearState(objName string, state string) (bool, error) {
	log.Printf("Mock: UI clear state %s from %s", state, objName)
	return true, nil
}

func uiGetLVGLVersion() string {
	return "8.3.0-arm-mock"
}

func uiObjAddFlag(objName string, flag string) (bool, error) {
	log.Printf("Mock: UI add flag %s to %s", flag, objName)
	return true, nil
}

func uiObjClearFlag(objName string, flag string) (bool, error) {
	log.Printf("Mock: UI clear flag %s from %s", flag, objName)
	return true, nil
}

func uiObjHide(objName string) (bool, error) {
	log.Printf("Mock: UI hide object: %s", objName)
	return uiObjAddFlag(objName, "LV_OBJ_FLAG_HIDDEN")
}

func uiObjShow(objName string) (bool, error) {
	log.Printf("Mock: UI show object: %s", objName)
	return uiObjClearFlag(objName, "LV_OBJ_FLAG_HIDDEN")
}

func uiObjSetOpacity(objName string, opacity int) (bool, error) {
	log.Printf("Mock: UI set opacity %d for %s", opacity, objName)
	return true, nil
}

func uiObjFadeIn(objName string, duration uint32) (bool, error) {
	log.Printf("Mock: UI fade in %s over %d ms", objName, duration)
	return true, nil
}

func uiObjFadeOut(objName string, duration uint32) (bool, error) {
	log.Printf("Mock: UI fade out %s over %d ms", objName, duration)
	return true, nil
}

func uiLabelSetText(objName string, text string) (bool, error) {
	log.Printf("Mock: UI set text '%s' for %s", text, objName)
	return true, nil
}

func uiImgSetSrc(objName string, src string) (bool, error) {
	log.Printf("Mock: UI set image source '%s' for %s", src, objName)
	return true, nil
}

func uiDispSetRotation(rotation uint16) (bool, error) {
	log.Printf("Mock: UI set display rotation: %d", rotation)
	return true, nil
}

func uiEventCodeToName(code int) string {
	return fmt.Sprintf("MOCK_ARM_EVENT_%d", code)
}

// Video mock functions for ARM no-CGO builds
func videoInit() error {
	log.Println("Mock: Video init for ARM (no CGO)")
	return nil
}

func videoClose() error {
	log.Println("Mock: Video close for ARM")
	return nil
}

func videoSetStreamQualityFactor(factor float64) error {
	log.Printf("Mock: Setting video quality factor to %f", factor)
	return nil
}

func videoGetStreamQualityFactor() (float64, error) {
	return 0.8, nil
}

func videoGetState() VideoState {
	return VideoState{
		Ready:          true,
		Width:          1920,
		Height:         1080,
		FramePerSecond: 30.0,
	}
}

// Additional functions needed by native.go and video.go
func setUpNativeHandlers() {
	log.Println("Mock: Setting up native handlers for ARM (no CGO)")
}

func crash() {
	panic("Mock crash for ARM (no CGO)")
}

func videoSetEDID(edid string) error {
	log.Printf("Mock: Setting video EDID: %s", edid)
	return nil
}

func videoGetEDID() (string, error) {
	// Return a mock EDID for 1920x1080 display
	return "00ffffffffffff0010ac72404c384145" +
		"2e120103802f1e78eaee95a3544c9926" +
		"0f5054a54b00b300d100714fa9408180" +
		"8140010101011d007251d01e206e2855" +
		"00d9281100001e8c0ad08a20e02d1010" +
		"3e9600138e2100001e023a8018713827" +
		"40582c4500d9281100001e011d80d072" +
		"1c1620102c2580d9281100009e000000", nil
}

func videoLogStatus() string {
	return "Mock video status: ARM no-CGO implementation running"
}

func videoStop() {
	log.Println("Mock: Video stop for ARM (no CGO)")
}

func videoStart() {
	log.Println("Mock: Video start for ARM (no CGO)")
	// Simulate video state change
	go func() {
		videoState := VideoState{
			Ready:          true,
			Error:          "",
			Width:          1920,
			Height:         1080,
			FramePerSecond: 30.0,
		}
		select {
		case videoStateChan <- videoState:
		default:
		}
	}()
}

func videoShutdown() {
	log.Println("Mock: Video shutdown for ARM (no CGO)")
}

// Native struct methods for display.go compatibility

// UI initialization and tick methods
func (n *Native) setUIVars() {
	appVersionStr := "unknown"
	systemVersionStr := "unknown"
	
	if n.appVersion != nil {
		appVersionStr = n.appVersion.String()
	}
	if n.systemVersion != nil {
		systemVersionStr = n.systemVersion.String()
	}
	
	log.Printf("Mock ARM: Setting UI vars - app_version: %s, system_version: %s", 
		appVersionStr, systemVersionStr)
}

func (n *Native) initUI() {
	log.Printf("Mock ARM: Initializing UI with rotation: %d", n.displayRotation)
	n.setUIVars()
}

func (n *Native) tickUI() {
	log.Println("Mock ARM: UI tick loop started (no-CGO)")
	for {
		time.Sleep(16 * time.Millisecond) // ~60 FPS
		// Mock UI tick - no actual UI operations
	}
}

func (n *Native) SwitchToScreenIfDifferent(screenName string) {
	log.Printf("Mock: Switch to screen %s for ARM (no CGO)", screenName)
}

func (n *Native) UpdateLabelIfChanged(labelName, text string) {
	log.Printf("Mock: Update label %s to '%s' for ARM (no CGO)", labelName, text)
}

func (n *Native) UpdateLabelAndChangeVisibility(labelName, text string, visible bool) {
	log.Printf("Mock: Update label %s to '%s', visible=%t for ARM (no CGO)", labelName, text, visible)
}

func (n *Native) UIObjHide(objName string) (bool, error) {
	log.Printf("Mock: Hide UI object %s for ARM (no CGO)", objName)
	return true, nil
}

func (n *Native) UIObjShow(objName string) (bool, error) {
	log.Printf("Mock: Show UI object %s for ARM (no CGO)", objName)
	return true, nil
}

func (n *Native) UIObjAddState(objName string, state string) (bool, error) {
	log.Printf("Mock: Add state %s to UI object %s for ARM (no CGO)", state, objName)
	return true, nil
}

func (n *Native) UIObjClearState(objName string, state string) (bool, error) {
	log.Printf("Mock: Clear state %s from UI object %s for ARM (no CGO)", state, objName)
	return true, nil
}

func (n *Native) UISetVar(varName string, value interface{}) {
	log.Printf("Mock: Set UI variable %s to %v for ARM (no CGO)", varName, value)
}

func (n *Native) SwitchToScreenIf(screenName string, shouldSwitch []string) {
	log.Printf("Mock: Switch to screen %s with conditions %v for ARM (no CGO)", screenName, shouldSwitch)
}

func (n *Native) UIObjSetImageSrc(objName, imageSrc string) (bool, error) {
	log.Printf("Mock: Set image source of %s to %s for ARM (no CGO)", objName, imageSrc)
	return true, nil
}

func (n *Native) UIObjFadeOut(objName string, duration uint32) (bool, error) {
	log.Printf("Mock: Fade out %s over %dms for ARM (no CGO)", objName, duration)
	return true, nil
}

func (n *Native) UIObjFadeIn(objName string, duration uint32) (bool, error) {
	log.Printf("Mock: Fade in %s over %dms for ARM (no CGO)", objName, duration)
	return true, nil
}

func (n *Native) DisplaySetRotation(rotation uint16) (bool, error) {
	log.Printf("Mock: Set display rotation to %d degrees for ARM (no CGO)", rotation)
	return true, nil
}