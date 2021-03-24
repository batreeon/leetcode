/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

// @lc code=start
func singleNumber(nums []int) int {
	ans := 0
	for _,num := range nums {
		// 若有重复的就异或成0了
		// 且x^0 = x
		ans ^= num
	}
	return ans
}
// @lc code=end

