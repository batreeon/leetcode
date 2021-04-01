/*
 * @lc app=leetcode.cn id=641 lang=golang
 *
 * [641] 设计循环双端队列
 */

// @lc code=start
package main

type MyCircularDeque struct {
	q   []int
	cap int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	mycirculardeque := MyCircularDeque{
		q:   []int{},
		cap: k,
	}
	return mycirculardeque
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if len(this.q) >= this.cap {
		return false
	}
	// 不可行，不会更改len
	// copy(this.q[1:],this.q[0:])
	// this.q[0] = value

	this.q = append([]int{value}, (this.q)...)
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if len(this.q) >= this.cap {
		return false
	}
	this.q = append(this.q, value)
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	l := len(this.q)
	if l == 0 {
		return false
	}
	copy(this.q[0:], this.q[1:])
	this.q = this.q[:l-1]
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	l := len(this.q)
	if l == 0 {
		return false
	}
	this.q = this.q[:l-1]
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if len(this.q) == 0 {
		return -1
	}
	return this.q[0]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	l := len(this.q)
	if l == 0 {
		return -1
	}
	return this.q[l-1]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return len(this.q) == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return len(this.q) == this.cap
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// @lc code=end
