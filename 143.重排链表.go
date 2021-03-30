/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
	/*
	array1 := []*ListNode{}
	for p := head ; p != nil ; {
		array1 = append(array1,p)
		p = p.Next
	}
	n := len(array1)
	array2 := make([]*ListNode,n)
	for i := 1 ; i < n ; i++ {
		if i&1 == 1 {
			array2[i] = array1[n-1-i/2]
		}else{
			array2[i] = array1[i/2]
		}
	}
	pNode := head
	for i := 1 ; i < n ; i++ {
		pNode.Next = array2[i]
		pNode = pNode.Next
	}
	pNode.Next = nil
	*/

	array := []*ListNode{}
	for p := head ; p != nil ; {
		array = append(array,p)
		p = p.Next
	}
	n := len(array)
	pNode := head
	for i := 1 ; i < n ; i++ {
		if i&1 == 1 {
			pNode.Next = array[n-1-i/2]
		}else{
			pNode.Next = array[i/2]
		}
		pNode = pNode.Next
	}
	pNode.Next = nil
}
// @lc code=end

