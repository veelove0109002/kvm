//go:build linux && amd64

package main

import "fmt"

// Mock implementation of gspt.SetProcTitle for X86_64
func setProcTitle(title string) {
	// Mock implementation - just log the title change
	fmt.Printf("Mock: Setting process title to: %s\n", title)
}