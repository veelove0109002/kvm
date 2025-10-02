//go:build !cgo

package native

import (
	"log"
	"time"
)

// Mock channels for non-CGO builds
var (
	indevEventChan  = make(chan int, 100)
	videoStateChan  = make(chan VideoState, 100)
	videoFrameChan  = make(chan []byte, 100)
	logChan         = make(chan nativeLogMessage, 100)
	rpcEventChan    = make(chan interface{}, 100)
)

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

func (n *Native) handleVideoStateChan() {
	log.Println("Mock: handleVideoStateChan started (no-CGO)")
	for {
		select {
		case state := <-videoStateChan:
			log.Printf("Mock: Received video state change: %+v", state)
			n.onVideoStateChange(state)
		case <-time.After(10 * time.Second):
			// Simulate periodic video state updates
			mockState := VideoState{
				Ready:          true,
				Width:          1920,
				Height:         1080,
				FramePerSecond: 30.0,
			}
			log.Println("Mock: Simulating video state update")
			n.onVideoStateChange(mockState)
		}
	}
}

func (n *Native) handleVideoFrameChan() {
	log.Println("Mock: handleVideoFrameChan started (no-CGO)")
	for {
		select {
		case frame := <-videoFrameChan:
			log.Printf("Mock: Received video frame of size: %d bytes", len(frame))
			n.onVideoFrameReceived(frame, time.Since(time.Now()))
		case <-time.After(33 * time.Millisecond):
			// Simulate 30 FPS video frames
			mockFrame := make([]byte, 1920*1080*3) // Mock RGB frame
			n.onVideoFrameReceived(mockFrame, 33*time.Millisecond)
		}
	}
}

func (n *Native) handleLogChan() {
	log.Println("Mock: handleLogChan started (no-CGO)")
	for {
		select {
		case logMsg := <-logChan:
			log.Printf("Mock: Native log [%s]: %s", logMsg.Level.String(), logMsg.Message)
		case <-time.After(1 * time.Second):
			// Periodic log simulation
			continue
		}
	}
}

func (n *Native) handleRpcEventChan() {
	log.Println("Mock: handleRpcEventChan started (no-CGO)")
	for {
		select {
		case event := <-rpcEventChan:
			log.Printf("Mock: Received RPC event: %+v", event)
			n.onRpcEvent(event.(string))
		case <-time.After(15 * time.Second):
			// Simulate periodic RPC events
			log.Println("Mock: Simulating RPC event")
			n.onRpcEvent("mock_rpc_event")
		}
	}
}