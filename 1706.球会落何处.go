/*
 * @lc app=leetcode.cn id=1706 lang=golang
 *
 * [1706] 球会落何处
 */

// @lc code=start
func findBall(grid [][]int) []int {
	res := make([]int, len(grid[0]))
	for j := 0; j < len(grid[0]); j++ {
		res[j] = dfs(grid, j)
	}
	return res
}

func dfs(grid [][]int, j int) int {
	m, n := len(grid), len(grid[0])
	i, j := -1, j
	oris := [][]int{{1,0},{0,-1},{0,1}} // 向下，向左，向右
	oriIndex := 0
	for {
		newi, newj := i + oris[oriIndex][0], j + oris[oriIndex][1]
		if newj < 0 || newj >= n {
			return -1
		}
		if newi >= m {
			return j
		} 
		i, j = newi, newj
		if oriIndex == 0 {
			if grid[i][j] == 1 {
				oriIndex = 2
			}else{
				oriIndex = 1
			}
		}else if oriIndex == 1 {
			if grid[i][j] == 1 {
				return -1
			}else{
				oriIndex = 0
			}
		}else{
			if grid[i][j] == 1 {
				oriIndex = 0
			}else{
				return -1
			}
		}
	}
}
// @lc code=end

