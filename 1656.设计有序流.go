/*
 * @lc app=leetcode.cn id=1656 lang=golang
 *
 * [1656] 设计有序流
 */

// @lc code=start
type OrderedStream struct {
    nums []string
	ptr int
}


func Constructor(n int) OrderedStream {
    return OrderedStream{
		nums: make([]string, n+1),
		ptr: 1,
	}
}


func (this *OrderedStream) Insert(idKey int, value string) []string {
    this.nums[idKey] = value
	res := []string{}
	if this.nums[this.ptr] != "" {
		for this.ptr < len(this.nums) && this.nums[this.ptr] != "" {
			res = append(res, this.nums[this.ptr])
			this.ptr++
		}
	}
	return res
}


/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
// @lc code=end

