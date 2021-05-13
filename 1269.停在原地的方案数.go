/*
 * @lc app=leetcode.cn id=1269 lang=golang
 *
 * [1269] 停在原地的方案数
 */

// @lc code=start
func numWays(steps int, arrLen int) int {
	const mod = 1e9 + 7
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	// dp[i][j]表示i步到达下标j的方法数
	// 那么第step步就是第step行
	// 能够到达最远的下标是step
	// 但要收arrLen的限制，最远arrLen-1。注意这是下标
	dp := make([][]int, steps+1)
	maxlength := min(arrLen-1, steps)
	for i := range dp {
		dp[i] = make([]int, maxlength+1)
	}
	// 0步到达下标0，方法只有一种，那就是不动
	// dp[0][1...maxlength-1]
	// 三种不同走法:
	// dp[i][j] = dp[i-1][j-1] + dp[i-1][j] + dp[i-1][j+1]
	dp[0][0] = 1
	for i := 1; i <= steps; i++ {
		for j := 0; j <= maxlength; j++ {
			dp[i][j] = dp[i-1][j]
			if j-1 >= 0 {
				dp[i][j] = (dp[i][j] + dp[i-1][j-1]) % mod
			}
			if j+1 <= maxlength {
				dp[i][j] = (dp[i][j] + dp[i-1][j+1]) % mod
			}
		}
	}
	return dp[steps][0]
}

// @lc code=end

