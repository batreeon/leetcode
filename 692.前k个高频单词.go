/*
 * @lc app=leetcode.cn id=692 lang=golang
 *
 * [692] 前K个高频单词
 */

// @lc code=start
package main
import "sort"

type sortedWords struct{
	words []string
	cnt map[string]int
}
func newSortedWords(words []string) sortedWords {
	s := sortedWords{[]string{},map[string]int{}}
	for i := range words {
		if _,ok := s.cnt[words[i]] ; !ok {
			s.words = append(s.words,words[i])
		}
		s.cnt[words[i]]++
	}
	return s
}
func (s sortedWords) Len() int {
	return len(s.words)
}
func (s sortedWords) Less(i,j int) bool {
	return s.cnt[s.words[i]] > s.cnt[s.words[j]] || (s.cnt[s.words[i]] == s.cnt[s.words[j]] && s.words[i] < s.words[j])
}
func (s sortedWords) Swap(i,j int) {
	s.words[i],s.words[j] = s.words[j],s.words[i]
}
func topKFrequent(words []string, k int) []string {
	s := newSortedWords(words)
	sort.Sort(s)
	return s.words[:k]
}
// @lc code=end

