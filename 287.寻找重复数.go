/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// @lc code=start
func findDuplicate(nums []int) int {
	n := len(nums)-1
	for i := 0 ; i < n+1 ; i++ {
		index := nums[i]%n
		if index == 0 {
			index = n
		}
		nums[index] += n
		if nums[index] > 2*n {
			return index
		}
	}
	return -1
}
// @lc code=end

