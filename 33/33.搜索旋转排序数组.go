/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
package main
func search(nums []int, target int) int {
	//二分查找
	//判断哪边有序，哪边有序就在哪边找
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	l,r := 0,n-1
	for l <= r {
		mid := (l+r)/2
		if nums[mid] == target {
			return mid
		}
		if nums[l] <= nums[mid] {//有序
			if nums[l] <= target && target < nums[mid] {//判断是否在[l,mid)区间里
				r = mid-1
			}else{
				l = mid + 1
			}
		}else{
			if nums[mid] < target && target <= nums[r] {//判断是否在(mid,r]区间里
				l = mid + 1
			}else{
				r = mid - 1
			}
		}
	}
	return -1
}
// @lc code=end

