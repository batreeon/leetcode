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
		// 如果只是一个单独的数字，该表达式就不用分了，直接返回
		if strings.IndexAny(s,"+-*") == -1 {
			num,_ := strconv.Atoi(s)
			return []int{num}
		}

		// 记录运算符号的位置，每个符号都是一个分割点
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

		// 记录最终的结果，是切片因为有多个答案
		temp := []int{}
		// 将表达式分成两半，递归处理左右部分
		// 然后再用他们中间的运算符号，计算左右切片中数值的所有组合
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

