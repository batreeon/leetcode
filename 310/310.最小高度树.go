/*
 * @lc app=leetcode.cn id=310 lang=golang
 *
 * [310] 最小高度树
 */

// @lc code=start
package main
func findMinHeightTrees(n int, edges [][]int) []int {
	// 若没有边，则所有节点都可以作为根
	if len(edges) == 0 {
		result := []int{}
		for i := 0 ; i < n ; i++ {
			result = append(result,i)
		}
		return result
	}

	// 思路，每一遍删除所有的叶子节点，即只有一个（或0个）邻节点的节点，（若最后只剩下一个节点，那么其邻接节点数就为0）
	// 并将其从其邻节点的邻接表中删除，这也是我们为什么把cnt的value值设为map的原因，为了方便删除其邻接点
	// 若经过最后依次删除后，cnt中为空了，那么最后依次被删除的节点就是答案，因为nodes会在上述判断之后才会进行更新，
	// 因此nodes可用于恢复最后一轮剩余的节点
	// （若最后只剩下一个节点，其没有邻节点）

	// cnt用来记录每个节点的邻节点
	cnt := map[int]map[int]bool{}
	// nodes存放依然存在的点
	nodes := map[int]bool{}
	// 相当于建邻接表
	for _,edge := range edges {
		if cnt[edge[0]] == nil {
			cnt[edge[0]] = map[int]bool{}
		}
		if cnt[edge[1]] == nil {
			cnt[edge[1]] = map[int]bool{}
		}
		cnt[edge[0]][edge[1]] = true
		cnt[edge[1]][edge[0]] = true
		nodes[edge[0]] = true
		nodes[edge[1]] = true
	}

	for {
		// if len(cnt) == 1 {
		// 	for k,_ := range cnt {
		// 		return []int{k}
		// 	}
		// }
		// 记录要删除的叶子节点
		del := []int{}
		for k,v := range cnt {
			if len(v) == 1 || len(v) == 0 {
				del = append(del,k)
			}
		}
		
		// 删除节点并更新其他节点的邻接表
		for _,delnode := range del {
			for kk,_ := range cnt[delnode] {
				delete(cnt[kk],delnode)
			}
			delete(cnt,delnode)
		}

		// 若最后没有节点了，取出最后一轮删除的节点，返回结果
		if len(cnt) == 0 {
			result := []int{}
			for k,_ := range nodes {
				result = append(result,k)
			}
			return result
		}
		// 更新剩余节点集合
		for _,delnode := range del {
			delete(nodes,delnode)
		}
	}
	return []int{}
}
// @lc code=end

