package numPrimeArrangements

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	Count()
	for i := 1; i < len(Arr); i++ {
		fmt.Printf("%v : %v \n", i, Arr[i])
	}
	//fmt.Println(1e7)
}

func TestXX(t *testing.T) {
	fmt.Println(numPrimeArrangements(100))
}
