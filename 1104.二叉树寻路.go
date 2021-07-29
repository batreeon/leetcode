/*
 * @lc app=leetcode.cn id=1104 lang=golang
 *
 * [1104] 二叉树寻路
 */

// @lc code=start
package main

import (
	"math"
)

func pathInZigZagTree(label int) []int {
	// 记录每一行节点的最大标记1，3，7，15。。。
	big := []int{}
	// i用于计算big元素值，也可以用来记录层数
	i := 1
	for  ; ; i++ {
		cur := int(math.Pow(2,float64(i)))-1
		if cur >= label {
			break
		}
		big = append(big,cur)
	}
	// 共有i层
	result := make([]int,i)
	result[i-1] = label
	i--
	// 第一层一定为1，就不用算了
	for i > 1 {
		// 如果不是之字形，那么节点n的父亲应为n/2，但因为是之字形的
		// 所以应根据该层元素的最大值和最小值，算出与n/2位置对称的那个值
		// 第i层的最小值是第i-1层最大值+1，最大值就是big[i-1]
		// n/2在该层的位置对称的值记为x,那么n/2到最小值的差应该等于x到最大值的差
		result[i-1] = big[i-2]+1+(big[i-1]-result[i]/2)
		i--
	}
	result[0] = 1
	return result
}
// @lc code=end

