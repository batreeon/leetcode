/*
 * @lc app=leetcode.cn id=306 lang=golang
 *
 * [306] 累加数
 */

// @lc code=start
package main

import (
	"strconv"	
	"strings"
)


func isAdditiveNumber(num string) bool {

	var backtrack func(pre1,pre2 int , s string) bool 
	backtrack = func(pre1,pre2 int , s string) bool {
		pre1,pre2 = pre2,pre1+pre2
		pre2str := strconv.Itoa(pre2)
		// 到末尾了
		if pre2str == s {
			return true
		}

		// 因为使用字符串匹配的，不用再单独考虑开头为0的情况了
		if strings.Index(s,pre2str) != 0 {
			return false
		}else{
			if backtrack(pre1,pre2,s[len(pre2str):]) {
				return true
			}
			return false
		}
	}

	n := len(num)
	for i := 1 ; i <= n-2 ; i++ {
		// 0开头只能是单0
		if num[0] == '0' && i > 1 {
			break
		}
		for j := i+1 ; j <= n-1 ; j++ {
			// 0开头只能是单0
			if num[i] == '0' && j != i+1 {
				break
			}
			pre1,_ := strconv.Atoi(num[:i])
			pre2,_ := strconv.Atoi(num[i:j])
			if backtrack(pre1,pre2,num[j:]) {
				return true
			}
		}
	}

	return false
}
// @lc code=end

