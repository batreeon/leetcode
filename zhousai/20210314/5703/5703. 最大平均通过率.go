
/*
一所学校里有一些班级，每个班级里有一些学生，现在每个班都会进行一场期末考试。给你一个二维数组 classes ，其中 classes[i] = [passi, totali] ，表示你提前知道了第 i 个班级总共有 totali 个学生，其中只有 passi 个学生可以通过考试。

给你一个整数 extraStudents ，表示额外有 extraStudents 个聪明的学生，他们 一定 能通过任何班级的期末考。你需要给这 extraStudents 个学生每人都安排一个班级，使得 所有 班级的 平均 通过率 最大 。

一个班级的 通过率 等于这个班级通过考试的学生人数除以这个班级的总人数。平均通过率 是所有班级的通过率之和除以班级数目。

请你返回在安排这 extraStudents 个学生去对应班级后的 最大 平均通过率。与标准答案误差范围在 10-5 以内的结果都会视为正确结果。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-average-pass-ratio
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"container/heap"
)

type bigHeap [][]int
func (h bigHeap) Len() int {
	return len(h)
}
func (h bigHeap) Less(i,j int) bool {
	progess := func(i int) float64 {
		old := float64(h[i][0])/float64(h[i][1])
		new := float64(h[i][0]+1)/float64(h[i][1]+1)
		return new-old
	}
	return progess(i) > progess(j)
}
func (h bigHeap) Swap(i,j int) {
	h[i],h[j] = h[j],h[i]
}
func (h *bigHeap) Push(x interface{}) {
	*h = append(*h,x.([]int))
}
func (h *bigHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	h := &bigHeap{}
	heap.Init(h)
	n := len(classes)
	for i := 0 ; i < n ; i++ {
		heap.Push(h,classes[i])
	}
	for extraStudents > 0 {
		class := heap.Pop(h)
		class.([]int)[0]++
		class.([]int)[1]++
		heap.Push(h,class)
		extraStudents--
	}
	var sum float64
	for _,c := range *h {
		sum += (float64(c[0])/float64(c[1]))
	}
	return sum/float64(n)
}