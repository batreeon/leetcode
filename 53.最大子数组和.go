/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
    res := nums[0]
	pre := res
	if pre < 0 {
		pre = 0
	}
	for _, num := range nums[1:] {
		pre += num
		if pre > res {
			res = pre
		}
		if pre < 0 {
			pre = 0
		}
	}
	return res
}
// @lc code=end

