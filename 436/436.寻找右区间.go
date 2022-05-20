/*
 * @lc app=leetcode.cn id=436 lang=golang
 *
 * [436] 寻找右区间
 */

// @lc code=start
package main

import (
	"sort"
)

func findRightInterval(intervals [][]int) []int {
	m1 := map[int]int{}
	for i, interval := range intervals {
		m1[interval[0]] = i 
	}
	
	l := len(intervals)
	intervalsCopy := make([][]int, l)
	copy(intervalsCopy, intervals)
	sort.Slice(intervalsCopy, func(i,j int) bool {
		if intervalsCopy[i][1] <= intervalsCopy[j][0] {
			return true
		}else if intervalsCopy[i][0] <= intervalsCopy[j][0]{
			return true
		}
		return false
	})
	m2 := map[int]int{}
	for i, interval := range intervalsCopy {
		m2[interval[0]] = i
	}

	result := []int{}
	for _, interval := range intervals {
		sorti := m2[interval[0]]
		if sorti == l-1 {
			result = append(result, -1)
		}else{
			// 被坑傻了，原来j可以等于i，例如某个区间[1,1],它可以是他自身的右区间
			right := sorti
			for ; right <= l-1; right++{
				if intervalsCopy[right][0] >= interval[1] {
					result = append(result, m1[intervalsCopy[right][0]])
					break
				}
			}
			if right == l {
				result = append(result, -1)
			}
		}
	}
	return result
}
// @lc code=end

