/*
 * @lc app=leetcode.cn id=313 lang=golang
 *
 * [313] 超级丑数
 */

// @lc code=start
func nthSuperUglyNumber(n int, primes []int) int {
	nums := []int{1}
	k := len(primes)
	ip := make([]int,k)
	for i := 1 ; i < n ; i++ {
		v := primes[0] * nums[ip[0]]
		for j := 1 ; j < k ; j++ {
			if primes[j] * nums[ip[j]] < v {
				v = primes[j] * nums[ip[j]]
			}
		}
		nums = append(nums,v)
		for j := 0 ; j < k ; j++ {
			for primes[j] * nums[ip[j]] <= v {
				ip[j]++
			}
		}
	}
	return nums[n-1]
}
// @lc code=end

