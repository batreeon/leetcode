/*
 * @lc app=leetcode.cn id=1018 lang=golang
 *
 * [1018] 可被 5 整除的二进制前缀
 */

// @lc code=start
func prefixesDivBy5(A []int) []bool {
	pre := 0
	ans := []bool{}
	for _,a := range A {
		pre = ((pre<<1)|a) % 5
		ans = append(ans,pre==0)
	}
	return ans
}
// @lc code=end

