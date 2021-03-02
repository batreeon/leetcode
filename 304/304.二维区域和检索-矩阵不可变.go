/*
 * @lc app=leetcode.cn id=304 lang=golang
 *
 * [304] 二维区域和检索 - 矩阵不可变
 */

// @lc code=start
package main
type NumMatrix struct {
	nM []int
	r,c int
}


func Constructor(matrix [][]int) NumMatrix {
	r := len(matrix)
	c := 0
	if r != 0 {
		c = len(matrix[0])
	}
	new := NumMatrix{nM:make([]int,r*c+1) , r:r , c:c}
	for i := 0 ; i < len(matrix) ; i++ {
		for j := 0 ; j < len(matrix[0]) ; j++ {
			if i == 0 && j == 0 {
				new.nM[1] = matrix[0][0]
			}
			if i != 0 || j != 0 {//不是[0,0]
				if i == 0 {//不是首行
					new.nM[j+1] = new.nM[j] + matrix[i][j]
				}else if j == 0 {//不是首列
					new.nM[i*c+j+1] = new.nM[(i-1)*c+j+1] + matrix[i][j]
				}else{
					new.nM[i*c+j+1] = new.nM[(i-1)*c+j+1] + new.nM[i*c+j] - new.nM[(i-1)*c+j]+ matrix[i][j]
				}
			}
		}
	}
	return new
}

/*
(row1-1,col1-1)	(row1-1,col2)
(row2,col1-1) (row2,col2) 
*/
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	r1,c1,r2,c2 := row1-1,col1-1,row2,col2
	col := this.c
	rightdown := this.nM[r2*col+c2+1]
	var leftup,rightup,leftdown int
	if row1 != 0 || col1 != 0 {
		if row1 == 0 {
			leftdown = this.nM[r2*col+c1+1]
		}else if col1 == 0 {
			rightup = this.nM[r1*col+c2+1]
		}else{
			leftdown = this.nM[r2*col+c1+1]
			rightup = this.nM[r1*col+c2+1]
			leftup = this.nM[r1*col+c1+1]
		}
	}
	return rightdown - rightup - leftdown + leftup
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end

