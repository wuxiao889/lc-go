package main

func countCharacters(words []string, chars string) int {
	// 思路1：hash构筑好了以后，每次copy一份出来做--
	hash := [26]uint8{} // (1) 按题设至多100个字母uint8足够应付
	for _, ch := range chars {
		hash[ch-'a']++
	}
	ret := 0
	// fmt.Println(hash)
	for _, word := range words {
		// fmt.Println(word)
		ret += len(word) // (2) 先蒸
		cp := hash       // (4) Golang中Array赋值即拷贝
		for _, ch := range word {
			if cp[ch-'a'] == 0 {
				ret -= len(word) // (3) 后烤
				break
			}
			cp[ch-'a']--
			// fmt.Println(cp)
		}
	}
	return ret
}
