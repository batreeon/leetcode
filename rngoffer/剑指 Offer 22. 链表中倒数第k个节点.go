/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func getKthFromEnd(head *ListNode, k int) *ListNode {
	p1 := head
	for k > 0 {
		p1 = p1.Next
		k--
	}
	p2 := head
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}