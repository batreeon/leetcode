/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 */

// @lc code=start
func reverseWords(s string) string {
	reverseword := func(w []byte) {
		l := len(w)
		for i := 0 ; i < l/2 ; i++ {
			w[i],w[l-i-1] = w[l-i-1],w[i]
		}
	}
	arrayS := strings.Fields(s)
	for i := 0 ; i < len(arrayS) ; i++ {
		ss := []byte(arrayS[i])
		reverseword(ss)
		arrayS[i] = string(ss)
	}
	return strings.Join(arrayS," ")
}
// @lc code=end

