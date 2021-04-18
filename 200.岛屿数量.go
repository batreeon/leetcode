/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
    r,c := len(grid),len(grid[0])
    visited := make([]bool,r*c)

	oris := make([][]int,4)
	oris[0] = []int{-1,0}
	oris[1] = []int{0,-1}
	oris[2] = []int{0,1}
	oris[3] = []int{1,0}
	var dfs func(x,y int)
	dfs = func(x,y int) {
		visited[x*c+y] = true
		for _,ori := range oris {
			newx,newy := x+ori[0],y+ori[1]
			if newx >= 0 && newx < r && newy >= 0 && newy < c && grid[newx][newy] == byte('1') && !visited[newx*c+newy] {
				dfs(newx,newy)
			}
		}
	}
	result := 0
	for i := 0 ; i < r ; i++ {
		for j := 0 ; j < c ; j++ {
			if grid[i][j] == byte('1') && !visited[i*c+j] {
				dfs(i,j)
				result++
			}
		}
	}

	return result
}
// @lc code=end

