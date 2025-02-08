/*
 * @lc app=leetcode.cn id=712 lang=golang
 *
 * [712] 两个字符串的最小ASCII删除和
 */

// @lc code=start
func minimumDeleteSum(s1 string, s2 string) int {
    n1 := len(s1)
	n2 := len(s2)
	// dp[i][j] 表示 s1[:i+1]和s2[:j+1]的最小删除和
	dp := make([][]int, n1+1)
	for i := 0; i < n1+1; i++ {
		dp[i] = make([]int, n2+1)
	}

	
	for j := 1; j < n2+1; j++ {
		dp[0][j] = dp[0][j-1] + int(s2[j-1])
	}
	for i := 1; i < n1+1; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}

	for i := 1; i < n1+1; i++ {
		for j := 1; j < n2+1; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}else{
				dp[i][j] = min(dp[i-1][j] + int(s1[i-1]), dp[i][j-1]+int(s2[j-1]))
			}
		}
	}
	return dp[n1][n2]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

