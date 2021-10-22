/*
 * @lc app=leetcode.cn id=218 lang=golang
 *
 * [218] 天际线问题
 */

// @lc code=start
package main
import "sort"

func getSkyline(buildings [][]int) [][]int {
	// 从左到右扫描边
	nodes := [][]int{}
	for _,b := range buildings {
		nodes = append(nodes,[]int{b[0],-b[2]})
		nodes = append(nodes,[]int{b[1],b[2]})
	}
	// 上一个最高的高度
	height := []int{0}
	// 上一个拐点
	last := []int{0,0}
	result := [][]int{}
	for _,node := range nodes {
		if node[1] < 0 {
			height = append(height,-node[1])
		}else{
			idx := sort.Search(len(height),func(i int) bool {return height[i] == node[1]})
			copy(height[idx:len(height)-1],height[idx+1:])
			height = height[:len(height)-1]
		}

		sort.Ints(height)
		maxH := height[len(height)-1]
		if maxH != last[1] {
			last[0],last[1] = node[0],maxH
			lastCopy := make([]int,2)
			lastCopy[0],lastCopy[1] = last[0],last[1]
			result = append(result,lastCopy)
		}
	}
	return result
}
// @lc code=end

