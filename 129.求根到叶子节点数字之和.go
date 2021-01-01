/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根到叶子节点数字之和
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
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Right == nil && root.Left == nil {
		return root.Val
	}
	if root.Left != nil {
		root.Left.Val += 10*root.Val
	}
	if root.Right != nil {
		root.Right.Val += 10*root.Val
	}
	return sumNumbers(root.Left) + sumNumbers(root.Right)
}
// @lc code=end

