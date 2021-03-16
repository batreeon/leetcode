/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	var sp func(head *ListNode) *ListNode
	sp = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil{
			return head
		}
		/*
		p := head
		pp := head.Next
		p.Next = pp.Next
		pp.Next = p
		head = pp
		*/
		p := head.Next
		head.Next = p.Next
		p.Next = head
		head = p
		head.Next.Next = sp(head.Next.Next)
		return head
	}
	return sp(head)
}
// @lc code=end

