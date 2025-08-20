/*
 * @lc app=leetcode.cn id=1277 lang=golang
 *
 * [1277] 统计全为 1 的正方形子矩阵
 */

// @lc code=start
func countSquares(matrix [][]int) int {
    r, c := len(matrix), len(matrix[0])
	dp := make([][]int, r+1)
	for i := 0; i < r+1; i++ {
		dp[i] = make([]int, c+1)
	}
	result := 0
	for i, row := range matrix {
		for j, x := range row {
			if x == 1 {
				dp[i+1][j+1] = min(dp[i+1][j], dp[i][j+1], dp[i][j]) + 1
				result += dp[i+1][j+1]
			}
		}
	}
	return result
}
// @lc code=end

