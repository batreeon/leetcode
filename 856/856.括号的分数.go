/*
 * @lc app=leetcode.cn id=856 lang=golang
 *
 * [856] 括号的分数
 */

// @lc code=start
package main

func scoreOfParentheses(S string) int {
	var score func(s string) int 
	score = func(s string) int {
		if len(s) == 0 {
			return 0
		}
		if s[1] == ')' {
			return 1 + score(s[2:])
		}else{
			stack := []byte{'(','('}
			i := 2
			for len(stack) > 0 {
				if s[i] == '(' {
					stack = append(stack,'(')
				}else{
					stack = stack[:len(stack)-1]
				}
				i++
			}
			return 2 * score(s[1:i-1]) + score(s[i:])
		}
	}
	return score(S)
}
// @lc code=end

