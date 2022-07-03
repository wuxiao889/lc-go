package nthUglyNumber

func nthUglyNumber1(n int) int {
	//1base,存丑数
	ans := make([]int, n+1)
	ans[1] = 1
	//i2,i3,i5分别表示三个质因数的指针
	i2, i3, i5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		//a,b,c表示三个质因数指针指向的丑数
		a, b, c := ans[i2]*2, ans[i3]*3, ans[i5]*5
		//下一个丑数
		val := min(a, b, c)
		ans[i] = val
		//指针后移，去重
		if a == val {
			i2++
		}
		if b == val {
			i3++
		}
		if c == val {
			i5++
		}
	}
	return ans[n]
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if a[i] < res {
			res = a[i]
		}
	}
	return res
}
