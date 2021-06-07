/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	var f func(nums1 []int,target1 int) int
	f = func(nums1 []int,target1 int) int {
		if len(nums1) == 1 {
			// 下面两个分支顺序不能调换，第二个会覆盖第一个
			if nums1[0] == target1 && nums1[0]*(-1) == target1 {
				// 这个就是nums[0] = 0 ,target = 0的情况
				return 2
			}
			if nums1[0] == target1 || nums1[0]*(-1) == target1 {
				return 1
			}
			return 0
		}
		n := len(nums1)
		return f(nums1[:n-1],target1+nums1[n-1]) +  f(nums1[:n-1],target1-nums1[n-1])
	}
	return f(nums,target)
}
// @lc code=end

