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
    /*
	var cur *ListNode  = head.Next
    var last *ListNode = head
	*/
    cur,last := head.Next,head
    for cur != nil {
		if cur.Val != last.Val{
			last.Next = cur
            last = cur
		}
		cur = cur.Next
	}
    last.Next = nil
	return head
}
// @lc code=end

