/*
 * @lc app=leetcode.cn id=1928 lang=golang
 *
 * [1928] 规定时间内到达终点的最小花费
 */

// @lc code=start
// 与787相似，只不过这里的约束由步数变成了时间
var INF = 1000*1000+1
func minCost(maxTime int, edges [][]int, passingFees []int) int {
	n := len(passingFees)
	// f[t][i]表示耗费t时间到达i的花费
	f := make([][]int,maxTime+1)
	for i := range f {
		f[i] = make([]int,n)
		for j := range f[i] {
			f[i][j] = INF
		}
	}
	// 注意这一点，起点也要付通行费
	f[0][0] = passingFees[0]

	for t := 1 ; t <= maxTime ; t++ {
		for _,edge := range edges {
			n1,n2,time := edge[0],edge[1],edge[2]
			if time <= t {
				// 双向道路
				f[t][n1] = min(f[t][n1],f[t-time][n2]+passingFees[n1])
				f[t][n2] = min(f[t][n2],f[t-time][n1]+passingFees[n2])
			}
		}
	}

	result := INF
	for t := 1 ; t <= maxTime ; t++ {
		result = min(result,f[t][n-1])
	}
	if result == INF {
		return -1
	}
	return result
}
func min(x,y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

