/*
 * @lc app=leetcode.cn id=187 lang=golang
 *
 * [187] 重复的DNA序列
 */

// @lc code=start
package main
import "strings"

func findRepeatedDnaSequences(s string) []string {
	resultMap := map[string]bool{}
	result := []string{}
	n := len(s)
	for i := 0 ; i <= n-11 ; i++ {
		if resultMap[s[i:i+10]] {
			continue
		}
		if strings.Index(s[i+1:],s[i:i+10]) != -1 {
			resultMap[s[i:i+10]] = true
			result = append(result,s[i:i+10])
		}
	}
	return result
}
// @lc code=end

