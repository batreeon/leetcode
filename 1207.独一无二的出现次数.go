/*
 * @lc app=leetcode.cn id=1207 lang=golang
 *
 * [1207] 独一无二的出现次数
 */

// @lc code=start
func uniqueOccurrences(arr []int) bool {
	num := make(map[int]int)
	for _,v := range arr {
		num[v]++
	}
	numnum := make(map[int]int)
	for _,v := range num {
		numnum[v]++
		if numnum[v] > 1 {
			return false
		}
	}
	return true
}
// @lc code=end

