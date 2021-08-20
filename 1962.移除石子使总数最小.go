/*
 * @lc app=leetcode.cn id=1962 lang=golang
 *
 * [1962] 移除石子使总数最小
 */

// @lc code=start
package main

import (
	"container/heap"
)

type maxHeap []int
func (h maxHeap) Len() int {return len(h)}
func (h maxHeap) Less(i,j int) bool {return h[i] > h[j]}
func (h maxHeap) Swap(i,j int) {h[i],h[j] = h[j],h[i]}
func (h *maxHeap) Push(x interface{}) {
	(*h) = append(*h,x.(int))
}
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	(*h) = (*h)[:n-1]
	return x
}
func (h maxHeap) Top() interface{} {
	return h[0]
}


func minStoneSum(piles []int, k int) int {
	sum := 0
	for _,pile := range piles {
		sum += pile
	}
	if sum <= 1 {
		return sum
	}

	// 只有一堆石头，就不需要使用堆了
	if len(piles) == 1 {
		for sum != 1 && k > 0 {
			sum -= sum/2
			k--
		}
		return sum
	}

	// 多于1堆，需要建堆
	h := maxHeap(piles)
	heap.Init(&h)
	for k > 0 {
		x := heap.Pop(&h).(int)
		top := h.Top().(int)
		for x >= top && k > 0 {
			sum -= x/2
			if sum == len(h)+1 {
				// 所有堆都只有一个石子
				return sum
			}
			x -= x/2
			k--
		}
		// 不需要判断x == 0 ,x最小只能为1了
		heap.Push(&h,x)
	}
	return sum
}
// @lc code=end

