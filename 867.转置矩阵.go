/*
 * @lc app=leetcode.cn id=867 lang=golang
 *
 * [867] 转置矩阵
 */

// @lc code=start
func transpose(matrix [][]int) [][]int {
	// if len(matrix) == 0 {
	// 	return [][]int{}
	// }
	r,c := len(matrix),len(matrix[0])
	ans := [][]int{}
	for j := 0 ; j < c ; j++ {
		ans = append(ans,[]int{})
		for i := 0 ; i < r ; i++ {
			ans[j] = append(ans[j],matrix[i][j])
		}
	}//空间局部性有点差吧
	return ans
}
// @lc code=end

