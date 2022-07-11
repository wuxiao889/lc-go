package reversePrefix

import (
	"bytes"
)

func reversePrefix(word string, ch byte) string {
	wb := []byte(word)
	indexByte := bytes.IndexByte(wb, ch)
	if indexByte == -1 {
		return word
	}
	for i, j := 0, indexByte; i < j; i, j = i+1, j-1 {
		wb[i], wb[j] = wb[j], wb[i]
	}
	return string(wb)
}
