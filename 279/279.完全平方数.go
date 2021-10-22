/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
package main
import "math"

func numSquares(n int) int {
	result := 0
	var backtrack func(n int) bool
	backtrack = func(n int) bool {
		if n == 0 {
			return true
		}
		for i := n ; i >= 0 ; i-- {
			sqrt := math.Sqrt(float64(i))
			if int(sqrt)*int(sqrt) == int(sqrt*sqrt) {
				result++
				if backtrack(n-i) {
					return true
				}
				result--
			}
		}
		return false
	}

	backtrack(n)
	return result
}
// @lc code=end

