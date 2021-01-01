/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
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
var paths []string
func binaryTreePaths(root *TreeNode) []string {
	paths = []string{}
	constructPath(root , "")
	return paths
}
func constructPath(root *TreeNode , path string) {
	if root != nil {
		pathNow := path
		pathNow += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			paths = append(paths,pathNow)
		} else {
			constructPath(root.Left,pathNow + "->")
			constructPath(root.Right,pathNow + "->")
		}
	}
}
// @lc code=end

