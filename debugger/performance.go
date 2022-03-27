package debugger

import (
	"fmt"
	"runtime"
	"time"

	"github.com/foamzou/audio-get/logger"
)

var appStartTime time.Time

func AppStart() {
	PrintMemUsage("App start")
	appStartTime = time.Now()
}

func AppEnd() {
	PrintMemUsage("App end")
	logger.Debug(fmt.Sprintf("Cost time: %s", time.Since(appStartTime).String()))
}

func PrintMemUsage(str string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	info := ""
	info += fmt.Sprintf("---------------------------%s----------------------------\n", str)
	info += fmt.Sprintf("\tAlloc = %v MiB", bToMb(m.Alloc))
	info += fmt.Sprintf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	info += fmt.Sprintf("\tSys = %v MiB", bToMb(m.Sys))
	info += fmt.Sprintf("\tNumGC = %v\n", m.NumGC)
	info += "\n--------------------------------------------------------------"

	logger.Debug(info)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
