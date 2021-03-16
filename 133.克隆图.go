/*
 * @lc app=leetcode.cn id=133 lang=golang
 *
 * [133] 克隆图
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    visited := map[*Node]*Node{}
	var cg func(node *Node) *Node
	cg = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if copy,ok := visited[node] ; ok {
			return copy
		}
		copy := &Node{Val:node.Val}
		visited[node] = copy
		for _,neighbor := range node.Neighbors {
			copy.Neighbors = append(copy.Neighbors,cg(neighbor))
		}
		return copy
	}
	return cg(node)
}
// @lc code=end

