/*
 * @lc app=leetcode.cn id=1210 lang=golang
 *
 * [1210] 穿过迷宫的最少移动次数
 */

// @lc code=start
package main

func minimumMoves(grid [][]int) int {
	n := len(grid)
	if n == 2 {
		return -1
	}

	// 以尾巴所在点为准
	// 记录水平状态
	// h[i][j] 表示到((i,j),(i,j+1)) 的最短移动次数
	h := make([][]int, n)
	for i := 0; i < n; i++ {
		h[i] = make([]int, n-1)
		for j := 0; j < n-1; j++ {
			h[i][j] = -1
		}
	}
	h[0][0] = 0

	// 记录竖直状态
	v := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		v[i] = make([]int, n)
		for j := 0; j < n; j++ {
			v[i][j] = -1
		}
	}

	// 使用bfs记录当前轮次到达的边缘
	statush := [][]int{{0, 0}}
	statusv := [][]int{}

	for len(statush) != 0 || len(statusv) != 0 {
		lh := len(statush)
		lv := len(statusv)

		for ih := 0; ih < lh; ih++ {
			sh := statush[ih]
			x, y := sh[0], sh[1]
			if y < n-2 && grid[x][y+2] == 0 {
				// 水平向右
				old := h[x][y+1]
				h[x][y+1] = min(h[x][y]+1, h[x][y+1])
				if old != h[x][y+1] {
					statush = append(statush, []int{x, y + 1})
				}
			}
			if x < n-1 && grid[x+1][y] == 0 && grid[x+1][y+1] == 0 {
				// 水平向下
				old := h[x+1][y]
				h[x+1][y] = min(h[x][y]+1, h[x+1][y])
				if old != h[x+1][y] {
					statush = append(statush, []int{x + 1, y})
				}

				// 顺时针
				old = v[x][y]
				v[x][y] = min(h[x][y]+1, v[x][y])
				if old != v[x][y] {
					statusv = append(statusv, []int{x, y})
				}
			}
		}

		for iv := 0; iv < lv; iv++ {
			sv := statusv[iv]
			x, y := sv[0], sv[1]
			if x < n-2 && grid[x+2][y] == 0 {
				// 竖直向下
				old := v[x+1][y]
				v[x+1][y] = min(v[x][y]+1, v[x+1][y])
				if old != v[x+1][y] {
					statusv = append(statusv, []int{x + 1, y})
				}
			}
			if y < n-1 && grid[x][y+1] == 0 && grid[x+1][y+1] == 0 {
				// 竖直向右
				old := v[x][y+1]
				v[x][y+1] = min(v[x][y]+1, v[x][y+1])
				if old != v[x][y+1] {
					statusv = append(statusv, []int{x, y + 1})
				}

				// 逆时针
				old = h[x][y]
				h[x][y] = min(v[x][y]+1, h[x][y])
				if old != h[x][y] {
					statush = append(statush, []int{x, y})
				}
			}
		}

		statush = statush[lh:]
		statusv = statusv[lv:]
	}

	return h[n-1][n-2]
}

func min(x, y int) int {
	if x == -1 && y == -1 {
		return -1
	}
	if x == -1 {
		return y
	}
	if y == -1 {
		return x
	}
	if x < y {
		return x
	}
	return y
}

// @lc code=end
