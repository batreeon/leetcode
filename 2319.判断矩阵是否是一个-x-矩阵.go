/*
 * @lc app=leetcode.cn id=2319 lang=golang
 *
 * [2319] 判断矩阵是否是一个 X 矩阵
 */

// @lc code=start
func checkXMatrix(grid [][]int) bool {
	n := len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || i + j == n-1 {
				if grid[i][j] == 0 {
					return false
				}
			}else{
				if grid[i][j] != 0 {
					return false
				}
			}
		}
	}
	return true
}
// @lc code=end

