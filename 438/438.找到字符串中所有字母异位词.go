/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
package main

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	mp := map[byte]int{}
	for _, b := range p {
		mp[byte(b)]++
	}
	l, r := 0, 0
	ms := map[byte]int{}
	result := []int{}

	// 这个for循环判断是要注意，s[l]还没有加入ms中
	for l + len(p) - 1  < len(s) {
		for r - l + 1 <= len(p) {
			// 将s中len(p)长度的子串加入到ms中
			ms[s[r]]++
			r++
		}
		if len(ms) == len(mp) {
			// 如果两个map长度相等在判断是否完全一致
			if equal(ms, mp) {
				result = append(result, l)
			}
		}

		// 无论是否完全一致，是s[l]都需要丢弃了，因为我们已经检查了以s[l]开头的子串
		ms[s[l]]--
		if ms[s[l]] == 0 {
			delete(ms, s[l])
		}
		l++

		for len(ms) > len(mp) {
			ms[s[l]]--
			if ms[s[l]] == 0 {
				delete(ms, s[l])
			}
			l++
		}
	}
	return result
}

func equal(m1, m2 map[byte]int) bool {
	for k1, v1 := range m1 {
		if v1 != m2[k1] {
			return false
		}
	}
	return true
}
// @lc code=end

