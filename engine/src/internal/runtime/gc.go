package runtime

import (
	"runtime/debug"
	"time"
)

func StartFreeOSMemoryLoop() {
	timeInterval := 1 * time.Minute
	go func() {
		ticker := time.NewTicker(timeInterval)
		for range ticker.C {
			FreeOSMemory()
		}
	}()
}

func FreeOSMemory() {
	debug.FreeOSMemory()
}
