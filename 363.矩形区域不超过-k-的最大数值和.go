/*
 * @lc app=leetcode.cn id=363 lang=golang
 *
 * [363] 矩形区域不超过 K 的最大数值和
 */

// @lc code=start
package main

import (
	"math"
	"sort"
)

func max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSumSubmatrix(matrix [][]int, k int) int {
	result := math.MinInt64
	r,c := len(matrix),len(matrix[0])
	for i := 0 ; i < r ; i++ {//上边界
		sum := make([]int,c)
		for j := i ; j < r ; j++ {//下边界
			for l := 0 ; l < c ; l++ {
				sum[l] += matrix[j][l]//更新累加值
			}
			sumSet := []int{0}
			s := 0
			// 找到最大值
			for _,v := range sum {
				s += v
				lb := sort.Search(len(sumSet),func(i int) bool {return sumSet[i] >= s-k})
				if lb != len(sumSet) {
					result = max(result,s-sumSet[lb])
				}
				sumSet = append(sumSet,s)
				if !sort.IsSorted(sort.IntSlice(sumSet)) {
					sort.Ints(sumSet)
				}
			}
		}
	}
	return result
}
// @lc code=end

