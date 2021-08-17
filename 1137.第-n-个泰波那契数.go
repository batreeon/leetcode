/*
 * @lc app=leetcode.cn id=1137 lang=golang
 *
 * [1137] 第 N 个泰波那契数
 */

// @lc code=start
func tribonacci(n int) int {
	t0,t1,t2 := 0,1,1
	if n == 0 {
		return t0
	}
	if n == 1 {
		return t1
	}
	if n == 2 {
		return t2
	}
	for n - 2 > 0 {
		t0,t1,t2 = t1,t2,t0+t1+t2
		n--
	}
	return t2
}
// @lc code=end

