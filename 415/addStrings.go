package addStrings

func addStrings(num1 string, num2 string) string {
	if len(num1) > len(num2) {
		return addStrings(num2, num1)
	}
	res := make([]byte, 0, len(num2))
	var o byte = '0'
	i, j := len(num1)-1, len(num2)-1
	for i >= 0 {
		add := num1[i] + num2[j] + o - '0'*3
		if add > 9 {
			o = '1'
			res = append(res, add-10+'0')
		} else {
			o = '0'
			res = append(res, add+'0')
		}
		i--
		j--
	}
	for j >= 0 {
		add := num2[j] + o - '0'*2
		if add > 9 {
			o = '1'
			res = append(res, add-10+'0')
		} else {
			o = '0'
			res = append(res, add+'0')
		}
		j--
	}
	if o != '0' {
		res = append(res, o)
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
