/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 */

// @lc code=start
func isPowerOfFour(n int) bool {
	if n == 1 {
		return true
	}
	if n == 0 {
		return false
	}
	if n & (n-1) != 0 {
		return false
	}
	for n != 0 {
		n >>= 2
		if n & 1 == 1 {
			return true
		}
	}
	return false
}
// @lc code=end

