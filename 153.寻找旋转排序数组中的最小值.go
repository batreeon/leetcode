/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	l,r := 0,n-1
	for l < r {
		if nums[l] < nums[r] {
			// 若有序，则一定包含最小值
			return nums[l]
		}
		if nums[l] > nums[r] {
			mid := (l+r)/2
			if nums[mid] < nums[r] {
				// 不能排除mid是否为最小
				r = mid
			}else if nums[mid] > nums[r] {
				// 可以排除mid为最小
				l = mid+1
			}
		}
	}
	return nums[l]
}
// @lc code=end

