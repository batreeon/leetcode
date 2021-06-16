/*
 * @lc app=leetcode.cn id=877 lang=golang
 *
 * [877] 石子游戏
 */

// @lc code=start
package main
func stoneGame(piles []int) bool {
	
	// 每次取最大？？想多了	3,2,10,4
	// 每次都取差值最大的？不行
	// 反向，从最大的开始依次填充？？错误
	
	// 最后还是看答案了，动态规划
	n := len(piles)
	f := make([][]int,n)
	// f[i][j]表示区间[i,j]，先手和后手的最大差值

	for i := range piles {
		f[i] = make([]int,n)
		f[i][i] = piles[i]
	}
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 0 ; i < n-1 ; i++ {
		for j := i+1 ; j < n ; j++ {
			f[i][j] = max(piles[i]-f[i+1][j],piles[j]-f[i][j-1])
		}
	}
	return f[0][n-1] > 0
}
// @lc code=end

