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

