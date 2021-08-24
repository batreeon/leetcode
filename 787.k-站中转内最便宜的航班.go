/*
 * @lc app=leetcode.cn id=787 lang=golang
 *
 * [787] K 站中转内最便宜的航班
 */

// @lc code=start
// 与1928相似
package main

// 由题条件，价格不会超过INF
var INF = 10000*100+1

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// f[t][i]表示经过t条航班，从src到i的价格，因为最多能经过k站中转，所以最多可经过k+1条航班
	f := make([][]int,k+2)
	for i := range f {
		f[i] = make([]int,n)
		for j := range f[i] {
			f[i][j] =  INF
		}
	}
	f[0][src] = 0

	for t := 1 ; t <= k+1 ; t++ {
		for _,flight := range flights {
			f[t][flight[1]] = min(f[t][flight[1]],f[t-1][flight[0]]+flight[2])
		}
	}
	result := INF
	for t := 1 ; t <= k+1 ; t++ {
		result = min(result,f[t][dst])
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

