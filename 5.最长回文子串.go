/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool,n)
	for i := 0 ; i < n ; i++ {
		dp[i] = make([]bool,n)
	}
	var result string
	for l := 1 ; l <= n ; l++ {
		for i := 0 ; i+l-1 < n ; i++ {
			j := i+l-1
			if l == 1 {
				dp[i][j] = true
			}else if l == 2 {
				if s[i] == s[j] {
					dp[i][j] = true
				}
			}else{
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && l > len(result) {
				result = s[i:i+l]
			}
		}
	}
	return result
}
// @lc code=end

