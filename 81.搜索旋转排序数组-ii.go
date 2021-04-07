/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start
package main
// import "sort"

func search(nums []int, target int) bool {

	/*
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		if nums[0] == target {
			return true
		}
		return false
	}
	l,r := 0,n-1
	for l <= r {
		mid := (l+r)/2
		if nums[mid] == target {
			return true
		}
		// 1,0,1,1,1,不能直接用nums[l] <= nums[mid]是否有序，
		// 因为当l==0,mid==2时，即1,0,1序列，他也判定为有序
		// 这是有重复元素带来的问题
		if sort.IsSorted(sort.IntSlice(nums[l:mid+1])) {
			if nums[l] <= target && target < nums[mid]{
				r = mid-1
			}else{
				l = mid+1
			}
		}else{
			if nums[mid] < target && target <= nums[r] {
				l = mid+1
			}else{
				r = mid-1
			}
		}
	}
	return false
	*/

	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return nums[0] == target
	}
	l,r := 0,n-1
	for l < r {
		mid := (l+r)/2
		if nums[mid] == target {
			return true
		}

		// 有重复元素，可能出现无法根据nums[l],nums[mid],nums[r]判断哪边有序的情况
		if nums[l] == nums[mid] && nums[r] == nums[mid] {
			l++
			r--
		}else if nums[mid] >= nums[l] {
			if nums[l] <= target && target < nums[mid]{
				r = mid-1
			}else{
				l = mid+1
			}
		}else{
			if nums[mid] < target && target <= nums[r] {
				l = mid+1
			}else{
				r = mid-1
			}
		}
	}
	return nums[l] == target
}
// @lc code=end

