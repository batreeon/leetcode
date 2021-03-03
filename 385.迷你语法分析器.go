/*
 * @lc app=leetcode.cn id=385 lang=golang
 *
 * [385] 迷你语法分析器
 */

// @lc code=start
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
package main
import (
	"strings"
	"strconv"
)
//[123,[456,[789]],[34]]
func deserialize(s string) *NestedInteger {
	ss := strings.Trim(s,"]")
	ans := NestedInteger{}
	start := 0
	for strings.IndexAny(ss[start:],"[,") != -1 {
		end := strings.Index(ss[start:],",[")
		if end == -1 {
			end = len(ss)
		}
		arrayN := strings.FieldsFunc(ss[start:end],func(r rune) bool {
			return r == ',' || r == ']' || r == '['
		})
		if start == 0 && ss[0] != '[' {
			for _,a := range arrayN {
				i,_ := strconv.Atoi(a)
				ans.SetInteger(i)
			}
		}else{
			res := NestedInteger{}
			for _,a := range arrayN {
				i,_ := strconv.Atoi(a)
				res.SetInteger(i)
			}
			ans.Add(res)
		}
		start = end+2
		if start > len(ss) {
			break
		} 
	}
	if start <= len(ss) {
		i,_ := strconv.Atoi(ss[start:])
		ans.SetInteger(i)
	}
	
	return &ans
}
// @lc code=end

