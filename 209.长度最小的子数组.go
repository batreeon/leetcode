/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	result := n+1
	left,right := 0,-1
	sum := 0
	loop:
	for {
		for sum < target {
			right++
			if right == n {
				break loop
			}
			sum += nums[right]
		}
		if right-left+1 < result {
			result = right-left+1
		}
		sum -= nums[left]
		left++
	}
	if result == n+1 {
		result = 0
	}
	return result
}
// @lc code=end

