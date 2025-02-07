/*
 * @lc app=leetcode.cn id=931 lang=golang
 *
 * [931] 下降路径最小和
 */

// @lc code=start
// 简单动态规划
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
    dp := make([]int, n)
	olddp := make([]int, n)
	for j := 0; j < n; j++ {
		dp[j] = matrix[0][j]
	}

	for i := 1; i < n; i++ {
		copy(olddp, dp)
		for j := 0; j < n; j++ {
			if j == 0 && j == n-1 {
				dp[j] = olddp[j] + matrix[i][j]
			}else if j == 0 {
				dp[j] = min(olddp[j], olddp[j+1]) + matrix[i][j]
			}else if j == n - 1 {
				dp[j] = min(olddp[j-1], olddp[j]) + matrix[i][j]
			}else{
				dp[j] = min(min(olddp[j-1], olddp[j]), olddp[j+1]) + matrix[i][j]
			}
		}
	}

	result := dp[0]
	for j := 1; j < n; j++ {
		result = min(result, dp[j])
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

