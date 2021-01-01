/*
 * @lc app=leetcode.cn id=941 lang=golang
 *
 * [941] 有效的山脉数组
 */

// @lc code=start
func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	if arr[0] > arr[1] || arr[len(arr)-2] < arr[len(arr)-1] {
		return false
	}
	i := 0
	for ; i < len(arr)-1 && arr[i+1] > arr[i] ; i++ {
	}
	if arr[i+1] > arr[i] {
		return false
	}
	for ; i < len(arr) - 1 && arr[i+1] < arr[i] ; i++ {
	}
	if i < len(arr)-1 {
		return false
	}
	return true
}

// @lc code=end

