/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start
func reverseString(s []byte)  {
	l := len(s)
	for i := 0 ; i < l/2 ; i++ {
		s[i],s[l-i-1] = s[l-i-1],s[i]
	}
}
// @lc code=end

