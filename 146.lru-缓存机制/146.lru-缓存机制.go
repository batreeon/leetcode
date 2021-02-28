/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */

// @lc code=start
/*
使用了一个map和slice，map的value是一个silce，slice的元素记录key。
每次使用了key就把对应key移动到slice的末尾。但是有一个问题，每次一移动，所有的key对应的下标都要变
需要遍历改变map中的所有元素，这成本太高了
package main
type LRUCache struct {
	lruM map[int][]int	//value是一个slice,第一个元素存key对应的值，第二个元素存key在lru中的下标
	lru []int
	len int
	cap int
}


func Constructor(capacity int) LRUCache {
	return LRUCache{lruM:make(map[int][]int) , lru:[]int{} , len:0 , cap:capacity}
}


func (this *LRUCache) Get(key int) int {
	if v,ok := this.lruM[key] ; ok {
		copy(this.lru[v[1]:] , this.lru[v[1]+1:])
		this.lru[this.len-1] = key
		this.lruM[key][1] = this.len-1
		return v[0]
	}else{
		return -1
	}
}


func (this *LRUCache) Put(key int, value int)  {
	if v,ok := this.lruM[key] ; ok {
		copy(this.lru[v[1]:] , this.lru[v[1]+1:])
		this.lru[this.len-1] = key
		this.lruM[key][0] = value
		this.lruM[key][1] = this.len-1
	}else{
		if this.len < this.cap {
			this.lru = append(this.lru,key)
			this.len++
			this.lruM[key] = []int{value,this.len-1}
		}else{
			delete(this.lruM,this.lru[0])
			copy(this.lru[:] , this.lru[1:])
			this.lru = this.lru[:this.len-1]
			this.lru = append(this.lru,key)
			
			this.lruM[key] = []int{value,this.len-1}
		}
	}
}
*/

//官方解使用哈希表和双向链表
type LRUCache struct {
    size int
    capacity int
    cache map[int]*DLinkedNode
    head, tail *DLinkedNode
}

type DLinkedNode struct {
    key, value int
    prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
    return &DLinkedNode{
        key: key,
        value: value,
    }
}

func Constructor(capacity int) LRUCache {
    l := LRUCache{
        cache: map[int]*DLinkedNode{},
        head: initDLinkedNode(0, 0),
        tail: initDLinkedNode(0, 0),
        capacity: capacity,
    }
    l.head.next = l.tail
    l.tail.prev = l.head
    return l
}

func (this *LRUCache) Get(key int) int {
    if _, ok := this.cache[key]; !ok {
        return -1
    }
    node := this.cache[key]
    this.moveToHead(node)
    return node.value
}


func (this *LRUCache) Put(key int, value int)  {
    if _, ok := this.cache[key]; !ok {
        node := initDLinkedNode(key, value)
        this.cache[key] = node
        this.addToHead(node)
        this.size++
        if this.size > this.capacity {
            removed := this.removeTail()
            delete(this.cache, removed.key)
            this.size--
        }
    } else {
        node := this.cache[key]
        node.value = value
        this.moveToHead(node)
    }
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
    node.prev = this.head
    node.next = this.head.next
    this.head.next.prev = node
    this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
    this.removeNode(node)
    this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
    node := this.tail.prev
    this.removeNode(node)
    return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

