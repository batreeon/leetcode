/*
 * @lc app=leetcode.cn id=451 lang=golang
 *
 * [451] 根据字符出现频率排序
 */

// @lc code=start
package main
import "sort"
func frequencySort(s string) string {
	times := map[byte]int{}
	bytes := []byte{}
	for i := range s {
		if _,ok := times[s[i]] ; !ok {
			bytes = append(bytes,s[i])
		}
		times[s[i]]++
	}
	sort.Slice(bytes,func(i,j int) bool {return times[bytes[i]] > times[bytes[j]]})
	var result string
	for _,b := range bytes {
		for i := 0 ; i < times[b] ; i++ {
			result += string(b)
		}
	}
	return result
}
// @lc code=end

