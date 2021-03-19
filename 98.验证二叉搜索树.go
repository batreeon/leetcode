/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
package main
// import "sort"

func isValidBST(root *TreeNode) bool {
	/*
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	if root.Left == nil {
		if root.Val >= root.Right.Val {
			return false
		}
		return isValidBST(root.Right)
	}else if root.Right == nil {
		if root.Val <= root.Left.Val {
			return false
		}
		return isValidBST(root.Left)
	}else{
		if root.Val >= root.Right.Val || root.Val <= root.Left.Val {
			return false
		}
		return isValidBST(root.Left) && isValidBST(root.Right)
	}
	return false
	*/

	/*
	var iVBLeft func(root *TreeNode,key int) bool
	var iVBRight func(root *TreeNode,key int) bool
	iVBLeft = func(root *TreeNode,key int) bool {
		if root == nil {
			return true
		}
		if root.Val >= key {
			return false
		}
		if root.Left == nil && root.Right == nil {
			return true
		}
		if root.Left == nil {
			if root.Val >= root.Right.Val {
				return false
			}
			return iVBRight(root.Right,key)
		}else if root.Right == nil {
			if root.Val <= root.Left.Val {
				return false
			}
			return iVBLeft(root.Left,key)
		}else{
			if root.Val >= root.Right.Val || root.Val <= root.Left.Val {
				return false
			}
			return iVBLeft(root.Left,key) && iVBRight(root.Right,key)
		}
	}
	iVBRight = func(root *TreeNode,key int) bool {
		if root == nil {
			return true
		}
		if root.Val <= key {
			return false
		}
		if root.Left == nil && root.Right == nil {
			return true
		}
		if root.Left == nil {
			if root.Val >= root.Right.Val {
				return false
			}
			return iVBRight(root.Right,key)
		}else if root.Right == nil {
			if root.Val <= root.Left.Val {
				return false
			}
			return iVBLeft(root.Left,key)
		}else{
			if root.Val >= root.Right.Val || root.Val <= root.Left.Val {
				return false
			}
			return iVBLeft(root.Left,key) && iVBRight(root.Right,key)
		}
	}
	if root == nil {
		return true
	}
	return iVBLeft(root.Left,root.Val) && iVBRight(root.Right,root.Val)
	*/

	// 上面两种都有漏洞
	// 不会遍历解法啊
	var nums []int
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		nums = append(nums,root.Val)
		inorder(root.Right)
	}
	inorder(root)
	// return sort.IsSorted(nums)
	n := len(nums)
	for i := 1 ; i < n ; i++ {
		if nums[i-1] >= nums[i] {
			return false
		}
	}
	return true
}
// @lc code=end

