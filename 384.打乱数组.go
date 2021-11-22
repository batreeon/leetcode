/*
 * @lc app=leetcode.cn id=384 lang=golang
 *
 * [384] 打乱数组
 */

// @lc code=start
package main

import "math/rand"

type Solution struct {
	nums []int
}


func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}


func (this *Solution) Reset() []int {
	return this.nums
}


func (this *Solution) Shuffle() []int {
	numsCopy := make([]int, len(this.nums))
	copy(numsCopy, this.nums)
	// 这个随机算法记一记，每次选择一个剩余的数，放在第i个位置上
	for i := 0; i < len(this.nums); i++ {
		idx := i + rand.Intn(len(this.nums) - i)
		numsCopy[i], numsCopy[idx] = numsCopy[idx], numsCopy[i]
	}
	return numsCopy
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end

