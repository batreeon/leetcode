/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */

// @lc code=start
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	n := len(nums)
	if nums[n-1] > nums[n-2] {
		return n-1
	}
	for i := 1 ; i < n-1 ; {
		if nums[i] > nums[i+1] {
			if nums[i] > nums[i-1] {
				return i
			}
			i = i+2
		}else{
			i = i+1
		}
	}
	return -1
}
// @lc code=end

