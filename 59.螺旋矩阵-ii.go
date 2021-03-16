/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	r,c := n,n
	start := 0
	num := 1
	ans := [][]int{}
	ans = append(ans,make([]int,n))

	// spiralOrderInCircle := func() {
	// 	endX,endY := c-start-1,r-start-1
	// 	//从左到右打印一行
	// 	for j := start ; j <= endX ; j++ {
	// 		ans[start][j] = num
	// 		num++
	// 	}
	// 	if start < endY {//从上到下，若只有一行，则不打印
	// 		for i := start+1 ; i <= endY ; i++ {
	// 			if endX == n-1 {
	// 				ans = append(ans,make([]int,n))
	// 			}
	// 			ans[i][endX] = num
	// 			num++
	// 		}
	// 	}
	// 	if start < endX && start < endY {//从右向左，若只有一行或一列，不打印
	// 		for j := endX-1 ; j >= start ; j-- {
	// 			ans[endY][j] = num
	// 			num++
	// 		}
	// 	}
	// 	if start < endX && start < endY - 1 {//从下到上，若不超过两行或者只有一列，不打印
	// 		for i := endY-1 ; i > start ; i-- { //这里循环条件就没有=了
	// 			ans[i][start] = num
	// 			num++
	// 		}
	// 	}
	// }

	// 每一圈的左上角起点坐标：(start,start)
	for r > start*2 && c > start*2 {
		func() {
			endX,endY := c-start-1,r-start-1
			//从左到右打印一行
			for j := start ; j <= endX ; j++ {
				ans[start][j] = num
				num++
			}

			// 至少有两行
			if start < endY {//从上到下，若只有一行，则不打印
				for i := start+1 ; i <= endY ; i++ {
					if endX == n-1 {
						ans = append(ans,make([]int,n))
					}//第一次进行向下打印时，初始化每行，这一步也可以单独在定义ans时进行吧，那样还简单易懂些
					ans[i][endX] = num
					num++
				}
			}

			// 至少有两行两列
			if start < endX && start < endY {//从右向左，若只有一行或一列，不打印
				for j := endX-1 ; j >= start ; j-- {
					ans[endY][j] = num
					num++
				}
			}

			// 至少有三行两列
			if start < endX && start < endY - 1 {//从下到上，若不超过两行或者只有一列，不打印
				for i := endY-1 ; i > start ; i-- { //这里循环条件就没有=了
					ans[i][start] = num
					num++
				}
			}
		}()
		
		start++
	}
	return ans
}
// @lc code=end

