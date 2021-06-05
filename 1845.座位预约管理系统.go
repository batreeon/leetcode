/*
 * @lc app=leetcode.cn id=1845 lang=golang
 *
 * [1845] 座位预约管理系统
 */

// @lc code=start
package main

import "container/heap"

type smallheap []int

func (h smallheap) Len() int           { return len(h) }
func (h smallheap) Less(i, j int) bool { return h[i] < h[j] }
func (h smallheap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *smallheap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *smallheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type SeatManager struct {
	available *smallheap
}

func Constructor(n int) SeatManager {
	a := &smallheap{}
	heap.Init(a)
	for i := 1; i <= n; i++ {
		heap.Push(a, i)
	}
	return SeatManager{available: a}
}

func (this *SeatManager) Reserve() int {
	return heap.Pop(this.available).(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(this.available, seatNumber)
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
// @lc code=end
