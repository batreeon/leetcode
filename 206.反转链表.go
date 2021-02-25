/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var last *ListNode
	for head != nil {
		next := head.Next
		head.Next = last
		last = head
		head = next
	}
	return last
}
// @lc code=end

