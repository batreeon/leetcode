/*
 * @lc app=leetcode.cn id=1004 lang=golang
 *
 * [1004] 最大连续1的个数 III
 */

// @lc code=start
// package main 

// import (
// 	"fmt"
// )

func longestOnes(A []int, K int) int {
	left := 0
	res := 0//记录当前窗口需要翻转的个数
	ans := 0
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	for right := 0 ; right < len(A) ; right++ {
		if A[right] == 0 {
			res++
		}
		for res > K {
			if A[left] == 0 {
				res--
			}
			left++
		}
		ans = max(ans,right-left+1)
	}
	return ans
}

// func main() {
// 	a := []int{1,1,1,0,0,0,1,1,1,1,0}
// 	fmt.Println(longestOnes(a,2))
// }
// @lc code=end

