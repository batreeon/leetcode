package lru

type ListNode struct {
	K int
	V int
	Pre *ListNode
	Next *ListNode
}

type lru struct {
	n int
	m map[int]*ListNode
	dummy *ListNode
}

func NewLru(n int) lru {
	dummy := &ListNode{}
	dummy.Pre = dummy
	dummy.Next = dummy

	return lru{
		n: n,
		m: make(map[int]*ListNode),
		dummy: dummy,
	}
}

func (lru *lru) Get(k int) int {
	listNode, ok := lru.m[k]
	if !ok {
		return -1
	}

	lru.moveToFront(listNode)

	return listNode.V
}

func (lru *lru) Set(k, v int) {
	listNode, ok := lru.m[k]
	if ok {
		listNode.V = v
		lru.moveToFront(listNode)
		return
	}

	if len(lru.m) == lru.n {
		tailNode := lru.dummy.Pre
		lru.remove(tailNode)
		delete(lru.m, tailNode.K)
	}

	newNode := &ListNode{
		K: k,
		V: v,
	}
	lru.insertToFront(newNode)
	lru.m[k] = newNode
}


func (lru *lru) moveToFront(ln *ListNode) {
	lru.remove(ln)

	lru.insertToFront(ln)
}

func (lru *lru) insertToFront(ln *ListNode) {
	ln.Pre = lru.dummy
	ln.Next = lru.dummy.Next
	ln.Pre.Next = ln
	ln.Next.Pre = ln
}

func (lru *lru) remove(ln *ListNode) {
	ln.Pre.Next = ln.Next
	ln.Next.Pre = ln.Pre
}