//go:build linux && arm && (ci || !cgo)

package main

import "fmt"

// Mock implementation of gspt.SetProcTitle for ARM CI/CD builds
func setProcTitle(title string) {
	// Mock implementation - just log the title change for ARM CI builds
	fmt.Printf("Mock: Setting process title to: %s (ARM CI)\n", title)
}