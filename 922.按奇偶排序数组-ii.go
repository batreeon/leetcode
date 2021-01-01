/*
 * @lc app=leetcode.cn id=922 lang=golang
 *
 * [922] 按奇偶排序数组 II
 */

// @lc code=start
func sortArrayByParityII(A []int) []int {
	for i := 0; i < len(A);i++ {
		if A[i] % 2 != i % 2 {
			j := i + 1
			for ; A[j] % 2 != i % 2;j++ {
			}
			A[i],A[j] = A[j],A[i]
		}
	}
	return A
}
// @lc code=end

