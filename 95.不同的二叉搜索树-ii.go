/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
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
func generateTrees(n int) []*TreeNode {
	var generatetrees func(int,int) []*TreeNode
	generatetrees = func(start int,end int) []*TreeNode {
		if start > end {
			// 不能直接返回nil，要返回列表，不然无法进入遍历，最终返回结果为空
			return []*TreeNode{nil}
		}
		allTrees := []*TreeNode{}
		for i := start ; i <= end ; i++ {
			left := generatetrees(start,i-1)
			right := generatetrees(i+1,end)
			for _,l := range left {
				for _,r := range right {
					root := &TreeNode{i,nil,nil}
					root.Left = l
					root.Right = r
					allTrees = append(allTrees,root)
				}
			}
		}
		return allTrees
	}
	if n == 0 {
		return nil
	}
	return generatetrees(1,n)
}
// @lc code=end

