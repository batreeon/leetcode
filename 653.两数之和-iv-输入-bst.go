/*
 * @lc app=leetcode.cn id=653 lang=golang
 *
 * [653] 两数之和 IV - 输入 BST
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
// package main
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }
// var nums map[int]bool = map[int]bool{}
func findTarget(root *TreeNode, k int) bool {
	nums := map[int]bool{}
	return findtarget(root, k, &nums)
}

func findtarget(root *TreeNode, k int, nums *map[int]bool) bool {
	if root == nil {
		return false
	}
	if (*nums)[k - root.Val] {
		return true
	}
	(*nums)[root.Val] = true
	return findtarget(root.Left, k, nums) || findtarget(root.Right, k, nums)

}
// @lc code=end

