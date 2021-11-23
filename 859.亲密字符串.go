/*
 * @lc app=leetcode.cn id=859 lang=golang
 *
 * [859] 亲密字符串
 */

// @lc code=start
func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	idx1 := -1
	idx2 := -1
	letters := make([]int, 26)
	for i := 0; i < len(s); i++ {
		letters[s[i] - 'a']++ 
		if s[i] != goal[i] {
			if idx1 == -1 {
				idx1 = i
			}else if idx2 == -1 {
				idx2 = i
				// ab,ca
				if s[idx1] != goal[idx2] || s[idx2] != goal[idx1] {
					return false
				}
			}else {
				// abc,bad
				return false
			}
		}
	}
	// abc, dbc
	if idx1 != -1 && idx2 == -1 {
		return false
	}
	//	aa,aa; aaab,aaab
	if idx1 == -1 {
		for _,j := range letters {
			if j >= 2 {
				return true
			}
		}
		return false
	}
	return true
}
// @lc code=end

