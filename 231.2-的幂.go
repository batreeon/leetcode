/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2的幂
 */

// @lc code=start
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	if n & (n-1) == 0 {
		return true
	}
	return false
}
// @lc code=end

