/*
 * @lc app=leetcode.cn id=1512 lang=golang
 *
 * [1512] 好数对的数目
 */

// @lc code=start
func numIdenticalPairs(nums []int) int {
    m := map[int]int{}

	result := 0
	for _, num := range nums {
		result += m[num]
		m[num]++
	}

	return result
}
// @lc code=end

