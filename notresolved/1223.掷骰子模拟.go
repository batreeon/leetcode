/*
 * @lc app=leetcode.cn id=1223 lang=golang
 *
 * [1223] 掷骰子模拟
 */

// @lc code=start
package main

const mo int =  1<<9 + 7

func dieSimulator(n int, rollMax []int) int {
	result := f(6, n)
	rollMaxM := map[int]int{}
	for _, rm := range rollMax {
		if rm < n {
			rollMaxM[rm]++
		}
	}

	
}

func f(x, n int) int {
	if n == 0 {
		if x > mo {
			return x%mo
		}
		return x%mo
	}
	result := x * f(x, n-1)
	if result > mo {
		result %= mo
	}
	return result
}
// @lc code=end

