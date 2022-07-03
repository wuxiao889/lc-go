package diffWaysToCompute

import "strconv"

func diffWaysToCompute(expression string) []int {
	if i, err := strconv.Atoi(expression); err == nil {
		return []int{i}
	}
	var res []int
	for i := range expression {
		c := expression[i]
		if c == '+' || c == '-' || c == '*' {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			var temp int
			for _, l := range left {
				for _, r := range right {
					switch c {
					case '+':
						temp = l + r
					case '-':
						temp = l - r
					case '*':
						temp = l * r
					}
					res = append(res, temp)
				}
			}
		}
	}
	return res
}
