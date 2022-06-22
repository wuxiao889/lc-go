package main

import "fmt"

func firstUniqChar(s string) byte {
	arr := [26]int{}
	for _, c := range s {
		arr[c-'a']++
	}
	fmt.Println(arr)
	for _, c := range s {
		if arr[c-'a'] == 1 {
			return byte(c)
		}
	}
	return ' '
}
