//go:build !cgo

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

func (n *Native) GetLVGLVersion() string {
	return "mock-lvgl-8.3.0"
}

func (n *Native) HideObj(name string) {
	log.Printf("Mock: Hiding UI object: %s", name)
}

func (n *Native) ShowObj(name string) {
	log.Printf("Mock: Showing UI object: %s", name)
}

func (n *Native) SetUIVar(name, value string) {
	log.Printf("Mock: Setting UI variable %s = %s", name, value)
}

func (n *Native) GetUIVar(name string) string {
	log.Printf("Mock: Getting UI variable: %s", name)
	return "mock_value"
}

func (n *Native) SetUIStatus(status string) {
	log.Printf("Mock: Setting UI status: %s", status)
}

func (n *Native) GetUIStatus() string {
	return "mock_status"
}

func (n *Native) SetUIProgress(progress int) {
	log.Printf("Mock: Setting UI progress: %d%%", progress)
}

func (n *Native) GetUIProgress() int {
	return 50 // Mock progress
}