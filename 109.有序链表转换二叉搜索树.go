/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	// 处理单节点的下面这个判断，是必须要加的，若只有单节点slow会指向head，
	// 左子树又会将head传入函数，无线循环，耗尽内存
	// （传入两节点链表，slow会指向第二个没这个问题，右子树会给函数传入nil,没有这个问题）
	if head.Next == nil {
		return &TreeNode{Val:head.Val}
	}
	pre := &ListNode{Next:head}
	slow,fast := head,head
	for fast.Next != nil {
		pre = pre.Next
		slow = slow.Next
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		}
	}
	pre.Next = nil
	root := &TreeNode{Val:slow.Val}
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(slow.Next)

	return root
}
// @lc code=end

