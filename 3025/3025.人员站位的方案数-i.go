/*
 * @lc app=leetcode.cn id=3025 lang=golang
 *
 * [3025] 人员站位的方案数 I
 */

// @lc code=start
package main

import (
	"math"
	"sort"
)

func numberOfPairs(points [][]int) int {
    sort.Slice(points, func(i, j int) bool {
        if points[i][0] < points[j][0] {
            return true
        }else if points[i][0] == points[j][0] {
            return points[i][1] >= points[j][1]
        }
        return false
    })

    result := 0
    for i, p := range points{
        y1 := p[1]
        maxY := math.MinInt32
        for _, q := range points[i+1:] {
            y2 := q[1]
            if y2 <= y1 && y2 > maxY {
                maxY = y2
                result++
            }
        }
    }

    return result
}
// @lc code=end

