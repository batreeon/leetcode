package main

import "math"
/*
有一个无向的 星型 图，由 n 个编号从 1 到 n 的节点组成。星型图有一个 中心 节点，并且恰有 n - 1 条边将中心节点与其他每个节点连接起来。

给你一个二维整数数组 edges ，其中 edges[i] = [ui, vi] 表示在节点 ui 和 vi 之间存在一条边。请你找出并返回 edges 所表示星型图的中心节点。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-center-of-star-graph
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func findCenter(edges [][]int) int {
	m := map[int]int{}
	for _, edge := range edges {
		for _, i := range edge {
			m[i]++
		}
	}
	maxTimes := math.MinInt64
	ans := -1
	for k,v := range m {
		if v > maxTimes {
			maxTimes = v
			ans = k
		}
	}
	return ans
}