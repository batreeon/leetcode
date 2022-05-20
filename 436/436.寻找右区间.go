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
	// 保存原数组每一个区间的下标，
	// 因为最终返回的结果是区间的下标，所以需要用到
	// 因为每一个区间的起始点不相同，所以只用起始点作为key就可以了
	m1 := map[int]int{}
	for i, interval := range intervals {
		m1[interval[0]] = i 
	}
	
	l := len(intervals)
	// 根据右区间的定义进行排序
	intervalsCopy := make([][]int, l)
	copy(intervalsCopy, intervals)
	sort.Slice(intervalsCopy, func(i,j int) bool {
		if intervalsCopy[i][1] <= intervalsCopy[j][0] {
			return true
		// golang的sort包，只用上面一个条件，排出来的不对，加了下面这条才对了，不太理解
		}else if intervalsCopy[i][0] <= intervalsCopy[j][0]{
			return true
		}
		return false
	})
	// 记录排序后区间的下标
	m2 := map[int]int{}
	for i, interval := range intervalsCopy {
		m2[interval[0]] = i
	}

	result := []int{}
	// 遍历原区间数组，然后看每一个区间时候存在右区间
	for _, interval := range intervals {
		// 区间排序后的下标
		sorti := m2[interval[0]]
		// 排在最右了，那自然不存在右区间
		if sorti == l-1 {
			result = append(result, -1)
		}else{
			// 被坑傻了，原来j可以等于i，例如某个区间[1,1],它可以是他自身的右区间
			right := sorti
			for ; right <= l-1; right++{
				if intervalsCopy[right][0] >= interval[1] {
					// 存在右区间
					result = append(result, m1[intervalsCopy[right][0]])
					break
				}
			}
			// 不存在右区间
			if right == l {
				result = append(result, -1)
			}
		}
	}
	return result
}
// @lc code=end

