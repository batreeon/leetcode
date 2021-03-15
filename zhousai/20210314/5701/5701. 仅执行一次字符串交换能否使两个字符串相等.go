package main

import "strings"

/*
给你长度相等的两个字符串 s1 和 s2 。一次 字符串交换 操作的步骤如下：选出某个字符串中的两个下标（不必不同），并交换这两个下标所对应的字符。

如果对 其中一个字符串 执行 最多一次字符串交换 就可以使两个字符串相等，返回 true ；否则，返回 false 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/check-if-one-string-swap-can-make-strings-equal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func areAlmostEqual(s1 string, s2 string) bool {
	l := len(s1)
	i := 0
	for ; i < l; i++ {
		if s1[i] != s2[i] {
			break
		}
	}
	if i == l {
		return true
	} else {
		j := i + 1
		for ; j < l; j++ {
			if s1[j] != s2[j] {
				break
			}
		}
		if j == l {
			return false
		} else {
            if (s1[i] == s2[j]) && (s1[j] == s2[i]) {
                return strings.Compare(string(s1[j+1:]), string(s2[j+1:])) == 0
            }else{
                return false
            }
		}
	}
	return false
}