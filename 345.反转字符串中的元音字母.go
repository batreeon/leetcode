/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */

// @lc code=start
func isVowel(b byte) bool {
	if b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' || b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U' {
		return true
	}
	return false
}
func reverseVowels(s string) string {
	bytes := []byte(s)
	l,r := 0,len(s)-1
	for {
		for l < len(s) && !isVowel(bytes[l]) {l++}
		for r >= 0 && !isVowel(bytes[r]) {r--}
		if l >= r {
			break
		}
		bytes[l],bytes[r] = bytes[r],bytes[l]
		l++
		r--
	}
	return string(bytes)
}
// @lc code=end

