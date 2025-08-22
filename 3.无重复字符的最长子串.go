/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
    result := 0
	left := 0
	for right := 1; right < len(s); right++ {
		for i := left; i < right; i++ {
			if s[i] == s[right] {
				left = i+1
			}
		}
		result = max(right-left+1, result)
	}
	return result
}
// @lc code=end

