/*
 * @lc app=leetcode.cn id=528 lang=golang
 *
 * [528] 按权重随机选择
 */

// @lc code=start
package main

import (
	"math/rand"
	"sort"
)

// 这里有一个问题，w[i] <= 100000,w.length <= 10000,
// 因此w数据总和最大10的9次方，可能会超过int范围
// 并且使用的随机函数rand.Intn也只返回int类型，可能无法覆盖sum的最大值，但是最后还是通过了
type Solution struct {
	sum []int
}

func Constructor(w []int) Solution {
	n := len(w)
	sum := make([]int,n+1)
	sum[0] = 0
	for i := 1 ; i <= n ; i++ {
		sum[i] = sum[i-1] + w[i-1]
	}
	return Solution{sum}
}

func (this *Solution) PickIndex() int {
	n := len(this.sum)
	max := this.sum[n-1]
	num := rand.Intn(max)
	return sort.Search(n,func(i int) bool {return this.sum[i] > num})-1
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
// @lc code=end

