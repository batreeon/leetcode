/*
 * @lc app=leetcode.cn id=460 lang=golang
 *
 * [460] LFU 缓存
 */

// @lc code=start
package main

import (
	"fmt"
	"sort"
)

type Value struct {
	key int
	value int
	count int
	time int
}

type list []Value

func (l list) Len() int {return len(l)}
func (l list) Less(i,j int) bool {
	return (l[i].count > l[j].count) || (l[i].count == l[j].count && l[i].time > l[j].time)
}
func (l list) Swap(i,j int) {
	l[i],l[j] = l[j],l[i]
}

type LFUCache struct {
	m map[int]*Value
	list
	time int
	cap int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		map[int]*Value{},
		list{},
		0,
		capacity,
	}
}


func (this *LFUCache) Get(key int) int {
	if this.cap == 0 {
		return -1
	}
	if v,ok := this.m[key]; ok {
		v.count++
		v.time = this.time
		for i := range this.list {
			if this.list[i].key == key {
				this.list[i].count++
				this.list[i].time = this.time
			}
		}
		this.time++
		return v.value 
	}
	return -1
}


func (this *LFUCache) Put(key int, value int)  {
	if this.cap == 0 {
		return
	}
	if v,ok := this.m[key]; ok {
		v.count++
		v.time = this.time
		v.value = value
		for i := range this.list {
			if this.list[i].key == key {
				this.list[i].count++
				this.list[i].time = this.time
			}
		}
	}else{
		newValue := Value{
			key,
			value,
			1,
			this.time,
		}
		if this.Len() == this.cap {
			sort.Sort(this.list)
			removed := this.list[this.cap-1]
			delete(this.m,removed.key)
			this.list[this.cap-1] = newValue
		}else{
			this.list = append(this.list, newValue)
		}
		this.m[key] = &newValue
	}
	this.time++
}

// func main() {
// 	lfu := Constructor(2)
// 	lfu.Put(1,1)
// 	lfu.Put(2,2)
// 	fmt.Println(lfu.Get(1))
// 	lfu.Put(3,3)
// 	fmt.Println(lfu.Get(2))
// 	fmt.Println(lfu.Get(3))
// 	lfu.Put(4,4)
// 	fmt.Println(lfu.Get(1))
// 	fmt.Println(lfu.Get(3))
// 	fmt.Println(lfu.Get(4))
// }

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

