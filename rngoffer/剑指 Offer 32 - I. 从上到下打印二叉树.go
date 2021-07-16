/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrder(root *TreeNode) []int {
    if root == nil {
		return []int{}
	}
	result := []int{}
	nodes := []*TreeNode{}
	nodes = append(nodes,root)
	for len(nodes) > 0 {
		l := len(nodes)
		for _,node := range nodes {
			result = append(result,node.Val)
			if node.Left != nil {
				nodes = append(nodes,node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes,node.Right)
			}
		}
		nodes = nodes[l:]
	}
	return result
}