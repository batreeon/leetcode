/*
 * @lc app=leetcode.cn id=1382 lang=golang
 *
 * [1382] 将二叉搜索树变平衡
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
func balanceBST(root *TreeNode) *TreeNode {
	treeNodes := []*TreeNode{}
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root != nil {
			inorder(root.Left)
			treeNodes = append(treeNodes,root)
			inorder(root.Right)
			root.Left = nil
			root.Right = nil
		}
	}
	var build func(l,r int) *TreeNode
	build = func(l,r int) *TreeNode {
		mid := (l+r) >> 1
		root := treeNodes[mid]
		if l < mid {
			root.Left = build(l,mid-1)
		}
		if mid < r {
			root.Right = build(mid+1,r)
		}
		return root
	}

	// 啊啊，还是得用数组存着啊
	inorder(root)
	return build(0,len(treeNodes)-1)
}
// @lc code=end

