/*
 * @lc app=leetcode.cn id=1833 lang=golang
 *
 * [1833] 雪糕的最大数量
 */

// @lc code=start
func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	result := 0
	for i := range costs {
		coins -= costs[i]
		if coins < 0 {
			break
		}
		result++
	}
	return result
}
// @lc code=end

