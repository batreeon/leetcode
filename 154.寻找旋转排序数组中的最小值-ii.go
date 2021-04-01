/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

// @lc code=start
package main
import "sort"

func findMin(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	l,r := 0,n-1
	for l < r {
		if sort.IsSorted(sort.IntSlice(nums[l:r+1])) {
			return nums[l]
		}
		mid := (l+r)/2
		// 中间元素不大于左边界，并且左侧不是有序的，若右侧是有序的，则一定一串相同的，
		// 因为nums[0]必定不是最小的，因此排除左侧，转向右侧
		// 若中间元素不大于左边界，且左侧不是有序的，那么最小值一定在左半侧
		if nums[mid] <= nums[l] && !sort.IsSorted(sort.IntSlice(nums[l:mid+1])) {
		// if nums[mid] <= nums[l] && nums[mid] <= nums[r] {	3313
		// if nums[mid] <= nums[l] {	22201
		// if nums[mid] <= nums[l] && sort.IsSorted(sort.IntSlice(nums[mid+1:r+1])) {	22201
			r = mid
		}else{
			l = mid+1
		}
	}
	return nums[l]
	// 2201
	// 3313
}
// @lc code=end

