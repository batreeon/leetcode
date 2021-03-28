/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	}

	rootIndex := 0
	for ; inorder[rootIndex] != postorder[n-1] ; rootIndex++ {}

	root := &TreeNode{Val:postorder[n-1]}
	root.Left = buildTree(inorder[:rootIndex],postorder[:rootIndex])
	root.Right = buildTree(inorder[rootIndex+1:],postorder[rootIndex:n-1])

	return root
}
// @lc code=end

