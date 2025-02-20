/*
 * @lc app=leetcode.cn id=1299 lang=golang
 *
 * [1299] 将每个元素替换为右侧最大元素
 */

// @lc code=start
func replaceElements(arr []int) []int {
    n := len(arr)
	max := arr[n-1]
	for i := n - 2; i >= 0; i-- {
		arri := arr[i]
		arr[i] = max
		if arri > max {
			max = arri
		}
	}
	arr[n-1] = -1
	return arr
}
// @lc code=end

