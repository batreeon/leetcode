/*
 * @lc app=leetcode.cn id=456 lang=golang
 *
 * [456] 132模式
 */

// @lc code=start
package main
import "math"


func find132pattern(nums []int) bool {
	l := len(nums)
	if l < 3 {
		return false
	}


	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}

	leftMin := nums[0]
	findMaxbelowThreshold := func(nums []int,threshold,start,end int) int {
		ans := math.MinInt64
		for i := start ; i <= end ; i++ {
			if nums[i] < threshold {
				ans = max(ans,nums[i])
			}
		}
		return ans
	}
	rightMax := math.MinInt64
	for i := 1 ; i < l-1 ; i++ {
		leftMin = min(leftMin,nums[i-1])
		if nums[i] <= leftMin {
			continue
		}
		
		if nums[i-1] >= nums[i] || rightMax != math.MaxInt64 {
			rightMax = findMaxbelowThreshold(nums,nums[i],i+1,l-1)
		}
		
		if nums[i] > rightMax && rightMax > leftMin {
			return true
		}
	}
	return false
}
// @lc code=end

