/*
 * @lc app=leetcode.cn id=519 lang=golang
 *
 * [519] 随机翻转矩阵
 */

// @lc code=start
package main

import "math/rand"

type Solution struct {
	matrix [][]int
	r int
	c int
	total int
	indexs []int
	num1 int
}


func Constructor(m int, n int) Solution {
	matrix := make([][]int, m)
	for i := range matrix {
		row := make([]int, n)
		matrix[i] = row
	}

	indexs := make([]int, m * n)
	for i := range indexs {
		indexs[i] = i
	}

	solution := Solution{
		matrix: matrix,
		r: m,
		c: n,
		total: m * n,
		indexs: indexs,
		num1: 0,
	}
	return solution
}


func (this *Solution) Flip() []int {
	idx := rand.Intn(this.total - this.num1) + this.num1
	flipIdx := this.indexs[idx]
	this.indexs[idx], this.indexs[this.num1] = this.indexs[this.num1], this.indexs[idx]
	r,c := flipIdx/this.c, flipIdx%this.c
	this.matrix[r][c] = 1
	this.num1++
	return []int{r,c}
}


func (this *Solution) Reset()  {
	for this.num1 > 0 {
		this.num1--
		r,c := this.indexs[this.num1]/this.c, this.indexs[this.num1]%this.c
		this.matrix[r][c] = 0
	}
	for i := range this.indexs {
		this.indexs[i] = i
	}
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(m, n);
 * param_1 := obj.Flip();
 * obj.Reset();
 */
// @lc code=end

