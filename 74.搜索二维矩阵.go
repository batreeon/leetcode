/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	r,c := len(matrix),len(matrix[0])
	t,d := 0,r-1
	// 找到末尾元素大于target的最小值所在的行
	for t < d {
		mid := (t+d)/2
		if matrix[mid][c-1] == target {
			return true
		}
		if matrix[mid][c-1] > target {
			d = mid
		}else{
			t = mid+1
		}
	}
	for _,v := range matrix[t] {
		if v == target {
			return true
		}
	}
	return false
}
// @lc code=end

