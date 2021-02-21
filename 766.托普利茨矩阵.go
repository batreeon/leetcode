/*
 * @lc app=leetcode.cn id=766 lang=golang
 *
 * [766] 托普利茨矩阵
 */

// @lc code=start
func isToeplitzMatrix(matrix [][]int) bool {
	//3*4
	//20
	//10 21
	//00 11 22
	//01 12 13
	//02 13
	//03
	if len(matrix) == 0 {
		return true
	}
	r,c := len(matrix),len(matrix[0])
	temp := make([]int,c)
	temp = matrix[0][:]
	for i := 1 ; i < r ; i++ {
		for j := 1 ; j < c ; j++ {
			if matrix[i][j] != temp[j-1] {
				return false
			}
		}
		temp = matrix[i][:]
	}
	return true
}
// @lc code=end

