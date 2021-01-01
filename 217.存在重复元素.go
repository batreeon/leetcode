/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	nummap := make(map[int]bool)
	for _,num := range nums {
		if nummap[num] == true {
			return true
		}else {
			nummap[num] = true
		}
	}
	return false
}
// @lc code=end

