/*
 * @lc app=leetcode.cn id=1480 lang=golang
 *
 * [1480] 一维数组的动态和
 */

// @lc code=start
func runningSum(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	sum := make([]int,n)
	sum[0] = nums[0]
	for i := 1 ; i < n ; i++ {
		sum[i] = nums[i] + sum[i-1]
	}
	return sum
}
// @lc code=end

