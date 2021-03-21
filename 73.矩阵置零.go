/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int)  {
	r := len(matrix)
	if r == 0 {
		return
	}
	c := len(matrix[0])
	row := make([]int,c)
	col := make([]int,r)

	for i := 0 ; i < c ; i++ {
		row[i] = 1
	}
	for i := 0 ; i < r ; i++ {
		col[i] = 1
	}

	for i,rr := range matrix {
		for j,k := range rr {
			if k == 0 {
				row[j] = 0
				col[i] = 0
			}
		}
	}

	for i,rr := range matrix {
		for j,_ := range rr {
			matrix[i][j] = matrix[i][j] * row[j] * col[i]
		}
	}
}
// @lc code=end

