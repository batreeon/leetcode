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
	
	/*这个还是不太好，繁琐
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
	*/

	//下面这个，m行n列，那么共需要向下m-1次，向右n-1次
	//排列结果：(m-1+n-1)! / [(m-1)! * (n-1)!] = (m+n-2) * (m+n-3) * .. * (m-1+1) / (n-1)!
	//评论区大神：机器人一定会走m+n-2步，即从m+n-2中挑出m-1步向下走不就行了吗？即C（（m+n-2），（m-1））
	//我脑壳笨啊，透！
	//相当于用了组合式吧，组合式C(m,n)展开分子分母的因式数都等于n,因此要选择n和m-n中较小的那个
	if m == 1 || n == 1 {
		return 1
	}//若在两边，则直接返回
	d,r := m-1,n-1

	// b,s := d,r
	// if d < r {
	// 	b,s = r,d
	// }
	b := d
	if d < r {
		b = r
	}
	ans := 1
	for i,j := d+r,1 ; i > b ; i,j = i-1,j+1 {
		ans *= i
		ans /= j	//边乘边除，防止溢出
	}
	// for i := s ; i > 0 ; i-- {
	// 	ans /= i
	// }

	return ans
}
// @lc code=end

