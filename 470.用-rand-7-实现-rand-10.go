/*
 * @lc app=leetcode.cn id=470 lang=golang
 *
 * [470] 用 Rand7() 实现 Rand10()
 */

// @lc code=start
func rand10() int {
    var r,c,idx int
	idx = 41
	for idx > 40 {
		r,c = rand7(),rand7()
		idx = c + (r-1) * 7
	}
	return (idx - 1)%10 + 1
}
// @lc code=end

