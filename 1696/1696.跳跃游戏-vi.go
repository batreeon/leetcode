/*
 * @lc app=leetcode.cn id=1696 lang=golang
 *
 * [1696] 跳跃游戏 VI
 */

// @lc code=start
package main

import (
	"container/heap"
)


type ele struct{
	sum,index int
}
type maxHeap []ele
func (h maxHeap) Len() int {return len(h)}
func (h maxHeap) Less(i,j int) bool {return h[i].sum > h[j].sum}
func (h maxHeap) Swap(i,j int) {h[i],h[j] = h[j],h[i]}
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h,x.(ele))
}
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = (*h)[:n-1]
	return x
}
func (h maxHeap) Top() interface{} {
	return h[0]
}

func maxResult(nums []int, k int) int {
	n := len(nums)
	h := &maxHeap{}
	heap.Init(h)
	heap.Push(h,ele{nums[n-1],n-1})
	result := nums[n-1]
	for i := n-2 ; i >= 0 ; i-- {
		for h.Top().(ele).index - i > k {
			heap.Pop(h)
		}
		topSum := h.Top().(ele).sum
		if i == 0 {
			result = topSum+nums[i]
			break
		}
		heap.Push(h,ele{topSum+nums[i],i})
	}
	return result
}
// @lc code=end

