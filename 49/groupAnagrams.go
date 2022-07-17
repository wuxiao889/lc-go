package groupAnagrams

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int]*[]string)
	for _, str := range strs {
		var w [26]int
		for _, b := range str {
			w[b-'a']++
		}
		if v, ok := m[w]; ok {
			*v = append(*v, str)
		} else {
			tmp := make([]string, 0)
			m[w] = &tmp
			*m[w] = append(*m[w], str)
		}
	}
	var res [][]string
	for i := range m {
		res = append(res, *m[i])
	}
	return res
}
