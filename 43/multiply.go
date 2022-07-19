package multiply

func multiply(num1 string, num2 string) string {
	//byte uint8 string[i]
	res := make([]int, len(num1)+len(num2))
	for i, x := range num1 {
		for j, y := range num2 {
			tmp := (x - '0') * (y - '0')
			res[i+j] += int(tmp / 10)
			res[i+j+1] += int(tmp % 10)
		}
	}
	for k := len(res) - 1; k > 0; k-- {
		if res[k] > 9 {
			res[k-1] += res[k] / 10
			res[k] = res[k] % 10
		}
	}
	s := make([]byte, len(num1)+len(num2))
	for i := range s {
		s[i] = byte(res[i] + '0')
	}
	for len(s) > 1 && s[0] == '0' {
		s = s[1:]
	}
	return string(s)
}

//func multiply(num1 string, num2 string) string {
//	if len(num1) > len(num2) {
//		return multiply(num2, num1)
//	}
//	res := "0"
//	for i := len(num1) - 1; i >= 0; i-- {
//		var o byte = '0'
//		b := num1[i]
//		if b == '0' {
//			continue
//		}
//		tmp := make([]byte, 0, len(num2))
//		for j := 0; j < len(num1)-1-i; j++ {
//			tmp = append(tmp, '0')
//		}
//		for j := len(num2) - 1; j >= 0; j-- {
//			b2 := num2[j]
//			mul := (b2-'0')*(b-'0') + o - '0'
//			if mul > 9 {
//				o = '0' + mul/10
//				mul = mul % 10
//				tmp = append(tmp, mul+'0')
//			} else {
//				o = '0'
//				tmp = append(tmp, mul+'0')
//			}
//		}
//		if o != '0' {
//			tmp = append(tmp, o)
//		}
//		for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
//			tmp[i], tmp[j] = tmp[j], tmp[i]
//		}
//		fmt.Println(string(tmp))
//		res = addStrings(res, string(tmp))
//	}
//	return res
//}
//
//func addStrings(num1 string, num2 string) string {
//	if len(num1) > len(num2) {
//		return addStrings(num2, num1)
//	}
//	res := make([]byte, 0, len(num2))
//	var o byte = '0'
//	i, j := len(num1)-1, len(num2)-1
//	for i >= 0 {
//		add := num1[i] + num2[j] + o - '0'*3
//		if add > 9 {
//			o = '1'
//			res = append(res, add-10+'0')
//		} else {
//			o = '0'
//			res = append(res, add+'0')
//		}
//		i--
//		j--
//	}
//	for j >= 0 {
//		add := num2[j] + o - '0'*2
//		if add > 9 {
//			o = '1'
//			res = append(res, add-10+'0')
//		} else {
//			o = '0'
//			res = append(res, add+'0')
//		}
//		j--
//	}
//	if o != '0' {
//		res = append(res, o)
//	}
//	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
//		res[i], res[j] = res[j], res[i]
//	}
//	return string(res)
//}
