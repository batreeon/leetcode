/*
 * @lc app=leetcode.cn id=2309 lang=golang
 *
 * [2309] 兼具大小写的最好英文字母
 */

// @lc code=start
func greatestLetter(s string) string {
	small := make([]int, 26)
	big := make([]int, 26)

	for _, b := range s {
		if b >= 'a' && b <= 'z' {
			small[int(b-'a')] = 1
		}
		if b >= 'A' && b <= 'Z' {
			big[int(b-'A')] = 1
		}
	}
	for i := 25; i >= 0; i-- {
		if small[i] & big[i] == 1 {
			return string('A' + byte(i))
		}
	}
	return ""
}
// @lc code=end

