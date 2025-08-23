/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
    idxs := make(map[int]int)
	for i, num := range nums {
		if lastNumIdx, ok := idxs[num]; ok {
			if i - lastNumIdx <= k {
				return true
			}
		}
		idxs[num] = i
	}
	return false
}
// @lc code=end

