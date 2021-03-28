/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode) {
	var f func(root *TreeNode) *TreeNode
	f = func(root *TreeNode) *TreeNode {
		if root == nil || (root.Left == nil && root.Right == nil) {
			return root
		}
		right := root.Right
		root.Right = f(root.Left)
		pTreeNode := root
		for pTreeNode.Right != nil {
			pTreeNode = pTreeNode.Right
		}
		pTreeNode.Right = f(right)
		root.Left = nil
		return root
	}
	root = f(root)
}

// @lc code=end

