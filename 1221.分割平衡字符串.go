/*
 * @lc app=leetcode.cn id=1221 lang=golang
 *
 * [1221] 分割平衡字符串
 */

// @lc code=start
func balancedStringSplit(s string) int {
	score := 0
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			score--
		}else{
			score++
		}
		if score == 0 {
			result++
		}
	}
	return result
}
// @lc code=end

