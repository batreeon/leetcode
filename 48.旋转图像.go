/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start
func rotate(matrix [][]int)  {
	n := len(matrix)
	for i := 0 ; i < n/2; i++{
		for j := i ; j < n-i-1 ; j++ {
			//m1,m2,m3,m4 := matrix[i][j],matrix[j][n-i-1],matrix[n-i-1][n-j-1],matrix[n-j-1][0]
			matrix[i][j],matrix[j][n-i-1],matrix[n-i-1][n-j-1],matrix[n-j-1][i] = matrix[n-j-1][i],matrix[i][j],matrix[j][n-i-1],matrix[n-i-1][n-j-1]
		}
	}
}
// @lc code=end

