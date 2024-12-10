package math

// GCD 求两个整数 a,b 的最大公约数
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

// LCM 求两个整数 a,b 的最小公倍数
func LCM(a, b int) int {
	return a / GCD(a, b) * b
}

// Primes 利用线性筛求 1 ~ n的所有质数 时间复杂度 O(n)
func Primes(n int) []int {
	vis := make([]int, n+1)
	var primes []int
	for i := 2; i <= n; i++ {
		if vis[i] == 0 {
			primes = append(primes, i)
		}
		for _, x := range primes {
			if x*i > n {
				break
			}
			vis[x*i] = 1
			if i%x == 0 {
				break
			}
		}
	}
	return primes
}

// IsPrime 判断一个数 n 是否是质数
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Pow 快速幂求 x 的 n 次方
func Pow(x, n int) int {
	ans := 1
	for ; n > 0; n /= 2 {
		if n&1 == 1 {
			ans *= x
		}
		x *= x
	}
	return ans
}
