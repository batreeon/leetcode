/*
 * @lc app=leetcode.cn id=289 lang=golang
 *
 * [289] 生命游戏
 */

// @lc code=start
package main

func gameOfLife(board [][]int)  {
    r,c := len(board), len(board[0])
	for i := range board {
		for j := range board[i] {
			neighbor := 0
			for k := -1; k <= 1; k++ {
				for z := -1; z <= 1; z++ {
					ii ,jj := i + k, j + z
					if ii >= 0 && ii < r && jj >= 0 && jj < c && !(ii == i && jj == j) {
						if board[ii][jj] == 1 || board[ii][jj] == -1 {
							neighbor++
						}
					}
				}
			}

			if board[i][j] == 1 {
				if neighbor < 2 || neighbor > 3 {
					board[i][j] = -1
				}
			}else{
				if neighbor == 3 {
					board[i][j] = 2
				} 
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] > 0 {
				board[i][j] = 1
			}else{
				board[i][j] = 0
			}
		}
	}
}
// @lc code=end

