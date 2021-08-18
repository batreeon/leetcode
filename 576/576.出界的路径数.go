/*
 * @lc app=leetcode.cn id=576 lang=golang
 *
 * [576] 出界的路径数
 */

// @lc code=start
package main
func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	/*
	grid := make([][][]int,m)
	for i := range grid {
		grid[i] = make([][]int,n)
		for j := range grid[i] {
			grid[i][j] = make([]int,maxMove)
		}
	}
	orix,oriy := []int{-1,0,0,1},[]int{0,-1,1,0}
	mod := 1000000007
	var bfs func(int,int,*[][][]int,int) int
	bfs = func(x,y int,grid *[][][]int,move int) int {
		if move == 0 {
			if x < 0 || x >= m || y < 0 || y >= n {
				return 1
			}else{
				return 0
			}
		}
		if x < 0 || x >= m || y < 0 || y >= n {
			return 0
		}
		if (*grid)[x][y][move-1] > 0 {
			return (*grid)[x][y][move-1]
		}
		res := 0
		for i := range orix {
			newX,newY := x+orix[i],y+oriy[i]
			res += bfs(newX,newY,grid,move-1)
			if res > mod {
				res %= mod
			}
		}
		(*grid)[x][y][move-1] = res
		return res
	}
	result := 0
	for i := 1 ; i <= maxMove ; i++ {
		result += bfs(startRow,startColumn,&grid,i)
	}
	if result > mod {
		result %= mod
	}
	return result
	*/

	mod := 1000000007
	orix,oriy := []int{-1,0,0,1},[]int{0,-1,1,0}
	// dp[i][j]表示的是从起点到达i,j的路径数
	dp := make([][]int,m)
	for i := range dp {
		dp[i] = make([]int,n)
	}
	// 步数为0时，只有到达起点的路径数为1，其他的都为0
	dp[startRow][startColumn] = 1
	// 出界的路径数
	result := 0
	for i := 0 ; i < maxMove ; i++ {
		// newDp[newX][newY]用来记录第i+1步到达newX,newY的路径数
		newDp := make([][]int,m)
		for i := range newDp {
			newDp[i] = make([]int,n)
		}
		for j := 0 ; j < m ; j++ {
			for k := 0 ; k < n ; k++ {
				count := dp[j][k]
				// count等于0说明第i步压根没到达这里
				if count > 0 {
					for ii := range orix {
						// 朝四个方向走
						newX,newY := j+orix[ii],k+oriy[ii]
						if newX < 0 || newX >= m || newY < 0 || newY >= n {
							// 出界了，加到result上
							result += count
							result %= mod
						}else{
							// 没出界，更新newDp[newX][newY]
							newDp[newX][newY] = (newDp[newX][newY] + count) % mod
						}
					}
				}
			}
		}
		// 这个就不解释了
		dp = newDp
	}
	return result
}
// @lc code=end

