/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	/*
	// 动态规划，到达[i][j]的路径数等于[i-1][j] + [i][j-1]
	r := len(obstacleGrid)
	if r == 0 {
		return 0
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	c := len(obstacleGrid[0])
	dp := make([][]int,r)

	// 处理第一行
	dp[0] = make([]int,c)
	for j := 0 ; j < c ; j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		dp[0][j] = 1
	}

	for i := 1 ; i < r ; i++ {
		dp[i] = make([]int,c)
		// 处理第一列
		if dp[i-1][0] != 0 && obstacleGrid[i][0] != 1 {
			dp[i][0] = 1
		}
		for j := 1 ; j < c ; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			if j >= 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[r-1][c-1]
	*/

	// 官解里面的滑动数组思想，还是很不错的
	r,c := len(obstacleGrid),len(obstacleGrid[0])
	f := make([]int,c)
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	f[0] = 1
	for i := 0 ; i < r ; i++ {
		for j := 0 ; j < c ; j++ {
			if obstacleGrid[i][j] == 1 {
				f[j] = 0
				continue
			}
			// 源f[j]记录着dp[i-1][j],现f[j-1]记录着dp[i][j-1]
			if j >= 1 && obstacleGrid[i][j-1] != 1 {
				f[j] = f[j] + f[j-1]
			}
		}
	}
	return f[c-1]
}	
// @lc code=end

