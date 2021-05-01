/*
 * @lc app=leetcode.cn id=633 lang=golang
 *
 * [633] 平方数之和
 */

// @lc code=start
package main
// import "math"

func judgeSquareSum(c int) bool {

	/*
	// 错误的答案
	isSquare := func(n int) bool {
		for i := 1; n > 0; i += 2 {
			n -= i
		}
		return n == 0
	}
	n := c
	for {
		if isSquare(n) {
			break
		}
		n++
	}
	return n-c == 0 || (n - c > 2 && (n - c) % 2 == 0)
	*/

	/*
	// 错误的
	isSquare := func(n int) bool {
		for i := 1; n > 0; i += 2 {
			n -= i
		}
		return n == 0
	}
	if isSquare(c) {
		return true
	}
	n := c
	lastn := c+1
	for n > 0 {
		if isSquare(n) {
			if isSquare(c-n) {
				return true
			}
			if lastn == c+1 {
				lastn = n
				n--
			}else{
				if lastn-n-2 > 0 {
					break
				}
				lastn,n = n,n-(lastn-n-2)
			}
		}else{
			n--
		}
	}
	return false
	*/

	/*
	// 官解，二分查找，好叼啊
	left,right := 0,int(math.Sqrt(float64(c)))
	for left <= right {
		temp := left * left + right * right
		if temp == c {
			return true
		}
		if temp > c {
			right--
		}
		if temp < c {
			left++
		}
	}
	return false
	*/

	for base := 2 ; base * base <= c ; base++ {
		if c % base != 0 {
			continue
		}
		exp := 0
		for c % base == 0 {
			c /= base
			exp++
		}
		if base % 4 == 3 && exp % 2 != 0 {
			return false
		}
	}
	return c % 4 != 3
}
// @lc code=end

