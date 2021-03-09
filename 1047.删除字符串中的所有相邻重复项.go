/*
 * @lc app=leetcode.cn id=1047 lang=golang
 *
 * [1047] 删除字符串中的所有相邻重复项
 */

// @lc code=start
func removeDuplicates(S string) string {
	// idx指示上一个未被删除的元素
	idx := -1
	n := len(S)
	s := []byte(S)
	for i := 0 ; i < n ; i++ {
		if idx == -1 || s[idx] != s[i] {
			idx++
			s[idx] = s[i]
		}else{
			idx--
		}
	}
	return string(s[:idx+1])
}
// @lc code=end

