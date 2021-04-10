/*
 * @lc app=leetcode.cn id=1162 lang=golang
 *
 * [1162] 地图分析
 */

// @lc code=start
func maxDistance(grid [][]int) int {
	// 尝试用bfs，先找到所有距陆地距离为1的地点，然后用这些地点继续做更新
	pos := [][]int{}
	visited := map[int]bool{}
	r,c := len(grid),len(grid[0])
	for i := 0 ; i < r ; i++ {
		for j := 0 ; j < c ; j++ {
			if grid[i][j] == 1 {
				visited[i*c+j] = true
				pos = append(pos,[]int{i,j})
			}
		}
	}
	if len(pos) == 0 || len(pos) == r*c {
		return -1
	}

	ori := make([][]int,4)
	ori[0] = []int{-1,0}
	ori[1] = []int{0,-1}
	ori[2] = []int{0,1}
	ori[3] = []int{1,0}

	dis := 0
	for {
		l := len(pos)
		dis++
		for i := 0 ; i < l ; i++ {
			x,y := pos[i][0],pos[i][1]
			for _,orientation := range ori {
				newx,newy := x+orientation[0],y+orientation[1]
				if newx >= 0 && newx < r && newy >= 0 && newy < c && !visited[newx*c+newy] {
					visited[newx*c+newy] = true
					// grid[newx][newy] = dis
					if len(visited) == r*c {
						return dis
					}
					pos = append(pos,[]int{newx,newy})
				}
			}
		}
		pos = pos[l:]
	}

	return -1
}
// @lc code=end

