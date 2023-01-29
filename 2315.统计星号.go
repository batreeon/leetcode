/*
 * @lc app=leetcode.cn id=2315 lang=golang
 *
 * [2315] 统计星号
 */

// @lc code=start
func countAsterisks(s string) int {
	result := 0
	out := false
	for _, b := range s {
		if b == '*' && !out {
			result++
		}
		if b == '|' {
			out = !out
		}
	}
	return result
}
// @lc code=end

