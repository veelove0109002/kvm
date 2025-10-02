//go:build linux && amd64

package kvm

import (
	"os"
)

// Mock NBD implementation for X86_64
func (d *NBDDevice) runClientConn() {
	d.l.Info().Msg("Mock NBD client for X86_64 - no actual NBD functionality")
}

func (d *NBDDevice) Close() {
	d.l.Info().Msg("Mock NBD close for X86_64")
	if d.dev != nil {
		_ = d.dev.Close()
	}
	if d.listener != nil {
		_ = d.listener.Close()
	}
	if d.clientConn != nil {
		_ = d.clientConn.Close()
	}
	if d.serverConn != nil {
		_ = d.serverConn.Close()
	}
}