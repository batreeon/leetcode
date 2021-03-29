/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start
func solve(board [][]byte)  {
	// 策略，创建一个与原数组大小相同的全是X的数据，然后填充没被包围的区域，从边界逐步生长
	// 因为被包围的无法从边界生长到，所以就不用动它
	r,c := len(board),len(board[0])
	res := make([][]byte,r)
	allx := make([]byte,c)
	for i := 0 ; i < c ; i++ {
		allx[i] = byte('X')
	}
	for i := 0 ; i < r ; i++ {
		res[i] = make([]byte,len(board[i]))
		copy(res[i],allx)
	}
	oris := make([][]int,4)
	oris[0] = []int{-1,0}
	oris[1] = []int{0,-1}
	oris[2] = []int{0,1}
	oris[3] = []int{1,0}
	var dfs func(i,j int)
	dfs = func(i,j int) {
		res[i][j] = byte('O')
		for _,ori := range oris {
			x,y := i+ori[0],j+ori[1]
			if x >= 0 && x < r && y >= 0 && y < c && res[x][y] != byte('O') && board[x][y] == byte('O') {
				dfs(x,y)
			}
		}
	}
	for i := 0 ; i < r ; i++ {
		for j := 0 ; j < c ; j++ {
			if board[i][j] == byte('O') && (i == 0 || i == r-1 || j == 0 || j == c-1) && res[i][j] != byte('O') {
				dfs(i,j)
			}
		}
	}
	copy(board,res)
}
// @lc code=end

