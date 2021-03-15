/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	//每一圈的起始左上角坐标是(0,0),(1,1)...(x,x)
	//终止条件，r>x*2 && c>x*2
	
	if len(matrix) == 0 {
		return []int{}
	}
	r,c := len(matrix),len(matrix[0])
	start := 0
	ans := []int{}
	spiralOrderInCircle := func() {
		endX,endY := c-start-1,r-start-1
		//从左到右打印一行
		for j := start ; j <= endX ; j++ {
			ans = append(ans,matrix[start][j])
		}

		// 至少有两行，才会向下打印
		if start < endY {//从上到下，若只有一行，则不打印
			for i := start+1 ; i <= endY ; i++ {
				ans = append(ans,matrix[i][endX])
			}
		}

		// 至少有两行两列，才会向左打印
		if start < endX && start < endY {//从右向左，若只有一行或一列，不打印
			for j := endX-1 ; j >= start ; j-- {
				ans = append(ans,matrix[endY][j])
			}
		}

		// 至少有三行两列，才会向上打印
		if start < endX && start < endY - 1 {//从下到上，若不超过两行或者只有一列，不打印
			for i := endY-1 ; i > start ; i-- { //这里循环条件就没有=了
				ans = append(ans,matrix[i][start])
			}
		}
	}
	for r > start*2 && c > start*2 {
		spiralOrderInCircle()
		start++
	}
	return ans
}
// @lc code=end

