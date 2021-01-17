/*
 * @lc app=leetcode.cn id=1024 lang=golang
 *
 * [1024] 视频拼接
 */

// @lc code=start
// package main

// import (
// 	"fmt"
// 	"sort"
// )

type by1dint [][]int

func (p by1dint) Len() int {
	return len(p)
}
func (p by1dint) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p by1dint) Less(i, j int) bool {
	return (p[i][0] < p[j][0]) || ((p[i][0] == p[j][0]) && (p[i][1] < p[j][1]))
	//开始时间短就排前面，开始时间相等时，按结束时间排
}

func videoStitching(clips [][]int, T int) int {
	sort.Sort(by1dint(clips))
	// fmt.Println(clips, len(clips))
	var i int
	num := 1
	for i = 1; clips[i][0] == clips[0][0]; i++ {
	}
	tail := clips[i-1][1]
	maxend := tail
	now2 := clips[i][1]
	i++
	// fmt.Println("i", "num", "now2", "tail")
	for ; i < len(clips); i++ {
		if clips[i][0] > maxend {
			return -1
		}
		if clips[i][1] > maxend {
			maxend = clips[i][1]
		}
		if clips[i][0] <= tail {
			if clips[i][1] > now2 {
				now2 = clips[i][1]
			}
		} else {
			num++
			tail = now2
		}
		// fmt.Println(i, num, now2, tail, maxend)
	}
	if maxend < T {
		return -1
	}
	if i==len(clips) && maxend == T && clips[i-1][0] <= tail{
		num++
	}
	return num
}
// func main() {
// 	// fmt.Println(videoStitching([][]int{{0,2}, {4,6}, {8,10}, {1,9}, {1,5}, {5,9}}, 10))
// 	fmt.Println(videoStitching([][]int{{0, 1}, {0, 2}, {5, 6}, {0, 4}, {0, 3}, {1, 3}, {4, 7}, {1, 4}, {2, 5}, {2, 6}, {3, 4}, {4, 5}, {5, 7}, {6, 9}}, 9))
// }

// @lc code=end
