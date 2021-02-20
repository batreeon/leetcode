/*
 * @lc app=leetcode.cn id=1438 lang=golang
 *
 * [1438] 绝对差不超过限制的最长连续子数组
 */

// @lc code=start
package main
import "math"
func longestSubarray(nums []int, limit int) int {
	left := 0
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	findMinV := func(nums []int) int {
		minV := math.MaxInt64
		for _,v := range nums {
			if v < minV {
				minV = v
			}
		}
		return minV
	}
	findMaxV := func(nums []int) int {
		maxV := math.MinInt64
		for _,v := range nums {
			if v > maxV {
				maxV = v
			}
		}
		return maxV
	}
	ans := 0
	maxV := math.MinInt64
	minV := math.MaxInt64
	for right := 0 ; right < len(nums) ; right++ {
		minV = min(nums[right],minV)
		maxV = max(nums[right],maxV)
		if maxV - minV > limit {//收缩窗口
			if nums[left] == minV {
				minV = findMinV(nums[left+1:right+1])
			}
			if nums[left] == maxV {
				maxV = findMaxV(nums[left+1:right+1])
			}
			left++
		}
		ans = max(ans,right-left+1)
	}
	return ans
}
// @lc code=end

