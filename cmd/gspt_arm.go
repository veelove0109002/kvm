//go:build linux && !amd64

package main

import "github.com/erikdubbelboer/gspt"

func setProcTitle(title string) {
	gspt.SetProcTitle(title)
}