/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	r,c := len(grid),len(grid[0])
	f := make([]int,c)
	f[0] = grid[0][0]
	for j := 1 ; j < c ; j++ {
		f[j] = grid[0][j] + f[j-1]
	}
	for i := 1 ; i < r ; i++ {
		f[0] += grid[i][0]
		for j := 1 ; j < c ; j++ {
			f[j] = grid[i][j] + min(f[j-1],f[j])
		}
	}
	return f[c-1]
}
// @lc code=end

