package countPrimes

//给定整数 n ，返回 所有小于非负整数 n 的质数的数量 。
func countPrimes(n int) int {
	isPrime := make([]bool, n)
	for i := 2; i*i < n; i++ {
		if !isPrime[i] {
			for j := i; j*i < len(isPrime); j++ {
				isPrime[j*i] = true
			}
		}
	}
	var cnt int
	for i := 2; i < n; i++ {
		if !isPrime[i] {
			cnt++
		}
	}
	return cnt
}
