/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(num int) []int {
	/*
	ans := make([]int,num+1)
	back := 1
	for i,j := 1,back ; i <= num ; i++ {
		ans[i] = ans[i-back]+1
		j--
		if j == 0 {
			back *= 2
			j = back
		}
	}
	return ans
	*/

	/*
	index 0 1 2 3 4 5 6 7 8
	value 0 1 1 2 1 2 2 3 1
	初始0，每位加1复制一份添加到末尾 0 1
	每位加1复制一份添加到末尾 0 1 1 2
	每位加1复制一份添加到末尾 0 1 1 2 1 2 2 3
	*/
	ans := make([]int,num+1)
	for i,start := 1,1 ; i <= num ; i++ {
		if i/2 == start {
			start *= 2
		}
		ans[i] = ans[i-start]+1
	}
	return ans

	/*
	ans := make([]int,num+1)
	if num == 0 {
		return ans
	}
	back := 1
	for i:= 1; i <= num ; i *= 2 {
		// ans[i:] = ans[i-back:i]+1
		//想法就是直接把前back个数据，再复制一份，并每位+1，但怎样实现呢？
		back *=2
	}
	return ans
	*/
}
// @lc code=end

