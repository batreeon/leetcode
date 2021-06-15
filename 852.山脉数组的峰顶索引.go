/*
 * @lc app=leetcode.cn id=852 lang=golang
 *
 * [852] 山脉数组的峰顶索引
 */

// @lc code=start
func peakIndexInMountainArray(arr []int) int {
	n := len(arr)
	l,r := 0,n-1
	for l < r {
		mid := (l+r)/2
		if mid != 0 && arr[mid] > arr[mid-1] {
			l = mid
		}
		if mid != n-1 && arr[mid] > arr[mid+1] {
			r = mid
		}
	}
	return l
}
// @lc code=end

