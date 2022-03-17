/*
 * @lc app=leetcode.cn id=720 lang=golang
 *
 * [720] 词典中最长的单词
 */

// @lc code=start
package main
import "sort"
func longestWord(words []string) string {
	wordsMap := map[string]bool{}
	result := ""
	lenResult := 0
	sort.Strings(words)
	for _, word := range words {
		if len(word) == 1 {
			wordsMap[word] = true
			if lenResult < 1 {
				lenResult = 1
				result = word
			}
		}else{
			wordsMap[word] = wordsMap[word[:len(word)-1]]
			if wordsMap[word] && lenResult < len(word) {
				lenResult = len(word)
				result = word
			}
		}
	}
	return result
}
// @lc code=end

