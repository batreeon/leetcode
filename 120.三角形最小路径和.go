/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	l := len(triangle)
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	for i := l - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] = min(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
		}
	}
	return triangle[0][0]
	// 怎么用了这么多时间
}

// @lc code=end

