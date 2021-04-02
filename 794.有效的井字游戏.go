/*
 * @lc app=leetcode.cn id=794 lang=golang
 *
 * [794] 有效的井字游戏
 */

// @lc code=start
package main
func validTicTacToe(board []string) bool {
	numO,numX,sumNum := 0,0,0
	for _,row := range board {
		for _,b := range row {
			if b == 'X' {
				numX++
				sumNum++
			}else if b == 'O' {
				numO++
				sumNum++
			}
		}
	}
	if numO > numX {
		return false
	}
	if numX-1 > numO {
		return false
	}
	isOver := func(pattern string,r byte) bool {
		numX := make([]int,3)
		for _,row := range board {
			if row == pattern {
				return true
			}
			for i,b := range row {
				if b == rune(r) {
					numX[i]++
					if numX[i] == 3 {
						return true
					}
				}
			}
		}
		if board[1][1] == r {
			if board[0][0] == r && board[2][2] == r {
				return true
			}
			if board[0][2] == r && board[2][0] == r {
				return true
			}
		}
		return false
	}
	isOverX,isOverO := isOver("XXX",byte('X')) , isOver("OOO",byte('O'))
	if isOverX && isOverO {
		return false
	}else if isOverX && !isOverO {
		if numO+1 != numX {
			return false
		}
	}else if !isOverX && isOverO {
		if numO != numX {
			return false
		}
	}

	// 搞不懂，这题，我认为题目的意思是，
	// 给出的棋盘必须是结束状态，原来不是，
	// 那么就不需要下面这句了
	// }else{
	
		// if sumNum != 9 {
		// 	return false
		// }
	// }
	return true
}
// @lc code=end

