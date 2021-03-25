/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	hhead := new(ListNode)
	hhead.Next = head
	pre := hhead
	p := head
	
	for p != nil && p.Next != nil {
		if p.Val == p.Next.Val {
			for p.Val == p.Next.Val {
				p = p.Next
				if p.Next == nil {
					break
				}
			}
			pre.Next = p.Next
		}else{
			pre.Next = p
			pre = pre.Next
		}
		p = p.Next
	}
	return hhead.Next
}
// @lc code=end

