/*
 * @lc app=leetcode.cn id=1418 lang=golang
 *
 * [1418] 点菜展示表
 */

// @lc code=start
package main

import (
	"sort"	
	"strconv"
)


func displayTable(orders [][]string) [][]string {
	foods := []string{}
	foodsVisited := map[string]bool{}
	tables := []string{}
	tablesVisited := map[string]bool{}
	m := map[string]map[string]int{}
	
	for _,order := range orders {
		if !foodsVisited[order[2]] {
			foodsVisited[order[2]] = true
			foods = append(foods,order[2])
		}
		if !tablesVisited[order[1]] {
			tablesVisited[order[1]] = true
			tables = append(tables,order[1])
		}
		if m[order[1]] == nil {
			m[order[1]] = map[string]int{}
		}
		m[order[1]][order[2]]++
	}

	sort.Slice(tables,func(i,j int) bool {
		x,_ := strconv.Atoi(tables[i])
		y,_ := strconv.Atoi(tables[j])
		return x < y
	})
	sort.Strings(foods)

	result := [][]string{{}}
	result[0] = append(result[0],"Table")
	result[0] = append(result[0],foods...)
	for _,table := range tables {
		temp := make([]string,len(foods)+1)
		temp[0] = table
		for i := 1 ; i < len(temp) ; i++ {
			if _,ok := m[table][foods[i-1]] ; !ok {
				temp[i] = "0"
			}else{
				temp[i] = strconv.Itoa(m[table][foods[i-1]])
			}
		}
		result = append(result,temp)
	}
	return result
}
// @lc code=end

