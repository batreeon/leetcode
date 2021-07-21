/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
		return nil
	}
	pA,pB := headA,headB
	for pA != pB {
		/*
        // 不能这样写，当没有公共节点时，他们最终都指向nil，而这样写直接跳过nil了
		if pA.Next == nil {
			pA = headB
		}else{
			pA = pA.Next
		}
		if pB.Next == nil {
			pB = headA
		}else{
			pB = pB.Next
		}
        */
        if pA == nil {
			pA = headB
		}else{
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		}else{
			pB = pB.Next
		}
	}
	return pA
}