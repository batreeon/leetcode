/*
 * @lc app=leetcode.cn id=456 lang=golang
 *
 * [456] 132模式
 */

// @lc code=start
package main
import "math"

func find132pattern(nums []int) bool {
	// 逻辑：从下下标1到n-1遍历，这是指示中间元素。
	// 找他左侧最小的元素，找他右侧比他小的最大的元素，如果符合132模式，就返回true
	// 若为false,就继续遍历
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
		
		// 第一次执行，即rightMax == math.MaxInt64时
		// nums[i-1]>=nums[i]是为了判断，rightMax是否变化，没有变化，就不需要再调用findMaxbelowThreshold函数了
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

