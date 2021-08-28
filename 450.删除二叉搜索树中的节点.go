/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] 删除二叉搜索树中的节点
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
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right,key)
	}else if key < root.Val {
		root.Left = deleteNode(root.Left,key)
	}else{
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 上面这两句别忘了

		pNode := root.Right
		for pNode.Left != nil {
			pNode = pNode.Left
		}
		pNode.Left = root.Left
		return root.Right
	}
	return root
}
// @lc code=end

