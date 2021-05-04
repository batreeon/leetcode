/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start
package main

import "strings"

func solveNQueens(n int) [][]string {
	// 回溯

	// 用来保存最终返回的结果
	result := [][]string{}
	var solvequeens func(res *[]string,row int) 
	solvequeens = func(res *[]string,row int) {
		// 下面要在第row行上填充皇后
		// 考察每一个位置
		for i := 0 ; i < n ; i++ {
			flag := true
			for j := 0 ; j < row ; j++ {
				// 同一列的上层有皇后了
				if (*res)[j][i] == 'Q' {
					flag = false
					break
				}

				y1,y2 := i+(row-j),i-(row-j)
				// 上层斜向右上有皇后了
				if y1 >= 0  && y1 < n {
					if (*res)[j][y1] == 'Q' {
						flag = false
						break
					}
				}
				// 上层斜向左上有皇后了
				if y2 >= 0 && y2 < n {
					if (*res)[j][y2] == 'Q' {
						flag = false
						break
					}
				}
			}
			// row层i列可以放皇后
			if flag {
				// 生成这一行
				temp := []string{}
				for k := 0 ; k < n ; k++ {
					temp = append(temp,".")
				}
				temp[i] = "Q"
				// 保存到中间结果中
				(*res) = append(*res,strings.Join(temp,""))
				if row == n-1 {
					// 若是最后一行，则将结果作为一种解保存到result中
					rescopy := make([]string,n)
					copy(rescopy,*res)
					result = append(result,rescopy)
				}else{
					// 若不是最后一行，则继续深入下一行
					solvequeens(res,row+1)
				}
				// 回退
				(*res) = (*res)[:len(*res)-1]
			}
		}
	}

	solvequeens(&([]string{}),0)

	return result
}
// @lc code=end

