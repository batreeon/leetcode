/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
package main

import "strings"

func lengthOfLastWord(s string) int {
    ss := strings.Fields(s)
	return len(ss[len(ss)-1])
}
// @lc code=end

