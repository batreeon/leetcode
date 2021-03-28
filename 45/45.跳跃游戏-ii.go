/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
package main
import "math"

func jump(nums []int) int {
	n := len(nums)
	f := make([]int,n)
	f[n-1] = 0
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	for i := n-2 ; i >= 0 ; i-- {
		// 只能跳0，不可能到终点
		if nums[i] == 0 {
			continue
		}

		// 只需一步就可以到终点
		if nums[i]+i >= n-1 {
			f[i] = 1
			continue
		}

		// 从后到前，考察所需的最小步数
		minStep := math.MaxInt64	
		for j := nums[i]+i ; j > i ; j-- {
			if f[j] == 0 {
				continue
			}
			// f[j]>0,则最终minStep>1
			// 若minStep==math.MaxInt64,则没有执行过下面这一句，不能到达终点，不需要更新f[i]
			minStep = min(minStep,1+f[j])
		}
		// minStep有两种可能，MaxInt64（f[i]应等于0）,大于1（f[i]应等于minStep）
		if minStep != math.MaxInt64 {
			f[i] = minStep
		}
	}
	return f[0]
}
// @lc code=end

