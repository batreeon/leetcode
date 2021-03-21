/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 */

// @lc code=start
package main

import "sort"

type doubleSort [][]int

func (ds doubleSort) Len() int {
	return len(ds)
}
func (ds doubleSort) Less(i, j int) bool {
	// if ds[i][0] < ds[j][0] && ds[i][1] < ds[j][1] {
	// 	return true
	// }
	// if ds[i][0] == ds[j][0] && ds[i][1] < ds[j][1] {
	// 	return true
	// }
	// if ds[i][0] < ds[j][0] && ds[i][1] == ds[j][1] {
	// 	return true
	// }
	// return false
	if ds[i][0] < ds[j][0] || (ds[i][0] == ds[j][0] && ds[i][1] > ds[j][1]) {
		return true
	}
	return false
}
func (ds doubleSort) Swap(i, j int) {
	ds[i], ds[j] = ds[j], ds[i]
}

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 || len(envelopes) == 1 {
		return len(envelopes)
	}
	ds := doubleSort(envelopes)
	sort.Sort(ds)

	/*
		triangle := [][][]int{}
		f := [][]int{}//标记路径
		start := 0
		for i := 1 ; i < len(ds) ; i++ {
			if ds.Less(i-1,i) {
				triangle = append(triangle,ds[start:i])
				f = append(f,make([]int,i-start))
				start = i
			}
		}
		if start < len(ds) {
			triangle = append(triangle,ds[start:])
			f = append(f,make([]int,len(ds)-start))
		}
		if len(triangle) == 1 {
			return 1
		}
		//ans := 0
		l := len(triangle)
		for i := 0 ; i < l ; i++ {
			for j := 0 ; j < len(triangle[i]) ; j++ {
				if i == 0 {
					f[i][j] = 1
				}else{

				}
			}
		}
		return len(triangle)
	*/
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	l := len(ds)
	f := make([]int, l)
	ans := 0
	// 本质和300题一样，找最长严格递增序列
	// 前面定义的doublesort排序规则很重要，先按长递增排序，若长相同，则按宽递减排序
	for i := 0; i < l; i++ {
		maxf := 0
		for j := 0; j < i; j++ {
			if ds[j][1] < ds[i][1] {
				maxf = max(maxf, f[j])
			}
		}
		f[i] = maxf + 1
		ans = max(ans, f[i])
	}
	return ans
}

// @lc code=end
