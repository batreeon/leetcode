/*
 * @lc app=leetcode.cn id=726 lang=golang
 *
 * [726] 原子的数量
 */

// @lc code=start
package main

import (
	"sort"
	"strconv"
	"unicode"
)


func countOfAtoms(formula string) string {
	i,n := 0,len(formula)

	parseAtom := func() string {
		start := i
		i++
		for i < n && unicode.IsLower(rune(formula[i])) {
			i++
		}
		return formula[start:i]
	}
	parseNum := func() int {
		num := 0
		if i == n || !unicode.IsDigit(rune(formula[i])) {
			return 1
		}
		for i < n && unicode.IsDigit(rune(formula[i])) {
			num = num*10 + int(formula[i]-'0')
			i++
		}
		return num
	}
	
	stk := []map[string]int{{}}
	// 初始化了数组的第一项，空map

	for i < n {
		if formula[i] == '(' {
			i++
			stk = append(stk,map[string]int{})
		}else if formula[i] == ')' {
			i++
			num := parseNum()
			last := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			for mole,v := range last {
				stk[len(stk)-1][mole] += num*v
			}
		}else{
			mole := parseAtom()
			num := parseNum()
			stk[len(stk)-1][mole] += num
		}
	}
	type pair struct{
		mole string
		times int
	}
	moles := stk[0]
	pairs := []pair{}
	for mole,v := range moles {
		pairs = append(pairs,pair{mole,v})
	}
	sort.Slice(pairs,func(i,j int) bool {return pairs[i].mole < pairs[j].mole})
	result := []byte{}
	for _,pair := range pairs {
		result = append(result,pair.mole...)
		if pair.times > 1 {
			result = append(result,strconv.Itoa(pair.times)...)
		}
	}
	return string(result)
}
// @lc code=end

