/*
 * @lc app=leetcode.cn id=463 lang=golang
 *
 * [463] 岛屿的周长
 */

// @lc code=start
func islandPerimeter(grid [][]int) int {
	var ans int
	for i := 0;i < len(grid);i++ {
		for j := 0;j < len(grid[i]);j++ {
			if grid[i][j] == 0 {
				continue
			}
			if len(grid[i]) == 1{
				if grid[i][j] == 1 {
					ans++
					ans++
				}else {

				}
				continue
			}
			if j == 0 {
				if grid[i][j+1] == 1 {
					ans++
				}else {
					ans++
					ans++
				}
			}else if j != len(grid[i])-1 {
				if grid[i][j-1] == 1 {

				}else {
					ans++
				}
				if grid[i][j+1] == 1{

				}else {
					ans++
				}
			}else {
				if grid[i][j-1] == 1 {
					ans++
				}else {
					ans++
					ans++
				}
			}
		}
	}
	for j := 0;j < len(grid[0]);j++ {
		for i := 0;i < len(grid);i++ {
			if grid[i][j] == 0 {
				continue
			}
			if len(grid) == 1{
				if grid[i][j] == 1 {
					ans++
					ans++
				}else {
					
				}
				continue
			}
			if i == 0 {
				if grid[i+1][j] == 1{
					ans++
				}else {
					ans++
					ans++
				}
			}else if i != len(grid)-1 {
				if grid[i-1][j] == 1{

				}else {
					ans++
				}
				if grid[i+1][j] == 1{

				}else {
					ans++
				}
			}else {
				if grid[i-1][j] == 1{
					ans++
				}else {
					ans++
					ans++
				}
			}
		}
	}
	return ans
}
// @lc code=end

