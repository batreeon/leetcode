/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */

// @lc code=start
func candy(ratings []int) int {
	// 从左向右遍历，从右向左遍历，取较大值
	n := len(ratings)
	ltor,rtol := make([]int,n),make([]int,n)
	ltor[0],rtol[n-1] = 1,1
	for i := 1 ; i < n ; i++ {
		j := n-1-i
		ltor[i] = 1
		rtol[j] = 1
		if ratings[i] > ratings[i-1] {
			ltor[i] = ltor[i-1]+1
		}
		if ratings[j] > ratings[j+1] {
			rtol[j] = rtol[j+1]+1
		}
	}

	ans := 0
	for i := 0 ; i < n ; i++ {
		if rtol[i] > ltor[i] {
			ans += rtol[i]
		}else{
			ans += ltor[i]
		}
	}
	return ans
}
// @lc code=end

