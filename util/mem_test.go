package util

import (
	"testing"
	"time"
)

func Test_printMem(t *testing.T) {
	_ = make([]int, 1024*128*100)
	printMem()
	time.Sleep(10 * time.Second)
}
