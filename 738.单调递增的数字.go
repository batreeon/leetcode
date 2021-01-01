import "strconv"

/*
 * @lc app=leetcode.cn id=738 lang=golang
 *
 * [738] 单调递增的数字
 */

// @lc code=start
func monotoneIncreasingDigits(N int) int {
	strN := []byte(strconv.Itoa(N)) //先转换成string,在转换成字节序
	i := 1
	for ; i < len(strN) && strN[i-1] <= strN[i]; i++ {
	}
	//找到最左边的不满足递增的地方，不用管有多处不满足的
	//我们是要找最大值，将最左处的调整了（高位影响大），将右边的数全部置为9
	if i < len(strN) { //若N不是递增
		//i指示的是小的那个数
		for i > 0 && strN[i-1] > strN[i] {
			//因为调整了一个数（减1），可能导致左边不满足递增，因此一直往左走
			strN[i-1]--
			i -= 1
		}
		//若i等于0，则要么最高为降1，要么最高为变成0，将右边的位全置9
		//若i不等于0，则i位降了1，将其右的位全置9
		for i += 1; i < len(strN); i++ {
			strN[i] = '9'
		} //这个已经将最高位变成0的情况考虑进去了
		// ans := 0
		// for i = 0 ; i < len(strN) ; i++ {
		// 	ans = 10 * ans + int(strN[i])
		// }
		ans, _ := strconv.Atoi(string(strN))
		return ans
	}
	return N
}

// @lc code=end

