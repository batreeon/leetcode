/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
    cur := 0
	next := 0
	result := 0
	for i, num := range nums[:len(nums)-1] {
		next = max(next, i+num)
		if i == cur {
			cur = next
			result++
		}
	}
	return result
}
// @lc code=end

