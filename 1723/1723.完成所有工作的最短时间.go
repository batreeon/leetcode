/*
 * @lc app=leetcode.cn id=1723 lang=golang
 *
 * [1723] 完成所有工作的最短时间
 */

// @lc code=start
package main

import (
	"sort"
)


func minimumTimeRequired(jobs []int, k int) int {
	sumjobs := 0.0
	for i := range jobs {
		sumjobs += float64(jobs[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(jobs)))
	// 上界是所有工作时间的总和
	high := int(sumjobs)
	// 下届是单个工作的最大时长
	low := jobs[0]
	for low < high {
		mid := (low+high)/2
		if !can(jobs,k,mid) {
			low = mid+1
		}else{
			high = mid
		}
	}
	return low
}
func can(jobs []int,k int,t int) bool {
	visited := make([]bool,len(jobs))
	kk := 1
	visitedNum := 0
	for visitedNum != len(visited) {
		temp := 0
		for i := range jobs {
			if visited[i] {
				continue
			}
			if temp + jobs[i] <= t {
				visited[i] = true
				visitedNum++
				temp += jobs[i]
			}
		}
		kk++
	}
	// 返回的kk，还没有使用，因为kk++是在尾部进行的
	// 返回kk，但其实只用了kk-1个工人
	// return kk-1 <= k
	if kk-1 > k {
		// t过小
		return false
	}
	// t够用
	return true
}
// @lc code=end

