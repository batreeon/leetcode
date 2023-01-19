/*
 * @lc app=leetcode.cn id=2351 lang=golang
 *
 * [2351] 第一个出现两次的字母
 */

// @lc code=start
func repeatedCharacter(s string) byte {
	letters := make([]int, 26)
	for _, b := range s {
		if letters[b - 'a'] == 1 {
			return byte(b)
		}
		letters[b - 'a']++
	}
	return 'a'
}
// @lc code=end

