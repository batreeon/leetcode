/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 */

// @lc code=start
package main

import (
	"strconv"
	"strings"
)

func calculate(s string) int {
	isNum := func(s string) bool {
		_, err := strconv.ParseFloat(s, 64)
		return err == nil
	}
	arrayN := strings.FieldsFunc(s,func(r rune) bool {
		return r == '+' || r == '-' || r == '*' || r == '/' || r == ' '
	})	//数字
	signals :=strings.FieldsFunc(s,func(r rune) bool {
		return (r >= '0' && r <= '9') || r == ' '
	})	//符号
	stoiSlice := func(ss *[]string) *[]int {
		nums := []int{}
		for _,s := range *ss {
			if isNum(s) {
				num,_ := strconv.Atoi(s)
				nums = append(nums,num)
			}
		}
		return &nums
	}
	nums := *stoiSlice(&arrayN)
	j := 0
	for _,signal := range signals {
		if  signal == "*" || signal == "/" {
			res1 := nums[j]
			res2 := nums[j+1]
			if signal == "*" {
				res1 = res1*res2
			}else{
				res1 = res1/res2
			}
			nums[j] = res1
			if len(nums[j:]) > 1 {
				copy(nums[j+1:],nums[j+2:])
				nums = nums[:len(nums)-1]
			}
		}else{
			j++
		}
	}
	j = 0
	for _,signal := range signals {
		if signal == "*" || signal == "/" {
			continue
		}
		switch signal {
		case "+":
			nums[0] = nums[0]+nums[1]
		case "-":
			nums[0] = nums[0]-nums[1]
		}
		if len(nums[1:]) > 1 {
			copy(nums[j+1:],nums[j+2:])
			nums = nums[:len(nums)-1]
		}
	}
	return nums[0]
}
// @lc code=end

