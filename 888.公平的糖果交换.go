/*
 * @lc app=leetcode.cn id=888 lang=golang
 *
 * [888] 公平的糖果交换
 */

// @lc code=start
func fairCandySwap(A []int, B []int) []int {
	sum1,sum2 := Sum(A),Sum(B)
	sumD2 := (sum1+sum2)/2
	for _,a := range A {
		//(sum1-ans1+ans2)*2 = sum
		ans2 := sumD2 - sum1 + a
		for _,b := range B {
			if b == ans2 {
				return []int{a,b}
			}
		}
	}
	return []int{}
}
func Sum(A []int) (sum int) {
	for _,a := range A {
		sum += a
	}
	return 
}
// @lc code=end

