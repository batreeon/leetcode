/*
 * @lc app=leetcode.cn id=1446 lang=golang
 *
 * [1446] 连续字符
 */

// @lc code=start
func maxPower(s string) int {
	result := 0
	l := len(s)
	for i := 0; l - i > result; {
		start := i
		j := start + 1
		for ; j < l; j++ {
			if s[j] != s[start] {
				break
			}
		}
		if j - start > result {
			result = j - start
		}
		i = j
	}
	return result
}
// @lc code=end

