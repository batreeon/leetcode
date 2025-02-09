/*
 * @lc app=leetcode.cn id=452 lang=golang
 *
 * [452] 用最少数量的箭引爆气球
 */

// @lc code=start
package main

import "sort"
// 简单区间题

func findMinArrowShots(points [][]int) int {
    res := 0
	n := len(points)
	if n == 1 {
		return 1
	}
	sort.Slice(points, func(i, j int) bool {return points[i][0] < points[j][0]})
	for i := 0; i < n; {
		l, r := points[i][0], points[i][1]
		j := i + 1
		for ; j < n && points[j][0] <= r; j++ {
			l = max(l, points[j][0])
			r = min(r, points[j][1])
		}
		i = j 
		res++
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

