/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	q := []*TreeNode{}
	q = append(q,root)
	result := []int{}
	for len(q) > 0 {
		result = append(result,q[0].Val)
		l := len(q)
		for i := 0 ; i < l ; i++ {
			if q[i].Right != nil {
				q = append(q,q[i].Right)
				// 不可以这样，你想的是找到了最右的，就可以直接退了
				// 但你不能保证在更深层不会用到左边的部分
				// 而那时你早早就在上层时就把左边的部分删掉了
				// break
			}
			if q[i].Left != nil {
				q = append(q,q[i].Left)
				// break
			}
		}
		q = q[l:]
	}
	return result
}
// @lc code=end

