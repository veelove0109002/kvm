//go:build linux && arm && cgo && !ci

package main

import "github.com/erikdubbelboer/gspt"

func setProcTitle(title string) {
	gspt.SetProcTitle(title)
}