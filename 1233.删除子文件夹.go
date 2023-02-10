/*
 * @lc app=leetcode.cn id=1233 lang=golang
 *
 * [1233] 删除子文件夹
 */

// @lc code=start
package main

import "strings"

// 笨方法，查找folder[i]的所以可能的父文件夹是否存在
func removeSubfolders(folder []string) []string {
	m := map[string]struct{}{}
	for _, f := range folder {
		m[f] = struct{}{}
	}

	for _, f := range folder {
		slashIndex := strings.LastIndex(f, "/")
		for ; slashIndex != 0; slashIndex = strings.LastIndex(f[:slashIndex-1], "/") {
			pre := f[:slashIndex]
			if _, ok := m[pre]; ok {
				delete(m, f)
			}
		}
	}

	result := []string{}
	for k := range m {
		result = append(result, k)
	}

	return result
}
// @lc code=end

