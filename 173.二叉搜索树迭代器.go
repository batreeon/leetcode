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

 /*
//  以前通过版
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
*/


// 20210328每日一题版
type BSTIterator struct {
	nodes []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	nodes := []*TreeNode{}
	if root != nil {
		nodes = append(nodes,root)
		for root.Left != nil {
			nodes = append(nodes,root.Left)
			root = root.Left
		}
	}
	return BSTIterator{nodes}
}

func (this *BSTIterator) Next() int {
	l := len(this.nodes)
	node := this.nodes[l-1]
	val := node.Val
	
	this.nodes = this.nodes[:l-1]

	if node.Right != nil {
		this.nodes = append(this.nodes,node.Right)
		node = node.Right
		for node.Left != nil {
			this.nodes = append(this.nodes,node.Left)
			node = node.Left
		}
	}
	return val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.nodes) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

