package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("%p\n", &s)
	fmt.Printf("%p\n", s)
	fmt.Println(s[0])
	var s1 []int
	s1 = []int{}
	fmt.Printf("%p\n", &s1)
	fmt.Printf("%p\n", s1)
	//s := []int{1, 2}
	//s1 := s
	//fmt.Printf("%p %p\n", s, s1)
	//fmt.Printf("%p %p\n", &s, &s1)
	//fmt.Println("----")
	//fmt.Printf("adress of 0th element:%p \n", s)
	//fmt.Println(s)
	//s = editslice(s)
	//
	//s = append(s, []int{1, 1, 1, 1, 1, 1, 1, 1}...)
	//fmt.Printf("adress of 0th element:%p \n", s)
	slice1 := []int{10, 20, 30, 40}
	fmt.Printf("slice1:adress of the slice  :%p\n", &slice1)
	fmt.Printf("slice1:adress of 0th element:%p \n", slice1)
	fmt.Printf("slice1:adress of 0th element:%p \n", &slice1[0])
	fmt.Printf("slice1:adress of 0th element:%p \n", &slice1[1])
	slice2 := slice1
	fmt.Printf("slice2:adress of the slice  :%p\n", &slice2)
	fmt.Printf("slice2:adress of 0th element:%p \n", slice2)
	fmt.Printf("slice2:adress of 0th element:%p \n", &slice2[0])
	fmt.Printf("slice2:adress of 0th element:%p \n", &slice2[1])

	for index, value := range slice1 {
		fmt.Printf("value = %d , value-addr = %p , slice-addr = %p\n", value, &value, &slice1[index])
	}
}

func editslice(s []int) []int {
	fmt.Printf("adress of 0th element:%p \n", s)
	for i := range s {
		s[i]++
	}
	fmt.Println(s)
	return s
}
