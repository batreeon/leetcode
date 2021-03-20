package main

import "strings"

/*
给你一个混合字符串 s ，请你返回 s 中 第二大 的数字，如果不存在第二大的数字，请你返回 -1 。

混合字符串 由小写英文字母和数字组成。
*/
func secondHighest(s string) int {
	mapping := func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}
	ss := strings.Map(mapping, s)
	if len(ss) <= 0 {
        return -1
    }
	maxb := rune(ss[0])
	for _, b := range ss {
		if b-'0' > maxb-'0' {
			maxb = b
		}
	}
	ansb := maxb
	for _, b := range ss {
		if b-maxb < 0 {
			if ansb == maxb {
				ansb = b
			}else{
				if b-'0' > ansb-'0' {
					ansb = b
				}
			}
			
		}
	}
	if ansb == maxb {
		return -1
	}
	return int(ansb - '0')
}