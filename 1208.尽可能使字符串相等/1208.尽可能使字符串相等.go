/*
 * @lc app=leetcode.cn id=1208 lang=golang
 *
 * [1208] 尽可能使字符串相等
 */

// @lc code=start
// package main 
// import "fmt"
func equalSubstring(s string, t string, maxCost int) int {
	// if maxCost == 0 {
	// 	return 0
	// } //不能直接通过maxCost判断，比如s="a",t="a"，字符串也可相等
	cost := make([]int,len(s))
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 0 ; i < len(s) ; i++ {
		// fmt.Println(s[i]," ",t[i])
		// fmt.Println(int(s[i])-int(t[i]))
		// fmt.Println(int(s[i]-t[i])) //不能用这个byte只能表示uint8
		// fmt.Println(s[i]-t[i])
		cost[i] = abs(int(s[i])-int(t[i]))
	}
	left,right := 0,0
	nowCost := 0
	maxLen := 0
	for ; right < len(s) ; right++ {
		nowCost += cost[right]
		if nowCost > maxCost {
			nowCost -= cost[left]
			left++
		}
		maxLen = max(maxLen,right-left+1)
	}
	return maxLen
}

// func main() {
// 	// s := string("abcd")
// 	// t := string("bcdf")
// 	// fmt.Println(equalSubstring(s,t,3))
// 	// s := string("abcd")
// 	// t := string("cdef")
// 	// fmt.Println(equalSubstring(s,t,1))
// 	s := string("pxezla")
// 	t := string("loewbi")
// 	fmt.Println(equalSubstring(s,t,25))
// }
// @lc code=end

