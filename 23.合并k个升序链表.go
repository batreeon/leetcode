/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 //自定义可排序类型实现sort接口，不行（多次排序，超时）！ 再进一步实现小顶堆，行！
import "container/heap"
type status struct {
	v int
	p *ListNode
}
type p_q []status
func (q p_q) Len() int {
	return len(q)
}
func (q p_q) Less(i,j int) bool {
	return q[i].v < q[j].v
}
func (q p_q) Swap(i,j int) {
	q[i],q[j] = q[j],q[i]
}

func (q *p_q) Push(x interface{}) {
    *q = append(*q, x.(status))
}
func (q *p_q) Pop() interface{} {
    old := *q
    n := len(old)
    x := old[n-1]
    *q = old[0 : n-1]
    return x
}
func mergeKLists(lists []*ListNode) *ListNode {
	

	q := &p_q{}
	heap.Init(q)
	for _,node := range lists {
		if node != nil {
			heap.Push(q,status{node.Val,node})
		}
	}
	head := &ListNode{}
	tail := head
	for len(*q) > 0 {
		next := heap.Pop(q).(status)
		tail.Next = next.p
		tail = tail.Next
		if next.p.Next != nil {
			heap.Push(q,status{next.p.Next.Val,next.p.Next})
		}
	}
	return head.Next
}
// @lc code=end

