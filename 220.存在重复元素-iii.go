/*
 * @lc app=leetcode.cn id=220 lang=golang
 *
 * [220] 存在重复元素 III
 */

// @lc code=start
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	abs := func(x int) int {
		if x < 0 {
			return -1 * x
		}
		return x
	}
	l := len(nums)
	for i := 0 ; i < l ; i++ {
		for j := 1 ; i+j < l && j <= k ; j++ {
			if abs(nums[i]-nums[i+j]) <= t {
				return true
			}
		}
	}
	return false
}
// @lc code=end

