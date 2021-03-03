/*
 * @lc app=leetcode.cn id=173 lang=golang
 *
 * [173] 二叉搜索树迭代器
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
type BSTIterator struct {
	stack []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
	stack := []*TreeNode{}
	if root != nil {
		stack = append(stack,root)
		for root.Left != nil {
			stack = append(stack,root.Left)
			root = root.Left
		}
	}
	return BSTIterator{stack}
}


func (this *BSTIterator) Next() int {
	l := len(this.stack)
	node := this.stack[l-1]
	ans := node.Val

	if node.Right != nil {
		node = node.Right
		this.stack[l-1] = node
		for node.Left != nil {
			node = node.Left
			this.stack = append(this.stack,node)
		}
	}else{
		this.stack = this.stack[:l-1]
	}
	return ans
}


func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

