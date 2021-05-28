/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	xor := x^y
	result := 0
	for xor != 0 {
		if xor & 1 == 1 {
			result++
		}
		xor >>= 1
	}
	return result
}
// @lc code=end

