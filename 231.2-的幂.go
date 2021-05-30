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
	// 是2的幂次，那么其二进制中只有一个1，去掉一个1若为0则是2的幂次，否则不是
	return n & (n-1) == 0
}
// @lc code=end

