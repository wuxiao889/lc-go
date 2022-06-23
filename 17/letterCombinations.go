package letterCombinations

var m = [10]string{
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

var res []string
var path []byte
var digit string

func letterCombinations(digits string) []string {
	res = make([]string, 0)
	path = make([]byte, 0)
	digit = digits
	backtrace(digit, 0)
	return res
}

func backtrace(digit string, k int) {
	if k == len(digit) {
		res = append(res, string(path))
		return
	}
	u := digit[k]
	s := m[u-'0']
	for i := 0; i < len(s); i++ {
		path = append(path, s[i])
		backtrace(digit, k+1)
		path = path[:len(path)-1]
	}
}
