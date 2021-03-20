/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */

// @lc code=start
package main
func maximalSquare(matrix [][]byte) int {
	r,c := len(matrix),len(matrix[0])
	lastRow := make([]int,c)
	maxSide := 0
	//写了三段这样的，都是为了应对答案为1的情况
	for i,v := range matrix[0] {
		if v == byte('1') {
			lastRow[i] = 1
			if maxSide == 0 {
				maxSide = 1
			}
		}
	}

	minAll := func(x,y,z int) int {
		if y < x {
			x,y = y,x
		}
		if z < x {
			x,z = z,x
		}
		return x
	}
	
	for i := 1 ; i < r ; i++ {
		curRow := make([]int,c)
		curRow[0] = int(matrix[i][0]-byte('0'))
		//
		if curRow[0] > maxSide {
			maxSide = curRow[0]
		}
		for j := 1 ; j < c ; j++ {
			if matrix[i][j] == byte('1') && minAll(lastRow[j-1],lastRow[j],curRow[j-1]) != 0 { 
				//
				if maxSide == 0 {
					maxSide = 1
				}
				curRow[j] = minAll(lastRow[j-1],lastRow[j],curRow[j-1])+1
				if curRow[j] > maxSide {
					maxSide = curRow[j]
				}
			}else{
				curRow[j] = int(matrix[i][j]-byte('0'))
				if curRow[j] > maxSide {
					maxSide = curRow[j]
				}
			}
		}
		copy(lastRow,curRow)
	}
	return maxSide*maxSide
}
// @lc code=end

