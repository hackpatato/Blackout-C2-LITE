package AntiSandbox

import (
	"runtime"
	"time"
)

func Checksandbox() bool{
	if runtime.NumCPU() < 4 {
		return true
	}
	startTime := time.Now()
	time.Sleep(3 * time.Second)
	duration := time.Since(startTime)
	if duration < 2800*time.Millisecond {
		return true
	}
	return false
}