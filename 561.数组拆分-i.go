/*
 * @lc app=leetcode.cn id=561 lang=golang
 *
 * [561] 数组拆分 I
 */

// @lc code=start
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0 ; i < len(nums) ; i += 2 {
		ans += nums[i]
	}
	return ans
}
// @lc code=end

