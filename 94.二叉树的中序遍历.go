/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	var ans []int
	var node []*TreeNode
	nowNode := root
	for ;; {
		for ; nowNode != nil ; {
			node = append(node,nowNode)
			nowNode = nowNode.Left
		}
		if len(node) == 0 {
			break
		}
		nowNode = node[len(node)-1]
		node = node[:len(node)-1]
		ans = append(ans,nowNode.Val)
		nowNode = nowNode.Right
	}
	return ans
}
// @lc code=end

