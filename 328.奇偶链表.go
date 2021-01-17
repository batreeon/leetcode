/*
 * @lc app=leetcode.cn id=328 lang=golang
 *
 * [328] 奇偶链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	evenHead := head.Next
	odd , even := head , head.Next
	for ; even != nil && even.Next != nil; {
		odd.Next = even.Next
		//由even.Next确保不会越界
		odd = odd.Next
		even.Next = odd.Next
		//循环条件even != nil确保even.Next可取
		//odd.Next是一定存在的，这一步不会越界，这是有循环条件even.Next != nil限制的
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
// @lc code=end

