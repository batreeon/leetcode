/*
 * @lc app=leetcode.cn id=1018 lang=golang
 *
 * [1018] 可被 5 整除的二进制前缀
 */

// @lc code=start
// package main

// import (
// 	"fmt"
// 	"math"
// )

func prefixesDivBy5(A []int) []bool {
	ans := []bool{}
	for i := 0 ; i < len(A); i++ {
		num := 0
		for j := 0 ; j <= i ; j++ {
			num = num + A[j]*int(math.Pow(2,float64(i-j)))
		}
		// fmt.Println(num)
		if num%5 == 0 {
			ans = append(ans,true)
			// fmt.Println("t")
		}else{
			ans = append(ans,false)
			// fmt.Println("f")
		}
	}
	return ans
}

// func main() {
// 	A := []int{1,0,1,1,1,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,1,1,1,1,0,0,0,0,1,1,1,0,0,0,0,0,1,0,0,0,1,0,0,1,1,1,1,1,1,0,1,1,0,1,0,0,0,0,0,0,1,0,1,1,1,0,0,1,0}
// 	prefixesDivBy5(A)
// }
// @lc code=end

