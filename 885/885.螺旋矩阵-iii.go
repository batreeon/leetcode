/*
 * @lc app=leetcode.cn id=885 lang=golang
 *
 * [885] 螺旋矩阵 III
 */

// @lc code=start
package main

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	dirs := [][]int{{0,1}, {1,0}, {0, -1}, {-1, 0}}
	di := 0
	i, j := rStart, cStart
	result := [][]int{{i, j}}
	r := 1
	change := 0
	for len(result) < rows * cols {
		for k := 0; k < r; k++ {
			i, j = i + dirs[di][0], j + dirs[di][1]
			if i >= 0 && i < rows && j >= 0 && j < cols {
				result = append(result, []int{i, j})
			}
		}
		di = (di + 1)%4
		change++
		if change == 2 {
			r++
			change = 0
		}
	}
	return result
}
// @lc code=end

