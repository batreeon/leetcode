/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */

// @lc code=start
package main

import "container/list"

type ListValue struct {
	k   int
	v   int
}

// 注意链表存的元素类型是*ListValue，不是LiatValue，因为要做修改，所以使用指针
type LRUCache struct {
	volume int
	l      *list.List
	m      map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		volume: capacity,
		l:      list.New(),
		m:      map[int]*list.Element{},
	}
	return lru
}

func (this *LRUCache) Get(key int) int {
	if ele,ok := this.m[key] ; ok {
		this.l.MoveToFront(ele)
		return ele.Value.(*ListValue).v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if ele,ok := this.m[key] ; ok {
		this.l.MoveToFront(ele)
		ele.Value.(*ListValue).v = value
	}else{
		ele := this.l.PushFront(&ListValue{
			k:key,
			v:value,
		})
		this.m[key] = ele
		if len(this.m) > this.volume {
			ele := this.l.Back()
			delete(this.m,ele.Value.(*ListValue).k)
			this.l.Remove(ele)
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
