/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 */

// @lc code=start

func reverseStr(s string, k int) string {
	var result string
	i := 0
	for ; i+2*k < len(s) ; i += 2*k {
		result += reverse(s[i:i+2*k],k)
	}
	result += reverse(s[i:],k)
	return result
}
func reverse(s string,k int) string {
	if k > len(s) {
		k = len(s)
	}
	bytes := []byte(s)
	for i,j := 0,k-1 ; i < j ; i,j = i+1,j-1 {
		bytes[i],bytes[j] = bytes[j],bytes[i]
	}
	return string(bytes)
}
// @lc code=end

