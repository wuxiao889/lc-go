package addBinary

import (
	"strconv"
)

//长字符串不通过
func addBinary2(a string, b string) string {
	ia, _ := strconv.ParseInt(a, 2, 64)
	ib, _ := strconv.ParseInt(b, 2, 64)
	for ib > 0 {
		ia, ib = ia^ib, (ia&ib)<<1
	}
	return strconv.FormatInt(ia, 2)
}

//0 0 0 0
//1 1 1 3 1 1
//1 1 0 2 0 1
//1 0 0 1 1 0

func addBinary1(a string, b string) string {
	res := make([]byte, 0)
	i, j := len(a)-1, len(b)-1
	o := uint8('0')
	for i >= 0 && j >= 0 {
		r := a[i] + b[j] + o - 3*'0'
		r, o = ro(r)
		res = append(res, r)
		i--
		j--
	}
	for i >= 0 {
		r := a[i] + o - 2*'0'
		r, o = ro(r)
		res = append(res, r)
		i--
	}
	for j >= 0 {
		r := b[j] + o - 2*'0'
		r, o = ro(r)
		res = append(res, r)
		j--
	}
	if o == '1' {
		res = append(res, o)
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func ro(r uint8) (uint8, uint8) {
	var o uint8
	switch r {
	case 3:
		r = '1'
		o = '1'
	case 2:
		r = '0'
		o = '1'
	case 1:
		r = '1'
		o = '0'
	case 0:
		r = '0'
		o = '0'
	}
	return r, o
}

func addBinary(a string, b string) string {
	res := make([]byte, 0)
	ca := uint8(0)
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		sum := ca
		if i >= 0 {
			sum += a[i] - '0'
		}
		if j >= 0 {
			sum += b[j] - '0'
		}
		sum, ca = sum%2, sum/2
		res = append(res, sum+'0')
	}
	if ca == 1 {
		res = append(res, ca+'0')
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
