/*
 * @lc app=leetcode.cn id=241 lang=golang
 *
 * [241] 为运算表达式设计优先级
 */

// @lc code=start
package main

import (
	"strconv"	
	"strings"
)


func diffWaysToCompute(expression string) []int {

	// result := []int{}
	var compute func(string) []int 
	compute = func(s string) []int {
		if strings.IndexAny(s,"+-*") == -1 {
			num,_ := strconv.Atoi(s)
			return []int{num}
		}
		operators := []int{}
		index := 0
		lastindex := -1
		for {
			index = strings.IndexAny(s[lastindex+1:],"+-*")
			if index == -1 {
				break
			}
			index += lastindex+1
			operators = append(operators,index)
			lastindex = index 
		}
		temp := []int{}
		for _,operator := range operators {
			left := compute(s[:operator])
			right := compute(s[operator+1:])
			for _,l := range left {
				for _,r := range right {
					if s[operator] == '+' {
						temp = append(temp,l+r)
					}else if s[operator] == '-' {
						temp = append(temp,l-r)
					}else{
						temp = append(temp,l*r)
					}
				}
			}
		}
		return temp
	}

	return compute(expression)
}
// @lc code=end

