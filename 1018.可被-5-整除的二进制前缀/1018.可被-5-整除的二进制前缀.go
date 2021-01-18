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
	// for i := 0 ; i < len(A); i++ {
	// 	num := 0
	// 	for j := 0 ; j <= i ; j++ {
	// 		num = num + A[j]*int(math.Pow(2,float64(i-j)))
	//		计算出完整值容易溢出，可能是这个原因造成了无法通过
	// 	}
	// 	// fmt.Println(num)
	// 	if num%5 == 0 {
	// 		ans = append(ans,true)
	// 		// fmt.Println("t")
	// 	}else{
	// 		ans = append(ans,false)
	// 		// fmt.Println("f")
	// 	}
	// }

	//计算到第ｉ位时，除以5得到的余数
	prefix := 0
	for _, v := range A {
		prefix = (prefix<<1 | v) % 5
		ans = append(ans, prefix == 0)
	}
	return ans
}

// func main() {
// 	A := []int{1,0,1,1,1,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,1,1,1,1,0,0,0,0,1,1,1,0,0,0,0,0,1,0,0,0,1,0,0,1,1,1,1,1,1,0,1,1,0,1,0,0,0,0,0,0,1,0,1,1,1,0,0,1,0}
// 	prefixesDivBy5(A)
// }
// @lc code=end

