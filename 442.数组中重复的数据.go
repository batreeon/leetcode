/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

// @lc code=start
func findDuplicates(nums []int) []int {
	// 长度为n
	// 下标范围0~n-1
	// 数范围1~n
	n := len(nums)
	for _,v := range nums {
		i := (v-1)%n
		nums[i] += n
	}
	ans := []int{}
	for i,v := range nums {
		if v > 2*n {
			ans = append(ans,i+1)
		}
	}
	return ans
}
// @lc code=end

