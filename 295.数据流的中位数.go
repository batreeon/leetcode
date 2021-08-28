/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 */

// @lc code=start
package main

import (
	"sort"
	"container/heap"
)

type _heap struct {sort.IntSlice}
func (h *_heap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice,x.(int))
}
func (h *_heap) Pop() interface{} {
	old := h.IntSlice
	n := len(old)
	x := old[n-1]
	h.IntSlice = old[:n-1]
	return x
}
func (h _heap) Top() interface{} {
	return h.IntSlice[0]
}
type smallHeap struct {_heap}
type bigHeap struct {_heap}
func (b bigHeap) Less(i,j int) bool {return b.IntSlice[i] > b.IntSlice[j]}


type MedianFinder struct {
	*bigHeap
	*smallHeap
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	b := &bigHeap{}
	heap.Init(b)
	s := &smallHeap{}
	heap.Init(s)
	return MedianFinder{b,s}
}


func (this *MedianFinder) AddNum(num int)  {
	if (len(this.smallHeap.IntSlice) + len(this.bigHeap.IntSlice)) & 1 == 1 {
		heap.Push(this.bigHeap,num)
		heap.Push(this.smallHeap,heap.Pop(this.bigHeap))
	}else{
		heap.Push(this.smallHeap,num)
		heap.Push(this.bigHeap,heap.Pop(this.smallHeap))
	}
}


func (this *MedianFinder) FindMedian() float64 {
	if (len(this.bigHeap.IntSlice) + len(this.smallHeap.IntSlice)) & 1 == 1 {
		return float64(this.bigHeap.Top().(int))
	}else{
		return float64(this.bigHeap.Top().(int) + this.smallHeap.Top().(int)) / 2.0
	}
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

