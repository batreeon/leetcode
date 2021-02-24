/*
 * @lc app=leetcode.cn id=832 lang=golang
 *
 * [832] 翻转图像
 */

// @lc code=start
func flipAndInvertImage(A [][]int) [][]int {
	reverse := func(x int) int {
		if x == 0 {
			return 1
		}else{
			return 0
		}
	}
	for i := 0 ; i < len(A) ; i++ {
		l := len(A[i])
		for j := 0 ; j < l/2 ; j++ {
			A[i][j],A[i][l-j-1] = reverse(A[i][l-j-1]),reverse(A[i][j])
		}
		if l%2 == 1 {
			A[i][l/2] = reverse(A[i][l/2])
		}
	}
	return A
}
// @lc code=end

