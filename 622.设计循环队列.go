/*
 * @lc app=leetcode.cn id=622 lang=golang
 *
 * [622] 设计循环队列
 */

// @lc code=start
/*
type MyCircularQueue struct {
	mycircularqueue []int
}


func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		make([]int,0,k),
	}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
	if len(this.mycircularqueue) == cap(this.mycircularqueue) {
		return false
	}
	this.mycircularqueue = append(this.mycircularqueue,value)
	return true
}


func (this *MyCircularQueue) DeQueue() bool {
	l := len(this.mycircularqueue)
	if l == 0 {
		return false
	}
	copy(this.mycircularqueue[0:l-1],this.mycircularqueue[1:l])
	this.mycircularqueue = this.mycircularqueue[0:l-1]
	return true
}


func (this *MyCircularQueue) Front() int {
	if len(this.mycircularqueue) == 0 {
		return -1
	}
	return this.mycircularqueue[0]
}


func (this *MyCircularQueue) Rear() int {
	l := len(this.mycircularqueue)
	if l == 0 {
		return -1
	}
	return this.mycircularqueue[l-1]
}


func (this *MyCircularQueue) IsEmpty() bool {
	return len(this.mycircularqueue) == 0
}


func (this *MyCircularQueue) IsFull() bool {
	return len(this.mycircularqueue) == cap(this.mycircularqueue)
}
// 使用分配好cap的切片，性能反倒不好了
*/

type MyCircularQueue struct {
	mycircularqueue []int
	cap int
}


func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		[]int{},
		k,
	}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
	if len(this.mycircularqueue) >= this.cap {
		return false
	}
	this.mycircularqueue = append(this.mycircularqueue,value)
	return true
}


func (this *MyCircularQueue) DeQueue() bool {
	if len(this.mycircularqueue) == 0 {
		return false
	}
	this.mycircularqueue = this.mycircularqueue[1:]
	return true
}


func (this *MyCircularQueue) Front() int {
	if len(this.mycircularqueue) == 0 {
		return -1
	}
	return this.mycircularqueue[0]
}


func (this *MyCircularQueue) Rear() int {
	l := len(this.mycircularqueue)
	if l == 0 {
		return -1
	}
	return this.mycircularqueue[l-1]
}


func (this *MyCircularQueue) IsEmpty() bool {
	return len(this.mycircularqueue) == 0
}


func (this *MyCircularQueue) IsFull() bool {
	return len(this.mycircularqueue) == this.cap
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
// @lc code=end

