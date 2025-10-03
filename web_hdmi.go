package kvm

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

// HDMI Output Web API handlers

// handleHDMIOutputStatus returns the current HDMI output status
func handleHDMIOutputStatus(c *gin.Context) {
	// Check if platform supports HDMI output
	if runtime.GOARCH != "amd64" || runtime.GOOS != "linux" {
		c.JSON(http.StatusOK, gin.H{
			"enabled":   false,
			"available": false,
			"error":     "HDMI output not supported on this platform",
			"platform":  runtime.GOOS + "/" + runtime.GOARCH,
		})
		return
	}

	status := getHDMIOutputStatus()
	c.JSON(http.StatusOK, status)
}

// handleHDMIOutputEnable enables HDMI output
func handleHDMIOutputEnable(c *gin.Context) {
	// Check if platform supports HDMI output
	if runtime.GOARCH != "amd64" || runtime.GOOS != "linux" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "HDMI output not supported on this platform",
			"platform": runtime.GOOS + "/" + runtime.GOARCH,
		})
		return
	}

	logger.Info().Msg("Web API: Enabling HDMI output")

	if err := rpcEnableHDMIOutput(); err != nil {
		logger.Error().Err(err).Msg("Failed to enable HDMI output")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Update config to remember the setting
	config.HDMIOutputEnabled = true
	if err := SaveConfig(); err != nil {
		logger.Warn().Err(err).Msg("Failed to save HDMI output config")
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "HDMI output enabled successfully",
		"status":  getHDMIOutputStatus(),
	})
}

// handleHDMIOutputDisable disables HDMI output
func handleHDMIOutputDisable(c *gin.Context) {
	// Check if platform supports HDMI output
	if runtime.GOARCH != "amd64" || runtime.GOOS != "linux" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "HDMI output not supported on this platform",
			"platform": runtime.GOOS + "/" + runtime.GOARCH,
		})
		return
	}

	logger.Info().Msg("Web API: Disabling HDMI output")

	if err := rpcDisableHDMIOutput(); err != nil {
		logger.Error().Err(err).Msg("Failed to disable HDMI output")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Update config to remember the setting
	config.HDMIOutputEnabled = false
	if err := SaveConfig(); err != nil {
		logger.Warn().Err(err).Msg("Failed to save HDMI output config")
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "HDMI output disabled successfully",
		"status":  getHDMIOutputStatus(),
	})
}

// handleHDMIOutputToggle toggles HDMI output on/off
func handleHDMIOutputToggle(c *gin.Context) {
	// Check if platform supports HDMI output
	if runtime.GOARCH != "amd64" || runtime.GOOS != "linux" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "HDMI output not supported on this platform",
			"platform": runtime.GOOS + "/" + runtime.GOARCH,
		})
		return
	}

	logger.Info().Msg("Web API: Toggling HDMI output")

	if err := rpcToggleHDMIOutput(); err != nil {
		logger.Error().Err(err).Msg("Failed to toggle HDMI output")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Update config based on current status
	status := getHDMIOutputStatus()
	if enabled, ok := status["enabled"].(bool); ok {
		config.HDMIOutputEnabled = enabled
		if err := SaveConfig(); err != nil {
			logger.Warn().Err(err).Msg("Failed to save HDMI output config")
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "HDMI output toggled successfully",
		"status":  status,
	})
}

// HDMI Output configuration API

type HDMIOutputConfigRequest struct {
	Enabled   bool `json:"enabled"`
	AutoStart bool `json:"auto_start"`
}

// handleHDMIOutputConfig handles HDMI output configuration
func handleHDMIOutputConfig(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		c.JSON(http.StatusOK, gin.H{
			"enabled":    config.HDMIOutputEnabled,
			"auto_start": config.HDMIOutputAutoStart,
			"available":  runtime.GOARCH == "amd64" && runtime.GOOS == "linux",
			"platform":   runtime.GOOS + "/" + runtime.GOARCH,
		})
	case "POST", "PUT":
		var req HDMIOutputConfigRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Update configuration
		config.HDMIOutputEnabled = req.Enabled
		config.HDMIOutputAutoStart = req.AutoStart

		if err := SaveConfig(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save configuration"})
			return
		}

		// Apply the setting immediately if on supported platform
		if runtime.GOARCH == "amd64" && runtime.GOOS == "linux" {
			if req.Enabled {
				if err := rpcEnableHDMIOutput(); err != nil {
					logger.Error().Err(err).Msg("Failed to enable HDMI output after config change")
				}
			} else {
				if err := rpcDisableHDMIOutput(); err != nil {
					logger.Error().Err(err).Msg("Failed to disable HDMI output after config change")
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"message":    "HDMI output configuration updated",
			"enabled":    config.HDMIOutputEnabled,
			"auto_start": config.HDMIOutputAutoStart,
		})
	default:
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
	}
}