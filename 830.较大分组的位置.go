/*
 * @lc app=leetcode.cn id=830 lang=golang
 *
 * [830] 较大分组的位置
 */

// @lc code=start
func largeGroupPositions(s string) [][]int {
	ans := [][]int{}
	for i := 0 ; i < len(s) ; {

		start := i
		// if i > 0 && s[i] != s[i-1] {
		// 	start = i
		// }//此时i指示的是一个分组的首位

		i++//此时指示的是一个分组首位的下一位

		for i < len(s) && s[i] == s[i-1] {
			i++
		}//上一个循环退出时，i指示的是上一分组末尾的下一位

		if i-start > 2 {
			ans = append(ans,[]int{start,i-1})
		}
	}
	return ans
}
// @lc code=end

