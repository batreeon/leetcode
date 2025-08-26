/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
func isSymmetric(root *TreeNode) bool {
	var f func(*TreeNode, *TreeNode) bool
	f = func(l,r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}
		if l == nil || r == nil {
			return false
		}
		lv, rv := l.Val, r.Val
		if lv != rv {
			return false
		}
		return f(l.Left, r.Right) && f(l.Right, r.Left)
	}
	if root == nil {
		return true
	}
	return f(root.Left, root.Right)
}
// @lc code=end

