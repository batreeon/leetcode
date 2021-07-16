/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

 func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
    pNode := head
	for pNode != nil {
		node := Node{
			Val:pNode.Val,
			Next:pNode.Next,
		}
		pNode.Next = &node
		pNode = node.Next
	}

	pNode = head
	for pNode != nil {
        if pNode.Random != nil {
            pNode.Next.Random = pNode.Random.Next
        }
		pNode = pNode.Next.Next
	}
    
	pNode = head
	newHead := pNode.Next
	pNew := newHead
	for pNode != nil {
		pNode.Next = pNode.Next.Next
		if pNew.Next != nil {
			pNew.Next = pNew.Next.Next
            pNew = pNew.Next
		}
        pNode = pNode.Next
	}
	return newHead
}