/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

// @lc code=start
// package main

// import "fmt"

func reverseBits(num uint32) uint32 {
	ans := ^uint32(0)//全1的32位值
	for i := 0; i < 32;i++ {
		if ((num>>i) & 1) == 0 {
			ans -= 1<<(31-i)
			//fmt.Println(ans)
		}
	}
	return ans
}

// func main() {
// 	fmt.Println(reverseBits(uint32(0)))
// }
// @lc code=end

