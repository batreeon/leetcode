/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
package main
import (
	"container/heap"
)
type MaxHeap []int
func (h MaxHeap) Len() int {return len(h)}
func (h MaxHeap) Less(i,j int) bool {return h[i] > h[j]}//大顶堆
func (h MaxHeap) Swap(i,j int) {h[i],h[j] = h[j],h[i]}
func (h *MaxHeap) Push(x interface{}) {*h = append(*h,x.(int))}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
type MinHeap []int
func (h MinHeap) Len() int {return len(h)}
func (h MinHeap) Less(i,j int) bool {return h[i] < h[j]}
func (h MinHeap) Swap(i,j int) {h[i],h[j] = h[j],h[i]}
func (h *MinHeap) Push(x interface{}) {*h = append(*h,x.(int))}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	big := &MaxHeap{}
	heap.Init(big)
	small := &MinHeap{}
	//想法是这样的，但是这样写不对 small := heap.interface(sort.Reverse(&MaxHeap{}))
	heap.Init(small)
	l := len(nums1)+len(nums2)
	for _,v := range nums1 {
		// if len(*big)+len(*small) > l/2 {
		// 	break
		// }//不能写这个，这个不是一个输入流，是两个啊，不能这样写
		if (len(*big)+len(*small)) & 1 == 1 {
			heap.Push(big,v)
			heap.Push(small,heap.Pop(big))
		}else{
			heap.Push(small,v)
			heap.Push(big,heap.Pop(small))
		}
	}
	for _,v := range nums2 {
		// if len(*big)+len(*small) > l/2 {
		// 	break
		// }
		if (len(*big)+len(*small)) & 1 == 1 {
			heap.Push(big,v)
			heap.Push(small,heap.Pop(big))
		}else{
			heap.Push(small,v)
			heap.Push(big,heap.Pop(small))
		}
	}
	var ans float64
	if l & 1 == 1 {
		ans = float64(heap.Pop(big).(int))
	}else{
		ans = (float64(heap.Pop(big).(int)) + float64(heap.Pop(small).(int))) / 2.0
	}
	return ans
}
//上面这个使用了大小堆，内存消耗比较大啊
//如果len(nums1)+len(nums2)是奇数，可以用快排，当分界点下标是n/2时，就找到了中位数
// @lc code=end