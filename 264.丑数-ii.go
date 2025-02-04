/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	res := []int{1}
	index2,index3,index5 := 0,0,0
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	
	for {
		if len(res) == n {
			return res[n-1]
		}
		next := min(2*res[index2],min(3*res[index3],5*res[index5]))
		res = append(res,next)
		// 下面不应该用不等号<=，而应该用==，因为不可能出现<的情况
		for 2*res[index2] <= next {
			index2++
		}
		for 3*res[index3] <= next {
			index3++
		}
		for 5*res[index5] <= next {
			index5++
		}
	}
}
// @lc code=end

