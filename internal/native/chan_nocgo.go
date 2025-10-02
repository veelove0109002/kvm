//go:build !cgo

package native

import (
	"log"
	"time"
)

// Mock channel for non-CGO builds
var indevEventChan = make(chan int, 100)

func (n *Native) handleIndevEventChan() {
	log.Println("Mock: handleIndevEventChan started (no-CGO)")
	for {
		select {
		case event := <-indevEventChan:
			log.Printf("Mock: Received indev event: %d", event)
			n.onIndevEvent("mock_event")
		case <-time.After(5 * time.Second):
			// Simulate periodic events for testing
			log.Println("Mock: Simulating indev event")
			n.onIndevEvent("mock_periodic_event")
		}
	}
}