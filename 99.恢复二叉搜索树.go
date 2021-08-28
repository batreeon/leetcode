/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
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
func recoverTree(root *TreeNode)  {
	nums := []int{}
	inorder(root,&nums)

	nts := needToSwap(nums)
	recover(root,nts,2)
}
// slice不是传的指针吗，为什么这个函数参数nums要用切片指针才行呢
func inorder(root *TreeNode,nums *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left,nums)
	*nums = append(*nums,root.Val)
	inorder(root.Right,nums)
	return
}
func needToSwap(nums []int) ([]int) {
	// 这个太妙了
	x,y := -1,-1
	for i := 0 ; i < len(nums)-1 ; i++ {
		if nums[i+1] < nums[i] {
			y = nums[i+1]
			if x == -1 {
				x = nums[i]
			}else{
				break
			}
		}
	}
	return []int{x,y}
}
func recover(root *TreeNode,nts []int,count int) {
	if root == nil {
		return
	}
	if root.Val == nts[0] || root.Val == nts[1] {
		root.Val ^= nts[0]
		root.Val ^= nts[1]
		count--
		if count == 0 {
			return
		}
	}
	recover(root.Left,nts,count)
	recover(root.Right,nts,count)
}
// @lc code=end

