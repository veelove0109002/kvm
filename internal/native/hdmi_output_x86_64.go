//go:build linux && amd64

package native

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// HDMIOutput manages HDMI display output for X86_64 systems
type HDMIOutput struct {
	enabled        bool
	displayDevice  string
	resolution     string
	webServerURL   string
	browserProcess *exec.Cmd
	ctx            context.Context
	cancel         context.CancelFunc
}

var hdmiOutput *HDMIOutput

// InitHDMIOutput initializes HDMI output functionality
func InitHDMIOutput(webServerURL string) error {
	nativeLogger.Info().Msg("Initializing HDMI output for X86_64")
	
	ctx, cancel := context.WithCancel(context.Background())
	hdmiOutput = &HDMIOutput{
		enabled:       false,
		webServerURL:  webServerURL,
		ctx:           ctx,
		cancel:        cancel,
	}
	
	// Detect available display devices
	if err := hdmiOutput.detectDisplayDevices(); err != nil {
		nativeLogger.Warn().Err(err).Msg("Failed to detect display devices")
		return err
	}
	
	return nil
}

// detectDisplayDevices finds available HDMI/display outputs
func (h *HDMIOutput) detectDisplayDevices() error {
	// Check for DRM devices
	drmPath := "/sys/class/drm"
	entries, err := os.ReadDir(drmPath)
	if err != nil {
		return fmt.Errorf("failed to read DRM devices: %w", err)
	}
	
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), "card") && strings.Contains(entry.Name(), "-") {
			// Found a display connector
			h.displayDevice = entry.Name()
			nativeLogger.Info().Str("device", h.displayDevice).Msg("Found display device")
			break
		}
	}
	
	if h.displayDevice == "" {
		return fmt.Errorf("no display devices found")
	}
	
	// Try to detect resolution
	h.resolution = "1920x1080"
	nativeLogger.Info().Str("resolution", h.resolution).Msg("Using default resolution")
	
	return nil
}

// EnableHDMIOutput starts displaying the web interface on HDMI
func EnableHDMIOutput() error {
	if hdmiOutput == nil {
		return fmt.Errorf("HDMI output not initialized")
	}
	
	if hdmiOutput.enabled {
		nativeLogger.Info().Msg("HDMI output already enabled")
		return nil
	}
	
	nativeLogger.Info().Msg("Enabling HDMI output")
	
	// Install required packages if needed
	if err := hdmiOutput.ensureDisplayEnvironment(); err != nil {
		nativeLogger.Error().Err(err).Msg("Failed to setup display environment")
		return err
	}
	
	// Start X server and browser
	if err := hdmiOutput.startDisplayOutput(); err != nil {
		nativeLogger.Error().Err(err).Msg("Failed to start display output")
		return err
	}
	
	hdmiOutput.enabled = true
	nativeLogger.Info().Msg("HDMI output enabled successfully")
	
	return nil
}

// DisableHDMIOutput stops HDMI display output
func DisableHDMIOutput() error {
	if hdmiOutput == nil || !hdmiOutput.enabled {
		return nil
	}
	
	nativeLogger.Info().Msg("Disabling HDMI output")
	
	// Stop browser process
	if hdmiOutput.browserProcess != nil {
		hdmiOutput.browserProcess.Process.Kill()
		hdmiOutput.browserProcess.Wait()
		hdmiOutput.browserProcess = nil
	}
	
	// Stop X server
	exec.Command("pkill", "-f", "Xorg.*jetkvm").Run()
	
	hdmiOutput.enabled = false
	nativeLogger.Info().Msg("HDMI output disabled")
	
	return nil
}

// ensureDisplayEnvironment installs and configures required packages
func (h *HDMIOutput) ensureDisplayEnvironment() error {
	nativeLogger.Info().Msg("Setting up display environment")
	
	// Update package list
	cmd := exec.Command("apt-get", "update")
	if err := cmd.Run(); err != nil {
		nativeLogger.Warn().Err(err).Msg("Failed to update package list")
	}
	
	// Install required packages
	packages := []string{
		"xorg",
		"openbox",
		"chromium-browser",
		"unclutter",
		"xdotool",
	}
	
	for _, pkg := range packages {
		nativeLogger.Info().Str("package", pkg).Msg("Installing package")
		cmd := exec.Command("apt-get", "install", "-y", pkg)
		if err := cmd.Run(); err != nil {
			nativeLogger.Warn().Str("package", pkg).Err(err).Msg("Failed to install package")
		}
	}
	
	return nil
}

// startDisplayOutput starts X server and browser for HDMI output
func (h *HDMIOutput) startDisplayOutput() error {
	nativeLogger.Info().Msg("Starting display output")
	
	// Create X server configuration
	if err := h.createXConfig(); err != nil {
		return fmt.Errorf("failed to create X config: %w", err)
	}
	
	// Start X server
	if err := h.startXServer(); err != nil {
		return fmt.Errorf("failed to start X server: %w", err)
	}
	
	// Wait for X server to be ready
	time.Sleep(3 * time.Second)
	
	// Start window manager
	if err := h.startWindowManager(); err != nil {
		return fmt.Errorf("failed to start window manager: %w", err)
	}
	
	// Start browser in kiosk mode
	if err := h.startBrowser(); err != nil {
		return fmt.Errorf("failed to start browser: %w", err)
	}
	
	return nil
}

