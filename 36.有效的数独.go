/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	type mapBB map[byte]bool
	// 记录每个列，记录每个小方形区域
	mm := make([]mapBB,9)
	mmm := make([]mapBB,9)
	for i := 0 ; i < 9 ;i++ {
		mm[i] = mapBB{}
		mmm[i] = mapBB{}
	}
	for i := 0 ; i < 9 ; i++ {
		m := make(mapBB, 0)
		for j := 0 ; j < 9 ; j++ {
			if board[i][j] >= '0' && board[i][j] <= '9' {
				if m[board[i][j]] == true {
				return false
				}
				m[board[i][j]] = true

				if mm[j][board[i][j]] == true {
					return false
				}
				mm[j][board[i][j]] = true

				if mmm[(i/3)*3+(j/3)][board[i][j]] == true {
					return false
				}
				mmm[(i/3)*3+(j/3)][board[i][j]] = true
			}
		}
	}
	return true
}
// @lc code=end

