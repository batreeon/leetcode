/*
给你一个二进制字符串 s ，该字符串 不含前导零 。

如果 s 最多包含 一个由连续的 '1' 组成的字段 ，返回 true​​​ 。否则，返回 false

1 <= s.length <= 100
s[i]​​​​ 为 '0' 或 '1'
s[0] 为 '1'
*/
package main

import "strings"

func checkOnesSegment(s string) bool {
	pre := strings.Index(s,"0")
	if pre == -1 {
		return true
	}
	if strings.Index(s[pre+1:],"1") != -1 {
		return false
	}
	return true
}