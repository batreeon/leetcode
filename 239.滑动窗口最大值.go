/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
package main

import (
	"container/heap"
	"sort"
)

var subnums []int
// 匿名字段，继承sort.IntSlice的三个方法
type hp struct{sort.IntSlice}
func (h hp) Less(i,j int) bool {return subnums[h.IntSlice[i]] > subnums[h.IntSlice[j]]}
func (h *hp) Push(x interface{}) {h.IntSlice = append(h.IntSlice,x.(int))}
func (h *hp) Pop() interface{} {
	old := h.IntSlice
	n := len(old)
	x := old[n-1]
	h.IntSlice = old[:n-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	subnums = nums
	h := &hp{make([]int,k)}
	for i := 0 ; i < k ; i++ {
		h.IntSlice[i] = i
	}
	heap.Init(h)
	n := len(nums)
	result := make([]int,1,n-k+1)
	// 注意h.IntSilce存的是下标，还需要到subnums中取值
	result[0] = subnums[h.IntSlice[0]]
	for i := k ; i < n ; i++ {
		heap.Push(h,i)
		for h.IntSlice[0] <= i-k {
			// 当前的最大值下标已经不再窗口里面了，pop
			heap.Pop(h)
		}
		result = append(result,subnums[h.IntSlice[0]])
	}
	return result
}
// @lc code=end

