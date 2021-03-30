/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
package main
// import "math"

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	// f := make([]int,n)
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	// f[n-1] = nums[n-1]
	// f[n-2] = max(nums[n-2],nums[n-1])
	// for i := n-3 ; i >= 0 ; i-- {
	// 	f[i] = max(f[i+1],nums[i]+f[i+2])
	// }
	// return f[0]

	// 不用数组了，反正只会用到f[i-1]和f[i-2]这两项
	fi2 := nums[n-1]
	fi1 := max(nums[n-2],fi2)
	for i := n-3 ; i >= 0 ; i-- {
		f := max(fi1,nums[i]+fi2)
		fi1,fi2 = f,fi1
	}
	return fi1
}
// @lc code=end

