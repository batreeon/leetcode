/*
 * @lc app=leetcode.cn id=2303 lang=golang
 *
 * [2303] 计算应缴税款总额
 */

// @lc code=start
func calculateTax(brackets [][]int, income int) float64 {
	result := float64(0)
	lastUpper := 0
	for _, bracket := range brackets {
		upper, percent := bracket[0], bracket[1]
		if income <= upper {
			result += float64(percent)/float64(100) * float64(income - lastUpper)
			return result
		}
		result += float64(percent)/float64(100) * float64(upper-lastUpper)
		lastUpper = upper
	}
	return result
}
// @lc code=end

