/*
 * @lc app=leetcode.cn id=1849 lang=golang
 *
 * [1849] 将字符串拆分为递减的连续值
 */

// @lc code=start
package main

import (
	"strconv"
	"strings")


func splitString(s string) bool {
	ss := strings.TrimLeft(s,"0")
	var f func(str string,pre int) bool
	f = func(str string,pre int) bool {
		sstr := strings.TrimLeft(str,"0")
		if len(sstr) == 0 {
			return true
		}
		for i := 1 ; i <= len(sstr) ; i++ {
			if i == len(sstr) && pre == -1 {
				// 没有分两组，不符合条件
				return false
			}
			num,_ := strconv.Atoi(sstr[:i])
			if pre != -1 {
				if num >= pre {
					break
				}
				if num != pre-1 {
					continue
				}
			}
			if f(sstr[i:],num) {
				return true
			}
		}
		return false
	}

	return f(ss,-1)
}
// @lc code=end

