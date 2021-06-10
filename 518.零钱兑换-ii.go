/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
package main
import "sort"

func change(amount int, coins []int) int {
	f := make([]int,amount+1)
	sort.Ints(coins)
	f[0] = 1
	for _,coin := range coins {
		f1 := make([]int,amount+1)
		copy(f1,f)
		for i := 1 ; i <= amount ; i++ { 
			for k := 1 ; i-coin*k >= 0 ; k++ {
				f1[i] += f[i-coin*k]
			}
		}
		copy(f,f1)
	}
	return f[amount]
}
// @lc code=end

