package subsets

import "fmt"

func ExampleSubSets() {
	fmt.Println(subsets([]int{1, 2, 3}))
	//output:
	//[[] [1] [1 2] [1 2 3] [1 3] [2] [2 3] [3]]
}
