package main

import (
	"container/heap"
	"sort"
    )


type cpu [][]int

func (c cpu) Len() int { return len(c) }
func (c cpu) Less(i, j int) bool {return c[i][0] < c[j][0]}
func (c cpu) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

type wait [][]int

func (w wait) Len() int {return len(w)}
func (w wait) Less(i,j int) bool {return (w[i][1] < w[j][1]) || (w[i][1]==w[j][1] && w[i][2] < w[j][2])}
func (w wait) Swap(i,j int) {w[i],w[j] = w[j],w[i]}
func (w *wait) Push(x interface{}) {
    *w = append(*w, x.([]int))
}
func (w *wait) Pop() interface{} {
    old := *w
    n := len(old)
    x := old[n-1]
    *w = old[0: n-1]
    return x
}

func getOrder(tasks [][]int) []int {
	for i, _ := range tasks {
		tasks[i] = append(tasks[i], i)
	}

	cpu := cpu(tasks)
	sort.Sort(cpu)
    
    wait := &wait{}
    heap.Init(wait)
    
    // 注意每次更新nowtime时直接让其等于第一个任务的开始时间就行了
    // 不要递增，无意义地浪费时间
    nowtime := cpu[0][0]
	result := []int{}
    
    for len(cpu) != 0 {
        i := 0
        for ; i < len(cpu) ; i++ {
            if cpu[i][0] <= nowtime {
                heap.Push(wait,cpu[i])
            }else{
				break
			}
        }
        cpu = cpu[i:]
        if wait.Len() > 0 {
            waittop := heap.Pop(wait)
            result = append(result,waittop.([]int)[2])
            nowtime = nowtime + waittop.([]int)[1]
        }else{
            nowtime = cpu[0][0]
        }
    }
    for wait.Len() != 0 {
        waittop := heap.Pop(wait)
        result = append(result,waittop.([]int)[2])
    }

	return result
}

/*
// 用堆能过，用切片加排序过不了
// 还是得用堆，没有新的任务加进来就排序，这太慢了
package main

import "sort"

type cpu [][]int

func (c cpu) Len() int { return len(c) }
func (c cpu) Less(i, j int) bool {return c[i][0] < c[j][0]}
func (c cpu) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

type wait [][]int

func (w wait) Len() int {return len(w)}
func (w wait) Less(i,j int) bool {return (w[i][1] < w[j][1]) || (w[i][1]==w[j][1] && w[i][2] < w[j][2])}
func (w wait) Swap(i,j int) {w[i],w[j] = w[j],w[i]}

func getOrder(tasks [][]int) []int {
	for i, _ := range tasks {
		tasks[i] = append(tasks[i], i)
	}
    
	cpu := cpu(tasks)
	sort.Sort(cpu)
    
	result := []int{}
    // if cpu[0][0] == cpu[len(cpu)-1][0] {
    //     wait := wait(cpu)
    //     sort.Sort(wait)
    //     for _,wait := range wait {
    //         result = append(result,wait[2])
    //     }
    //     return result
    // }
    // if cpu[0][1] == cpu[len(cpu)-1][1] {
    //     wait := wait(cpu)
    //     sort.Sort(wait)
    //     for _,wait := range wait {
    //         result = append(result,wait[2])
    //     }
    //     return result
    // }
    wait := wait{}
    
    nowtime := cpu[0][0]
    
    for len(cpu) != 0 {
        i := 0
        addflag := false
        for ; i < len(cpu) ; i++ {
            if cpu[i][0] <= nowtime {
                wait = append(wait,cpu[i])
                addflag = true
            }else{
				break
			}
        }
        cpu = cpu[i:]
        if len(wait) > 0 {
            if addflag {
                sort.Sort(wait)
            }
            result = append(result,wait[0][2])
            nowtime = nowtime + wait[0][1]
            wait = wait[1:]
        }else{
            nowtime = cpu[0][0]
        }
    }
    if len(wait) != 0 {
        for _,w := range wait {
            result = append(result,w[2])
        }
    }
	return result
}
*/