/*
 * @lc app=leetcode.cn id=1765 lang=golang
 *
 * [1765] 地图中的最高点
 */

// @lc code=start
package main
func highestPeak(isWater [][]int) [][]int {
	r,c := len(isWater),len(isWater[0])
	seen := make([]bool,r*c)
	// 存入用于bfs的下标
	queue := []int{}
	for i := 0 ; i < r ; i++ {
		for j := 0 ; j < c ; j++ {
			if isWater[i][j] == 1 {
				seen[i*c+j] = true
				queue = append(queue,i*c+j)
			}
		}
	}
	// 下一步要填的数字
	num := 1
	// 保存返回的结果
	result := make([][]int,r)
	for i := 0 ; i < r ; i++ {
		result[i] = make([]int,c)
	}
	// 四周临近点
	ori := make([][]int,4)
	ori[0] = []int{-1,0}
	ori[1] = []int{0,-1}
	ori[2] = []int{0,1}
	ori[3] = []int{1,0}
	for len(queue) > 0 {
		l := len(queue)
		for i := 0 ; i < l ; i++ {
			x,y := queue[i]/c,queue[i]%c
			for j := 0 ; j < 4 ; j++ {
				newx,newy := x+ori[j][0],y+ori[j][1]
				if newx >= 0 && newx < r && newy >= 0 && newy < c {
					if !seen[newx*c+newy] {
						seen[newx*c+newy] = true
						queue = append(queue,newx*c+newy)
						result[newx][newy] = num
					}
				}
			}
		}
		queue = queue[l:]
		num++
	}
	// 说实在的这里并没有采取什么特别措施保证最后的最高高度最大
	return result
}
// @lc code=end

