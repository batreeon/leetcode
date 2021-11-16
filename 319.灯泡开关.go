/*
 * @lc app=leetcode.cn id=319 lang=golang
 *
 * [319] 灯泡开关
 */

// @lc code=start
package main

import "math"

func bulbSwitch(n int) int {
	// 一个灯泡是打开的：那么它被开关奇数次，那么它的因数有奇数个，而因数总是成对出现的，
	// 若他的因数有奇数个，只有一种可能，它是完全平方数
	// [1,n]之间完全平方数的个数等于 sqrt(n) ，向下取整
	return int(math.Sqrt(float64(n)))
}
// @lc code=end

