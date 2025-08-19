/*
 * @lc app=leetcode.cn id=2348 lang=golang
 *
 * [2348] 全 0 子数组的数目
 */

// @lc code=start
func zeroFilledSubarray(nums []int) int64 {
	var result int64
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroLen := 1
			i++
			for ; i < len(nums) && nums[i] == 0; i++ {
				zeroLen++
			}
			result += int64(zeroLen * (zeroLen + 1) / 2)
		}
	}
	return result
}
// @lc code=end

