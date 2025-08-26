/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 package main
 
// type ListNode struct {
//     Val int
//     Next *ListNode
//  }

func reverseKGroup(head *ListNode, k int) *ListNode {
    n := 0
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}

	dummy := &ListNode{Next: head}
	p0 := dummy

	pre := new(ListNode)
	cur = p0.Next
	for n >= k {
		n -= k 
		for i := 0; i < k; i++ {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}

		next := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = next
	}

	return dummy.Next
}
// @lc code=end

