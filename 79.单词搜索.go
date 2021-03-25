/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	// 速度太慢，暴力解
	seen := make(map[int]bool)
	r,c := len(board),len(board[0])
	ori := [][]int{}
	ori = append(ori,[]int{-1,0})
	ori = append(ori,[]int{0,-1})
	ori = append(ori,[]int{0,1})
	ori = append(ori,[]int{1,0})
	var backtrack func(i,j int,word string) bool
	backtrack = func(i,j int,word string) bool {
		if len(word) == 0 {
			return true
		}
		for k := 0 ; k < 4 ; k++ {
			newi,newj := i+ori[k][0],j+ori[k][1]
			if newi >= 0 && newi < r && newj >=0 && newj < c {
				if board[newi][newj] == word[0] && !seen[newi*c+newj] {
					seen[newi*c+newj] = true
					if backtrack(newi,newj,word[1:]) {
						return true
					}
					delete(seen,newi*c+newj)
				}
			}	
		}
		return false
	}

	for i := 0 ; i < r ; i++ {
		for j := 0 ; j < c ; j++ {
			if board[i][j] == word[0] {
				seen[i*c+j] = true
				if backtrack(i,j,word[1:]) {
					return true
				}
				delete(seen,i*c+j)
			}
		}
	}
	return false
}
// @lc code=end

