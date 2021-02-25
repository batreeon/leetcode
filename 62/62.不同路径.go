/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
package main
func uniquePaths(m int, n int) int {
	/*
	错误做法，path是共享的，数值多次计入，算的值大于正确值
	path := make([]int,m*n)
	path[0] = 1
	var nextPos func(int,int)
	nextPos = func(posi,posj int) {
		if posi + 1 < m  {
			path[(posi+1)*n+posj] += path[posi*n+posj]
			nextPos(posi+1,posj)
		}
		if posj + 1 < n {
			path[posi*n+(posj+1)] += path[posi*n+posj]
			nextPos(posi,posj+1)
		}
	}
	nextPos(0,0)
	return path[m*n-1]
	*/
	
	/*
	//m,n太大时居然超时了，经过测试结果是对的，但是耗时太长
	var nextPos func(int,int,int) int
	nextPos = func(paths int,posi,posj int) int {
		if posi == m-1 && posj == n-1 {
			return paths
		}
		res := 0
		if posi + 1 < m  {
			res += nextPos(paths,posi+1,posj)
		}
		if posj + 1 < n {
			res += nextPos(paths,posi,posj+1)
		}
		return res
	}
	return nextPos(1,0,0)
	*/
	
	//用纸写了下，看出来了，就是经过旋转的杨辉三角
	//第m行n列是杨辉三角中第m+n-1行，第m+n-1-(m-1) = n个元素
	if m == 1 || n == 1 {
		return 1
	}//若在两边，则直接返回

	rowIndex := m+n-1//从第一行开始计,x行有x个元素。
	row := make([]int, rowIndex)
    row[0] = 1
    for i := 1; i <= rowIndex/2; i++ {
        row[i] = row[i-1] * (rowIndex - i) / i
    }//每一行对称的，只计算一半

	index := n-1//n-1是下标
	if n - 1 > rowIndex/2 {
		index = rowIndex-n //写成index = rowIndex - index - 1这个比较常见好理解一些
	}
	return row[index]
}
// @lc code=end

