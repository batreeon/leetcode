/*
 * @lc app=leetcode.cn id=938 lang=golang
 *
 * [938] 二叉搜索树的范围和
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
func rangeSumBST(root *TreeNode, low int, high int) int {
	result := 0
	var rangesumbst func(root *TreeNode)
	rangesumbst = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Val >= low && root.Val <= high {
			result += root.Val
			rangesumbst(root.Left)
			rangesumbst(root.Right)
		}else if root.Val < low {
			rangesumbst(root.Right)
		}else{
			rangesumbst(root.Left)
		}
	}

	rangesumbst(root)
	return result
}
// @lc code=end

