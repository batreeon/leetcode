/*
 * @lc app=leetcode.cn id=879 lang=golang
 *
 * [879] 盈利计划
 */

// @lc code=start
func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	// f[i][j][k]前i项任务，使用不超过j个人，利益至少为k的方案数
	// f[i][j][k] = f[i-1][j][k]
	// 若j >= group[i-1],则还要f[i][j][k] += 发f[i-1][j-group[i-1]][max(k-profit[i-1],0)]
	// 很明显可以用滑动数组来优化
	mod := 1e9+7
	f := make([][]int,n+1)
	// 初始f[0][x][0]为1
	for i := range f {
		f[i] = make([]int,minProfit+1)
		f[i][0] = 1
	}
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 1 ; i <= len(group) ; i++ {
		f1 := make([][]int,n+1)
		// 第i项任务需要的人数，产生的利益
		g,p := group[i-1],profit[i-1]
		for j := 0 ; j <= n ; j++ {
			f1[j] = make([]int,minProfit+1)
			for k := 0 ; k <= minProfit ; k++ {
				f1[j][k] = f[j][k]
				if j >= g {
					f1[j][k] += f[j-g][max(k-p,0)]
				}
				if f1[j][k] > int(mod) {
					f1[j][k] -= int(mod)
				}
			}
		}
		copy(f,f1)
	}
	return f[n][minProfit]
}
// @lc code=end

