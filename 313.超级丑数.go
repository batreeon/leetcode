/*
 * @lc app=leetcode.cn id=313 lang=golang
 *
 * [313] 超级丑数
 */

// @lc code=start
func nthSuperUglyNumber(n int, primes []int) int {
	res := []int{1}
	min := func(x,y int) int {
		if x > y {
			return y
		}
		return x
	}
	idx := make([]int,len(primes))
	for {
		if len(res) == n {
			return res[n-1]
		}
		next := primes[0] * res[idx[0]]
		for i := 1 ; i < len(primes) ; i++ {
			next = min(next,primes[i] * res[idx[i]])
		}
		res = append(res,next)
		for i,prime := range primes {
			for res[idx[i]] * prime <= next {
				idx[i]++
			}
		}
	}
}
// @lc code=end

