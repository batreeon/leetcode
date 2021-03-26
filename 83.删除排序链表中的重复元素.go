/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
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
	if head == nil {
		return head
	}
	pre := head
	
	// 只有非空指针才能调用Next
	cur := head.Next
	for cur != nil {
		if pre.Val != cur.Val {
			pre.Next = cur
			pre = cur
		}
		cur = cur.Next
	}
	pre.Next = nil
	return head
}
// @lc code=end

