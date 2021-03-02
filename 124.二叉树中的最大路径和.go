/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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
func maxPathSum(root *TreeNode) int {

	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	maxPath := math.MinInt64
	var maxpathsum func(root *TreeNode) int
	maxpathsum = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		// left := maxpathsum(root.Left)
		// right := maxpathsum(root.Right)
		// maxPath = max(maxPath,root.Val + max(0,left) + max(0,right))
		// return root.Val + max(left,right)	//注意这里，子树的某一条路径小于0，那还不如不加呢
		left := max(0,maxpathsum(root.Left))
		right := max(0,maxpathsum(root.Right))
		maxPath = max(maxPath,root.Val + left + right)
		return root.Val + max(left,right)	//返回的值不分叉的最大路径值
	}
	maxpathsum(root)
	return maxPath
	
}
// @lc code=end

