/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
    r,c := len(matrix)-1,len(matrix[0])-1
	for i := 0 ; i <= r ; i++ {
		for j := c ; j >= 0 ; j-- {
			if matrix[i][j] == target {
				return true
			}
			if matrix[i][j] > target {
				// 若matrix[i][j] > target那么matrix[i+1][j]必大于target
				c = j-1
			}
			if matrix[i][j] < target {
				break
			}
		}
	}
	return false
}
// @lc code=end

