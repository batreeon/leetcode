/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 */

// @lc code=start
package main

import (
	"container/heap"
)
type bigHeap []int
func (h bigHeap) Len() int {
	return len(h)
}
func (h bigHeap) Less(i,j int) bool {
	return h[i] > h[j]
}
func (h bigHeap) Swap(i,j int) {
	h[i],h[j] = h[j],h[i]
}
func (h *bigHeap) Push(x interface{}) {
	*h = append(*h,x.(int))
}
func (h *bigHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}
func (h bigHeap) Top() interface{} {
	return h[0]
}

type smallHeap []int
func (h smallHeap) Len() int {
	return len(h)
}
func (h smallHeap) Less(i,j int) bool {
	return h[i] < h[j]
}
func (h smallHeap) Swap(i,j int) {
	h[i],h[j] = h[j],h[i]
}
func (h *smallHeap) Push(x interface{}) {
	*h = append(*h,x.(int))
}
func (h *smallHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}
func (h smallHeap) Top() interface{} {
	return h[0]	//顶是下标0
}

type MedianFinder struct {
	left *bigHeap
	right *smallHeap
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	left := &bigHeap{}
	heap.Init(left)
	right := &smallHeap{}
	heap.Init(right)
	return MedianFinder{left,right}
}


func (this *MedianFinder) AddNum(num int)  {
	if (len(*(this.left)) + len(*(this.right))) % 2 == 0 {//入left
		heap.Push(this.right,num)
		heap.Push(this.left,heap.Pop(this.right))
	}else{
		heap.Push(this.left,num)
		heap.Push(this.right,heap.Pop(this.left))
	}
}


func (this *MedianFinder) FindMedian() float64 {
	if (len(*(this.left)) + len(*(this.right))) % 2 == 0 {
		return (float64((*(this.left)).Top().(int)) + float64((*(this.right)).Top().(int))) / 2.0
	}else{
		return float64((*(this.left)).Top().(int))
	}
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

