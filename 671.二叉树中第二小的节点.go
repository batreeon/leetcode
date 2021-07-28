/*
 * @lc app=leetcode.cn id=671 lang=golang
 *
 * [671] 二叉树中第二小的节点
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
func findSecondMinimumValue(root *TreeNode) int {
	result := -1
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	var dfs func(r *TreeNode,rootV int)
	dfs = func(r *TreeNode,rootV int) {
		if r == nil {
			return
		}
		if r.Val > rootV {
			if result == -1 {
				result = r.Val
			}else{
				result = min(result,r.Val)
			}
			// 找到一个后面的肯定不会更小了
			// result == -1时就是递归分支中第一个遇到比根节点更小的
			// result != -1就是，其他递归分支里面已经遇到过比根节点更小的了，当前分支也遇到了一个比根节点更小的
			// 我们需要比较哪个分支里的更小
			return
		}
		dfs(r.Left,rootV)
		dfs(r.Right,rootV)
	}
	dfs(root.Left,root.Val)
	dfs(root.Right,root.Val)
	return result
}
// @lc code=end

