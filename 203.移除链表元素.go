/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	hhead := new(ListNode)
	hhead.Next = head
	pre := hhead
	cur := head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			cur = cur.Next
		}else{
			pre,cur = cur,cur.Next
		}
	}
	return hhead.Next
}
// @lc code=end

