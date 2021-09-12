/*
 * @lc app=leetcode.cn id=678 lang=golang
 *
 * [678] 有效的括号字符串
 */

// @lc code=start
package main

// import "fmt"

func checkValidString(s string) bool {
	/*好难啊，我不会
	// 贪心维护一个区间，左括号数量可能的值所在的区间
	lo, hi := 0, 0
	for _, b := range s {
		if b == '(' {
			lo++
			hi++
		} else if b == '*' {
			// 为什么要加这个判断，请考虑 (**)( 这个例子
			if lo > 0 {
				lo--
			}
			hi++
		} else {
			// 考虑 (*)
			if lo > 0 { lo-- }
			hi--
		}
		if hi < 0 {
			return false
		}
	}
	return lo == 0
	*/
	left,replace := []int{},[]int{}
	for i,b := range s {
		if b == '(' {
			left = append(left,i)
		}else if b == '*' {
			replace = append(replace,i)
		}else{
			if len(left) == 0 && len(replace) == 0 {
				return false
			}else if len(left) > 0 {
				left = left[:len(left)-1]
			}else{
				// 注意这个尽量把*留在后面用来匹配(
				// 这个太妙了
				replace = replace[1:]
			}
		}
	}
	if len(left) > len(replace) {
		return false
	}
	for i,j := 0,0 ; i < len(left) ; i,j = i+1,j+1 {
		// 因为星号可能也能替换成空，因此左括号左边有星号也可能是合法的
		for j < len(replace) && left[i] > replace[j] {
			j++
		}
		if j == len(replace) {
			return false
		}
	}
	return true
}

// @lc code=end
