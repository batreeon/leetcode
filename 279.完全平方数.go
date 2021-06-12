/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
package main
import "math"
func numSquares(n int) int {
	f := make([]int,n+1)
	for i := 1 ; i <= n ; i++ {
		minV := math.MaxInt64
		for j := 1 ; j * j <= i ; j++ {
			if minV > f[i-j*j] {
				minV = f[i-j*j]
			}
		}
		f[i] = minV+1
		// 为什么加一，因为是i-j*j + j*j，加了一个完全平方数j*j
	}
	return f[n]
}
// @lc code=end

