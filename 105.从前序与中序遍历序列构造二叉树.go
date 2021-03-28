/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootIndex := 0
	for ; preorder[0] != inorder[rootIndex] ; rootIndex++ {}
	root := &TreeNode{Val:preorder[0]}
	root.Left = buildTree(preorder[1:rootIndex+1],inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:],inorder[rootIndex+1:])

	return root
}
// @lc code=end

