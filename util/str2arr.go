package util

import (
	"regexp"
	"strconv"
	"strings"
)

func Str2A2rr(s string) [][]int {
	var err error
	s = s[1 : len(s)-1]
	compile := regexp.MustCompile(`\[([^\]]+)\]`)

	match := compile.FindAllString(s, -1)
	if match == nil {
		return [][]int{}
	}
	for i := range match {
		match[i] = match[i][1 : len(match[i])-1]
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

func Str2Arr(s string) []int {
	var err error
	compile := regexp.MustCompile(`\[([^\]]+)\]`)

	m := compile.FindSubmatch([]byte(s))
	if m == nil {
		return []int{}
	}
	match := string(m[1])
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
