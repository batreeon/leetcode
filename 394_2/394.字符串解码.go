/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
// 2[abc]3[cd]ef
// stack
// 2
// 2[
// 3[2[abc]] => 3[abcabc]
package main

import "strconv"

func decodeString(s string) string {
	stack := []rune{}

	for _, r := range s {
		if isAlpha(r) || isDigit(r) {
			stack = append(stack, r)
		}
		if r == '[' {
			stack = append(stack, r)
		}
		if r == ']' {
			i := len(stack) - 1
			for ; stack[i] != '['; i-- {
			}

			temp := make([]rune, len(stack)-(i+1))
			copy(temp, stack[i+1:])
			stack = stack[:i]

			numr := len(stack) - 1
			numl := len(stack) - 1
			for ; numl >= 0 && isDigit(stack[numl]); numl-- {
			}

			count, _ := strconv.Atoi(string(stack[numl+1 : numr+1]))
			stack = stack[:numl+1]
			for i := 0; i < count; i++ {
				stack = append(stack, temp...)
			}
		}
	}
	return string(stack)
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isAlpha(r rune) bool {
	return r >= 'a' && r <= 'z'
}

// @lc code=end
