//go:build !linux

package main

import "fmt"

// Mock implementation of gspt.SetProcTitle for non-Linux platforms
func setProcTitle(title string) {
	// Mock implementation - just log the title change
	fmt.Printf("Mock: Setting process title to: %s (non-Linux)\n", title)
}