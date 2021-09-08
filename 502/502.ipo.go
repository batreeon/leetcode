/*
 * @lc app=leetcode.cn id=502 lang=golang
 *
 * [502] IPO
 */

// @lc code=start
package main

import (
	"sort"
	"container/heap"
)

type maxHeap struct {sort.IntSlice}
func (h maxHeap) Less(i,j int) bool {return h.IntSlice[i] > h.IntSlice[j]}
func (h *maxHeap) Push(x interface{}) {h.IntSlice = append(h.IntSlice,x.(int))}
func (h *maxHeap) Pop() interface{} {
	n := len(h.IntSlice)
	x := h.IntSlice[n-1]
	h.IntSlice = h.IntSlice[:n-1]
	return x
}
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	indexs := make([]int,n)
	for i := range indexs {
		indexs[i] = i
	}
	sort.Slice(indexs,func(i,j int) bool {
		ii,jj := indexs[i],indexs[j]
		return capital[ii] < capital[jj]
	})
	h := &maxHeap{}
	heap.Init(h)
	index := 0
	for ; k > 0 ; k-- {
		for ; index < n && capital[indexs[index]] <= w ; index++ {
			heap.Push(h,profits[indexs[index]])
		}
		if h.Len() == 0 {
			break
		}
		w += heap.Pop(h).(int)
	}
	return w
}
/*
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	indexs := make([]int,len(profits))
	for i := range indexs {
		indexs[i] = i
	}
	for ; k > 0 ; k-- {
		sort.Slice(indexs, func(i,j int) bool {
			ii,jj := indexs[i],indexs[j]
			return (profits[ii] > profits[jj]) || ((profits[ii] == profits[jj]) && (capital[ii] < capital[jj]))
		})
		idx := func() int {
			for i := 0 ; i < len(indexs) ; i++ {
				ii := indexs[i]
				if capital[ii] <= w {
					return i
				}
			}
			return len(indexs)
		}()
		if idx == len(indexs) {
			break
		}
		w += profits[indexs[idx]]
		indexs = remove(indexs,idx)
	}
	return w
}
func remove(slice []int, i int) []int {
	copy(slice[i:],slice[i+1:])
	return slice[:len(slice)-1]
}
*/

/*
type _sort struct{
	sort.IntSlice	//capital
	bind []int		//profits
}
func (s _sort) Less(i,j int) bool {
	return (s.bind[i] > s.bind[j]) || ((s.bind[i] == s.bind[j]) && (s.IntSlice[i] < s.IntSlice[j]))	
}
func (s _sort) Swap(i,j int) {
	s.IntSlice[i],s.IntSlice[j] = s.IntSlice[j],s.IntSlice[i]
	s.bind[i],s.bind[j] = s.bind[j],s.bind[i]
}
func (s _sort) Remove(i int) {
	s.IntSlice = remove(s.IntSlice,i)
	s.bind = remove(s.bind,i)
}
func remove(slice []int, i int) []int {
	copy(slice[i:],slice[i+1:])
	return slice[:len(slice)-1]
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	_Sort := _sort{
		capital,
		profits,
	}

	for ; k > 0 ; k-- {
		_Sort.Sort()
		idx := sort.Search(_Sort.Len(),func(i int) bool {
			if i == 0 && _Sort.IntSlice[i] <= w {
				return true
			}
			if i > 0 && _Sort.IntSlice[i-1] > w && _Sort.IntSlice[i] <= w {
				return true
			}
			return false
		})
		if idx == len(capital) {
			break
		}
		w += _Sort.bind[idx] - _Sort.IntSlice[idx]
		_Sort.Remove(idx)
	}
	return w
}
*/
// @lc code=end

