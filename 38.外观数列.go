/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

// @lc code=start
var ss []string
func countAndSay(n int) string {
	if n == 1 {
		ss = append(ss, string("1"))
	}
	if n <= len(ss) {
		return ss[n-1]
	}
	for len(ss) < n {
		ss = append(ss, describe(ss[len(ss)-1]))
	}
	return ss[n-1]
}

func describe(s string) string {
	var result string
	i := 0
	for j := i; i < len(s); {
		for j < len(s) && s[j] == s[i] {
			j++
		}
		n := strconv.Itoa(j-i)
		result += n
		result += string(s[i])
		i = j
	}
	return result
}
// @lc code=end