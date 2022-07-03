package numPrimeArrangements

var Arr = [100]int{}

func numPrimeArrangements(n int) int {
	Count()
	mod := int(1e9 + 7)
	ans1, ans2 := 1, 1
	k := Arr[n]
	for i := 1; i <= k; i++ {
		ans1 = (ans1 * i) % mod
	}
	for i := 1; i <= n-k; i++ {
		ans2 = (ans2 * i) % mod
	}
	return (ans1 * ans2) % mod
}

func Count() {
	cnt := 0
	for i := 2; i < len(Arr); i++ {
		flag := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				flag = false
			}
		}
		if flag {
			cnt++
		}
		Arr[i] = cnt
	}
}
