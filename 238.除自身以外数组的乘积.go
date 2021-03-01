/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	//抄的官解下方的评论，华科江城慕雪，太强了
	n := len(nums)
	ans := make([]int,n)
	prefix,suffix := 1,1
	for i := 0 ; i < n ; i++ {
		if n-i-1 >= i {
			ans[i],ans[n-i-1] = 1,1

		}
		ans[i] *= prefix
		ans[n-i-1] *= suffix
		prefix *= nums[i]
		suffix *= nums[n-i-1]
	}
	return ans
}
// @lc code=end

