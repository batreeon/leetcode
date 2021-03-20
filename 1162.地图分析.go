/*
 * @lc app=leetcode.cn id=1162 lang=golang
 *
 * [1162] 地图分析
 */

// @lc code=start
func maxDistance(grid [][]int) int {
	visited := map[int]bool{}
	visitedNum := 0
	stack := []int{}
	r,c := len(grid),len(grid[0])
	for i := 0 ; i < r ; i++ {
		for j := 0 ; j < c ; j++ {
			if grid[i][j] == 1 {
				grid[i][j] = 0
				stack = append(stack,i*c+j)
				visited[i*c+j] = true
				visited++
			}
		}
	}
	if visited == 0 || visited == 1 {
		return -1
	}
	l := len(stack)
	ori := [][]int{
		[]int{-1,0},
		[]int{0,-1},
		[]int{0,1},
		[]int{1,0},
	}
	for i := 0 ; i < l ; i++ {
		x,y := stack[i]/c,stack[i]%c
		for j := 0 ; j < 4 ; j++ {
			newX,newY := x+ori[j][0],y+ori[j][1]
			if _,ok := visited[newX*c+r] ; !ok {
				grid[newX][newY] = grid[]
			}
		}
	}
}
// @lc code=end

