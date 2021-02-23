/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续1的个数
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	left := -1
	max := 0
	for right := 0 ; right < len(nums) ; right++ {
		if nums[right] == 0 {
			left = right
		}
		if right-left > max {
			max = right - left
		}
	}
	return max
	/*
	left := -1
	max := 0
	right := 0
	for  ; right < len(nums) ; right++ {
		if nums[right] == 0 {
			if right-left > max {
				max = right - left - 1
			}
			left = right
		}
	}
	if right-left > max {
		max = right - left - 1
	}
	return max
	*/
}
// @lc code=end

