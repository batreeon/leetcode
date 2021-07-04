/*
 * @lc app=leetcode.cn id=645 lang=golang
 *
 * [645] 错误的集合
 */

// @lc code=start
func findErrorNums(nums []int) []int {
	n := len(nums)
	times := make([]int,n+1)
	for i := 0 ; i < n ; i++ {
		times[nums[i]]++
	}
	var n1,n2 int
	for i := 1 ; i <= n ; i++ {
		if times[i] == 2 {
			n1 = i
		}
		if times[i] == 0 {
			n2 = i
		}
	}
	return []int{n1,n2}
}
// @lc code=end

