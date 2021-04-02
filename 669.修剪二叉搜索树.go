/*
 * @lc app=leetcode.cn id=669 lang=golang
 *
 * [669] 修剪二叉搜索树
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
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	var trimbst func(root *TreeNode) *TreeNode
	trimbst = func(root *TreeNode) *TreeNode {
		if root == nil {
			return root
		}
		if root.Val < low {
			return trimbst(root.Right)
		}else if root.Val > high {
			return trimbst(root.Left)
		}else{
			root.Left = trimbst(root.Left)
			root.Right = trimbst(root.Right)
		}
		return root
	}
	return trimbst(root)
}
// @lc code=end

