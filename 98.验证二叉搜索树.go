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

	/*
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
	*/

	// 上面是只会根据中序遍历结果看是否有序。。。
	return isvalidBST(root).valid

}

type Result struct{
	max int
	min int
	valid bool
}
func isvalidBST(root *TreeNode) (result Result) {
	// 一开始看这段代码时我疑问，result.valid什么时候会置false呢
	// 其实返回参数已经初始化result了，它初始状态就是false,
	// 只有我们判断root的两棵子树都符合二叉搜索树，并且root.Val，left.max,right.min大小满足要求时才会将result.valid置为true
	// 或者root==nil，会将result.valid置为true
	if root == nil {
		// 空树自然是合法的二叉搜索树
		// 并且以空树为子树也必定是合法的，因此该子树max为最小值，min为最大值 
		// 使得以空树为子树的父节点值符合大于左子树的最大值，小于右子树的最小值
		result.max = -1<<63
		result.min = 1<<63-1
		result.valid = true
		return
	}

	left,right := isvalidBST(root.Left),isvalidBST(root.Right)
	if left.max < root.Val && root.Val < right.min && left.valid && right.valid {
		result.valid = true
	}
	result.max = max(root.Val,max(left.max,right.max))
	result.min = min(root.Val,min(left.min,right.min))
	return
}
func max(x,y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x,y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

