/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
package main
import "math"

func wordBreak(s string, wordDict []string) bool {
	maxLen := math.MinInt64
	minLen := math.MaxInt64
	// 相当于set，记录word
	// m中的value只有true
	m := map[string]bool{}
	// 若s[i:]不合法，就把他记下来，后面如果再要考察s[i:]，就可以直接返回
	// f中value只有false,因为遇到true就可以直接返回了，不用保存
	f := map[string]bool{}
	for _,word := range wordDict {
		if len(word) > maxLen {
			maxLen = len(word)
		}
		if len(word) < minLen {
			minLen = len(word)
		}
		m[word] = true
	}
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	var wordbreak func(s string) bool 
	wordbreak = func(s string) bool {
		if _,ok := m[s] ; ok {
			return true
		}
		if len(s) <= minLen {
			return false
		}
		i := min(len(s),maxLen)
		for ; i >= minLen ; i-- {
			if _,ok := m[s[0:i]] ; ok {
				if _,ok := f[s[i:]] ; ok {
					continue
				}else{
					can := wordbreak(s[i:])
					if can {
						return true
					}else{
						f[s[i:]] = false
					}
				}
			}
		}
		return false
	}
	return wordbreak(s)
}
// @lc code=end

