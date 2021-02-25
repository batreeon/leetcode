/*
 * @lc app=leetcode.cn id=292 lang=golang
 *
 * [292] Nim 游戏
 */

// @lc code=start
func canWinNim(n int) bool {
	//n=1,2,3 true
	//n=4 false
	//若n是4的倍数，则false.
	//因为不管你每次取1,2,3，对手都可以取相应的3，2，1，每次取一轮，都还是会剩下4的倍数
	return n % 4 != 0
}
// @lc code=end

