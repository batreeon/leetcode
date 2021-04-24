/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	n1,n2 := len(haystack),len(needle)
	for i := 0 ; i < n1-n2+1 ; i++ {
		if haystack[i:i+n2] == needle {
			return i
		}
	}
	return -1
}
// @lc code=end

