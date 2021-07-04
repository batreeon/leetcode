/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */

// @lc code=start
func convertToTitle(columnNumber int) string {
	var result string
	for columnNumber > 0 {
		columnNumber--
		result = string(columnNumber%26 + 'A') + result
		columnNumber /= 26
	}
	return result
}
// @lc code=end

