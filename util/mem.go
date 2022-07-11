package util

import (
	"fmt"
	"runtime"
)

func printMem() {
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	fmt.Printf("%.2f MB\n", float64(rtm.Alloc)/1024./1024.)
}
