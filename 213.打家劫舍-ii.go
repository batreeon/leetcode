/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	n := len(nums)
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	if n == 1 {
		return nums[0]
	}

	pre2 := nums[0]
	pre1 := max(pre2,nums[1])
	for i := 2 ; i < n-1 ; i++ {
		f := max(pre1,pre2+nums[i])
		pre2,pre1 = pre1,f
	}
	result := pre1
	if n != 2 {
		pre2 = nums[1]
		pre1 = max(pre2,nums[2])
		for i := 3 ; i < n ; i++ {
			f := max(pre1,pre2+nums[i])
			pre2,pre1 = pre1,f
		}
		result = max(result,pre1)
	}

	return result
}
// @lc code=end

