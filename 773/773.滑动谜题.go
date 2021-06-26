/*
 * @lc app=leetcode.cn id=773 lang=golang
 *
 * [773] 滑动谜题
 */

// @lc code=start
package main
func slidingPuzzle(board [][]int) int {
	// bfs,可以说是很暴力了
	// 和752有异曲同工之妙，bfs可以使每一轮step加一
	getKey := func(b [][]int) string {
		bs := []byte{}
		for i := range b {
			for j := range b[i] {
				// int转byte
				bs = append(bs,byte(b[i][j]))
			}
		}
		return string(bs)
	}

	targetBoard := [][]int{}
	targetBoard = append(targetBoard,[]int{1,2,3})
	targetBoard = append(targetBoard,[]int{4,5,0})
	targetKey := getKey(targetBoard)

	// 记录当前局面有没有遇见过
	visited := map[string]bool{}
	q := [][][]int{}
	q = append(q,board)
	step := 0

	orix,oriy := []int{-1,0,0,1},[]int{0,-1,1,0}
	for len(q) > 0 {
		l := len(q)
		for _,bad := range q {
			if getKey(bad) == targetKey {
				return step
			}
			var i0,j0 int
			for i := range bad {
				for j := range bad[i] {
					if bad[i][j] == 0 {
						i0,j0 = i,j
						break
					}
				}
			}
			for i := 0 ; i < 4 ; i++ {
				ni,nj := i0+orix[i],j0+oriy[i]
				if ni < 0 || ni >= 2 || nj < 0 || nj >= 3 {
					continue
				}

				// 下一步交换后得到的新board
				resbad := make([][]int,2)
				for i := range resbad {
					resbad[i] = make([]int,3)
					copy(resbad[i],bad[i])
				}
				resbad[i0][j0],resbad[ni][nj] = resbad[ni][nj],resbad[i0][j0]
				k := getKey(resbad)
				if visited[k] {
					continue
				}
				visited[k] = true
				q = append(q,resbad)
			}
		}
		step++
		q = q[l:]
	}
	return -1
}
// @lc code=end

