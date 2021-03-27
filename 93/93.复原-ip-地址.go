/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */

// @lc code=start
package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	ans := []string{}
	var restoreipaddress func(s string, res []string)
	restoreipaddress = func(s string, res []string) {
		if len(res) == 4 {
			copyRes := make([]string, 4)
			copy(copyRes, res)
			ans = append(ans, strings.Join(copyRes, "."))
			return
		}
		l := len(s)
		// 不算当前组，后续还剩几组
		back := 4 - len(res) - 1
		// i是当前分段取得的部分尾下标，包括
		// 因此留给剩余部分的字符个数=l-1-i

		// 最后一组，就没必要分了，剩下的字符全是最后一组的
		if back == 0 {
			// 位数大于1且高位为0，则不合法
			if len(s) <= 1 || s[0] != '0' {
				num, _ := strconv.Atoi(s)
				if num >= 0 && num <= 255 {
					res = append(res, s)
					restoreipaddress(string(""), res)
					res = res[:len(res)-1]
				}
			}
		} else {
			for i := 0 ; i < l && i < 3 ; i++ {
				// 剩余元素不能过多
				if l-i-1 > 3*back {
					continue
				}
				// 剩余元素不能过少，过少就可以提前退出了
				if l-i-1 < back {
					break
				}
				// 位数大于1且高位为0，可以提前退出了
				if i != 0 && s[0] == '0' {
					break
				}
				num, _ := strconv.Atoi(s[:i+1])
				if num >= 0 && num <= 255 {
					res = append(res, s[:i+1])
					restoreipaddress(s[i+1:], res)
					res = res[:len(res)-1]
				}
				// 可以提前退出了
				if num >= 255 {
					break
				}
			}
		}

	}

	restoreipaddress(s, []string{})
	return ans
}

// @lc code=end
