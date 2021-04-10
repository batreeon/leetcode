/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */

// @lc code=start
func isUgly(n int) bool {
	// 负数一定不是丑数
	// 1，2，3，4，5都是丑数
	if n <= 0 {
		return false
	}
	if n <= 5 {
		return true
	}
	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}
	return n == 1
}
// @lc code=end

