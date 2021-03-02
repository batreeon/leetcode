/*
 * @lc app=leetcode.cn id=304 lang=golang
 *
 * [304] 二维区域和检索 - 矩阵不可变
 */

// @lc code=start

package main

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	matrix := [][]int{}
	matrix = append(matrix,[]int{3,0,1,4,2})
	matrix = append(matrix,[]int{5,6,3,2,1})
	matrix = append(matrix,[]int{1,2,0,1,5})
	matrix = append(matrix,[]int{4,1,0,1,7})
	matrix = append(matrix,[]int{1,0,3,0,5})
	got := Constructor(matrix)
	t.Error(got.nM,"\n",got.r,"\n",got.c)
}
