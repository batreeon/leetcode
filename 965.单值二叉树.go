/*
 * @lc app=leetcode.cn id=965 lang=golang
 *
 * [965] 单值二叉树
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
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Val != root.Left.Val {
		return false
	}
	if root.Right != nil && root.Val != root.Right.Val {
		return false
	}
	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}
// @lc code=end

