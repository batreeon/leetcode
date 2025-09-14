package lru

type ListNode struct {
	K int
	V int
	Pre *ListNode
	Next *ListNode
}

type LRU struct {
	n int
	m map[int]*ListNode
	dummy *ListNode
}

func NewLru(n int) LRU {
	dummy := &ListNode{}
	dummy.Pre = dummy
	dummy.Next = dummy

	return LRU{
		n: n,
		m: make(map[int]*ListNode),
		dummy: dummy,
	}
}

func (lru *LRU) Get(k int) int {
	listNode, ok := lru.m[k]
	if !ok {
		return -1
	}

	lru.moveToFront(listNode)

	return listNode.V
}

func (lru *LRU) Set(k, v int) {
	listNode, ok := lru.m[k]
	if ok {
		listNode.V = v
		lru.moveToFront(listNode)
		return
	}

	if len(lru.m) == lru.n {
		tailNode := lru.dummy.Pre
		lru.remove(tailNode)
		delete(lru.m, k)
	}

	newNode := &ListNode{
		K: k,
		V: v,
	}
	lru.insertToFront(newNode)
	lru.m[k] = newNode
}


func (lru *LRU) moveToFront(ln *ListNode) {
	lru.remove(ln)

	lru.insertToFront(ln)
}

func (lru *LRU) insertToFront(ln *ListNode) {
	ln.Pre = lru.dummy
	ln.Next = lru.dummy.Next
	ln.Pre.Next = ln
	ln.Next.Pre = ln
}

func (lru *LRU) remove(ln *ListNode) {
	ln.Pre.Next = ln.Next
	ln.Next.Pre = ln.Pre
}