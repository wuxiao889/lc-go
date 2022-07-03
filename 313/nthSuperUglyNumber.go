package nthSuperUglyNumber

func nthSuperUglyNumber(n int, primes []int) int {
	ans := make([]int, n+1)
	ans[1] = 1
	p := make([]int, 1000)
	for _, v := range primes {
		p[v] = 1
	}
	pp := make([]int, len(primes))
	for i := 2; i <= n; i++ {
		for i, v := range primes {
			pp[i] = ans[p[v]] * v
		}
		val := min(pp...)
		for i, v := range primes {
			if pp[i] == val {
				p[v]++
			}
		}
		ans[i] = val
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
