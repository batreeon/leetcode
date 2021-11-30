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
	for l + len(p) - 1  < len(s) {
		for r - l + 1 <= len(p) {
			ms[s[r]]++
			r++
		}
		if len(ms) == len(mp) {
			if equal(ms, mp) {
				result = append(result, l)
			}
			ms[s[l]]--
			if ms[s[l]] == 0 {
				delete(ms, s[l])
			}
			l++
		}else{
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

