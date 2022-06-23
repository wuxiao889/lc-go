package util

import (
	"fmt"
	"testing"
)

var tests = []string{
	"[[1,1,1,0,0],[0,1,1,1,1],[0,0,0,0,0],[1,0,0,0,0],[1,1,0,1,1]]",
	"[[1,1,1,01,0],[0,144,1,1,1],[0,0,0,0,0],[1,0,0,0,0],[1,1,0,1,1]]",
}

var test2 = []string{
	"[1,1,2,5,6,7,10]",
}

func TestStrToA2rr(t *testing.T) {
	for _, tt := range tests {
		fmt.Println(StrToA2rr(tt))
	}
}

func TestStrToArr(t *testing.T) {
	for _, tt := range test2 {
		fmt.Println(StrToArr(tt))
	}
}
