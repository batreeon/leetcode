/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 */

// @lc code=start
func numDistinct(s string, t string) int {
	sl,tl := len(s),len(t)
	dp := make([][]int,sl+1)
	for i := 0 ; i <= sl ; i++ {
		dp[i] = make([]int,tl+1)
		dp[i][tl] = 1
	}
	for i := sl-1 ; i >= 0 ; i-- {
		for j := tl-1 ; j >= 0 ; j-- {
			if s[i] == t[j] {
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			}else{
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	return dp[0][0]
}
// @lc code=end

