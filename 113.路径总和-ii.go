/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	// 暴力法吧，尝试遍历了所有可能

	ans := [][]int{}
	var pathsum func(root *TreeNode,target int,res []int) 
	pathsum = func(root *TreeNode,target int,res []int) {
		if root == nil {
			return
		}
		val := root.Val
		if root.Left == nil && root.Right  == nil && root.Val == target {
			resCopy := make([]int,len(res)+1)
			copy(resCopy,res)
			resCopy[len(res)] = val
			ans = append(ans,resCopy)
		}
		res = append(res,val)
		target -= val
		if root.Left != nil {
			pathsum(root.Left,target,res)
		}
		if root.Right != nil {
			pathsum(root.Right,target,res)
		}
	}

	pathsum(root,targetSum,[]int{})
	return ans
}
// @lc code=end

