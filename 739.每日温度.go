/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start
func dailyTemperatures(T []int) []int {
	// 单调栈
	l := len(T)
	ans := make([]int,l)
	//记录下标
	stack := []int{}
	sl := 0

	for i := 0 ; i < l; i++ {
		for sl > 0 && T[stack[sl-1]] < T[i] {
			ans[stack[sl-1]] = i - stack[sl-1]
			stack = stack[:sl-1]
			sl--
		}
		stack = append(stack,i)
		sl++
	}
	return ans
}
// @lc code=end

