/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
package main
import "strconv"

func numDecodings(s string) int {
	// 0没有对应的，但是10，20有对应的
	l := len(s)
	if l == 0 {
		return 0
	}
	if s[l-1] == '0' && l == 1{
		return 0
	}
	if s[l-1] == '0' && (s[l-2] != '1' && s[l-2] != '2') {
		return 0
	}
	dp := make([]int,l+1)
	dp[l] = 1
	for i := l-1 ; i >= 0 ; i-- {
		if s[i] == '0' {
			continue
		}
		res := dp[i+1]
		for j := i+1 ; j < l ; j++ {
			num,_ := strconv.Atoi(s[i:j+1])
			if num >= 0 && num <= 26 {
				res += dp[j+1]
			}
			if num >= 26 {//后面的分割没必要进行了，比26大了都
				break
			}
		}
		dp[i] = res
	}
	return dp[0]
}
// @lc code=end