// createXConfig creates X server configuration for HDMI output
func (h *HDMIOutput) createXConfig() error {
	configDir := "/tmp/jetkvm-x11"
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}
	
	configContent := fmt.Sprintf(`
Section "ServerLayout"
    Identifier "JetKVM Layout"
    Screen 0 "JetKVM Screen" 0 0
EndSection

Section "Monitor"
    Identifier "HDMI Monitor"
    Option "DPMS" "false"
EndSection

Section "Device"
    Identifier "Graphics Card"
    Driver "modesetting"
EndSection

Section "Screen"
    Identifier "JetKVM Screen"
    Device "Graphics Card"
    Monitor "HDMI Monitor"
    DefaultDepth 24
    SubSection "Display"
        Depth 24
        Modes "%s"
    EndSubSection
EndSection
`, h.resolution)
	
	configPath := filepath.Join(configDir, "xorg.conf")
	return os.WriteFile(configPath, []byte(configContent), 0644)
}

// startXServer starts the X server
func (h *HDMIOutput) startXServer() error {
	nativeLogger.Info().Msg("Starting X server")
	
	cmd := exec.CommandContext(h.ctx, "Xorg", 
		":1", 
		"-config", "/tmp/jetkvm-x11/xorg.conf",
		"-logfile", "/tmp/jetkvm-xorg.log",
		"vt7")
	
	if err := cmd.Start(); err != nil {
		return err
	}
	
	// Set DISPLAY environment variable
	os.Setenv("DISPLAY", ":1")
	
	return nil
}

// startWindowManager starts a minimal window manager
func (h *HDMIOutput) startWindowManager() error {
	nativeLogger.Info().Msg("Starting window manager")
	
	cmd := exec.CommandContext(h.ctx, "openbox", "--config-file", "/dev/null")
	cmd.Env = append(os.Environ(), "DISPLAY=:1")
	
	if err := cmd.Start(); err != nil {
		return err
	}
	
	// Hide cursor
	hideCursorCmd := exec.CommandContext(h.ctx, "unclutter", "-display", ":1", "-idle", "0")
	hideCursorCmd.Start()
	
	return nil
}

// startBrowser starts browser in kiosk mode
func (h *HDMIOutput) startBrowser() error {
	nativeLogger.Info().Str("url", h.webServerURL).Msg("Starting browser in kiosk mode")
	
	browserArgs := []string{
		"--no-sandbox",
		"--disable-dev-shm-usage",
		"--disable-gpu",
		"--disable-software-rasterizer",
		"--disable-background-timer-throttling",
		"--disable-backgrounding-occluded-windows",
		"--disable-renderer-backgrounding",
		"--disable-features=TranslateUI",
		"--disable-ipc-flooding-protection",
		"--kiosk",
		"--incognito",
		"--noerrdialogs",
		"--disable-session-crashed-bubble",
		"--disable-infobars",
		"--autoplay-policy=no-user-gesture-required",
		h.webServerURL,
	}
	
	cmd := exec.CommandContext(h.ctx, "chromium-browser", browserArgs...)
	cmd.Env = append(os.Environ(), "DISPLAY=:1")
	
	h.browserProcess = cmd
	
	if err := cmd.Start(); err != nil {
		return err
	}
	
	// Auto-refresh page every 30 seconds to keep it alive
	go h.autoRefreshBrowser()
	
	return nil
}

// autoRefreshBrowser periodically refreshes the browser to keep it responsive
func (h *HDMIOutput) autoRefreshBrowser() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			if h.enabled && h.browserProcess != nil {
				// Send F5 key to refresh
				cmd := exec.Command("xdotool", "search", "--name", "chromium", "key", "F5")
				cmd.Env = append(os.Environ(), "DISPLAY=:1")
				cmd.Run()
			}
		case <-h.ctx.Done():
			return
		}
	}
}

// IsHDMIOutputEnabled returns whether HDMI output is currently enabled
func IsHDMIOutputEnabled() bool {
	return hdmiOutput != nil && hdmiOutput.enabled
}

// GetHDMIOutputStatus returns the current HDMI output status
func GetHDMIOutputStatus() map[string]interface{} {
	if hdmiOutput == nil {
		return map[string]interface{}{
			"enabled":    false,
			"available":  false,
			"error":      "HDMI output not initialized",
		}
	}
	
	return map[string]interface{}{
		"enabled":       hdmiOutput.enabled,
		"available":     true,
		"display_device": hdmiOutput.displayDevice,
		"resolution":    hdmiOutput.resolution,
		"web_url":       hdmiOutput.webServerURL,
	}
}