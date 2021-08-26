/*
 * @lc app=leetcode.cn id=881 lang=golang
 *
 * [881] 救生艇
 */

// @lc code=start
package main

import (
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	l,r := -1,len(people)-1
	result := 0
	for l < r {
		sum := people[r]
		for l < r && sum+people[l+1] <= limit {
			sum += people[l+1]
			l++
			break
			// 每艘船只能载两人，栽了两人后就直接跳出
		}
		result++
		r--
	}
	return result
}
// @lc code=end

