package util

import (
	"regexp"
	"strconv"
	"strings"
)

func StrToA2rr(s string) [][]int {
	var err error
	compile := regexp.MustCompile(`\d[^\]]+`)

	match := compile.FindAllString(s, -1)
	for i := range match {
		match[i] = strings.Replace(match[i], ",", " ", -1)
	}

	res := make([][]int, len(match))
	for i := range match {
		fields := strings.Fields(match[i])
		res[i] = make([]int, len(fields))
		for j := range res[i] {
			res[i][j], err = strconv.Atoi(fields[j])
			if err != nil {
				panic(err)
			}
		}
	}
	return res
}

func StrToArr(s string) []int {
	var err error
	compile := regexp.MustCompile(`\d[^\]]+`)

	match := compile.FindString(s)
	match = strings.Replace(match, ",", " ", -1)
	fields := strings.Fields(match)
	res := make([]int, len(fields))
	for i := range res {
		res[i], err = strconv.Atoi(fields[i])
		if err != nil {
			panic(err)
		}
	}
	return res
}
