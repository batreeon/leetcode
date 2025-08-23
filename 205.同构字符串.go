/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	sr, tr := []rune(s), []rune(t)
	s2t := make(map[rune]rune)
	tset := make(map[rune]struct{})
	for i, sb := range sr {
		if _, ok := s2t[sb]; !ok {
			if _, ok1 := tset[tr[i]]; ok1 {
				return false
			}

			s2t[sb] = tr[i]
			tset[tr[i]] = struct{}{}
		} else {
			if s2t[sb] != tr[i] {
				return false
			}
		}
	}

	return true
}

// @lc code=end

