/*
 * @lc app=leetcode.cn id=1006 lang=golang
 *
 * [1006] 笨阶乘
 */

// @lc code=start
func clumsy(N int) int {
	// if N == 1 {return 1}
	// if N == 2 {return 2}
	// if N == 3 {return 6}
	// if N == 4 {return 7}
	switch N {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 6
	case 4:
		return 7
	}

	// var ans int
	switch (N-3)%4 {
	case 1:
		// 末尾剩+1
		// 按理说应该(N+1)+..-..+..-..+1 = N+2的，
		// 但是因为这种情况下末尾四位为-4*3/2+1,4*3/2不等于5而等于6，
		// （这是因为4刚好是2的倍数，因此一家）
		// 这与它前面的+5无法抵消，而应该在最后的结果中再多减1
		return N+1
	case 2:
		// 末尾剩+2-1
		return N+2
	case 3:
		// 末尾剩+3-2*1
		return N+2
	case 0:
		// 末尾剩+4-3*2/1
		return N-1
	}
	return -1
}
// @lc code=end

