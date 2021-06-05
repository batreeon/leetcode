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
	var f func(str string,idx,pre int) bool
	f = func(str string,idx,pre int) bool {
		if len(str) == 0 {
			// 拆分的子数组个数不够
			return idx >= 2
		}
		sstr := strings.TrimLeft(str,"0")
		if len(str) != 0 && len(sstr) == 0 {
			return pre == 1
		}
		for i := 1 ; i <= len(sstr) ; i++ {
			num,_ := strconv.Atoi(sstr[:i])
			if pre != -1 {
				if num >= pre {
					// i++之后，num只会更大，后面的就不用再看了
					break
				}
				if num != pre-1 {
					continue
				}
			}
			if f(sstr[i:],idx+1,num) {
				return true
			}
		}
		return false
	}

	return f(ss,0,-1)
}
// @lc code=end

