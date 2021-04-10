/*
 * @lc app=leetcode.cn id=1042 lang=golang
 *
 * [1042] 不邻接植花
 */

// @lc code=start
func gardenNoAdj(n int, paths [][]int) []int {

	// 颜色池，每个花园最初都有四种颜色可选
	// 后续若一个花园选了一种颜色，则将该颜色从邻居花园的颜色池中删掉
	colorPool := make([][]bool,n+1)
	for i := 1 ; i < n+1 ; i++ {
		colorPool[i] = []bool{false,true,true,true,true}
	}

	// 记录每个花园的邻居
	neighbors := make([][]int,n+1)
	for _,path := range paths {
		neighbors[path[0]] = append(neighbors[path[0]],path[1])
		neighbors[path[1]] = append(neighbors[path[1]],path[0])
	}

	// 从第一个花园开始，从其颜色池中选一个颜色（我们选标号小的），
	// 然后依次更新邻居的颜色池
	// color记录每个花园所选的颜色
	color := make([]int,n+1)
	for i := 1 ; i < n+1 ; i++ {
		// 查找颜色池中的可选颜色
		for j := 1 ; j < n+1 ; j++ {
			if colorPool[i][j] == true {
				color[i] = j
				// 选了颜色j，则将该颜色从邻居的颜色池中删掉
				for _,neighbor := range neighbors[i] {
					colorPool[neighbor][j] = false
				}
				break
			}
		}
	}
	return color[1:]

	// 这道题每个花园最多只有三个邻居，因此复杂度不会过于爆炸
}
// @lc code=end

