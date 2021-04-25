/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nodes := []*TreeNode{root}
	result := 0
	for len(nodes) > 0 {
		l := len(nodes)
		result += l
		for _,node := range nodes {
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
// @lc code=end

