/*
 * @lc app=leetcode.cn id=516 lang=golang
 *
 * [516] 最长回文子序列
 */

// @lc code=start
func longestPalindromeSubseq(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	dp := make([][]int,n)
	for i := range dp {
		dp[i] = make([]int,n)
	}
	// 注意循环顺序，从短到长
	for i := n-1 ; i >= 0 ; i-- {
		dp[i][i] = 1
		for j := i+1 ; j < n ; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			}else{
				dp[i][j] = max(dp[i+1][j],dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}
func max(x,y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

