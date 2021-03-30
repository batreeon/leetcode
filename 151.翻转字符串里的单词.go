/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 */

// @lc code=start
package main

import "strings"

func reverseWords(s string) string {
	// strings.Trim(s," ")
	// ss := strings.Fields(s)
	// reverse := func(b []byte) string {
	// 	l := len(b)
	// 	for i,j := 0,l-1 ; i < j ; i,j = i+1,j-1 {
	// 		b[i],b[j] = b[j],b[i]
	// 	}
	// 	return string(b)
	// }
	// for i,s := range ss {
	// 	ss[i] = reverse([]byte(s))
	// }
	// return strings.Join(ss," ")
	// 卧槽，题都没读懂

	strings.Trim(s, " ")
	ss := strings.Fields(s)
	reverse := func(ss []string) []string {
		l := len(ss)
		for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
			ss[i], ss[j] = ss[j], ss[i]
		}
		return ss
	}

	reversedSs := reverse(ss)
	return strings.Join(reversedSs, " ")
}

// @lc code=end
