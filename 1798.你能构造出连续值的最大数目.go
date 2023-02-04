/*
 * @lc app=leetcode.cn id=1798 lang=golang
 *
 * [1798] 你能构造出连续值的最大数目
 */

// @lc code=start
func getMaximumConsecutive(coins []int) int {
	// 可以用回溯吗 感觉复杂度有点高

	// 直接看的答案
	// 加入我们可以构造的连续序列是[0,x]
	// 新增一个数y, 那么可以构造出新的范围为[0,x] + [y, x+y]
	// 如果y <= x+1 , 那么新的范围为[0, x+y]
	// 所以为了连续，我们每次取coins中最小的数；如果最小的数都不能连续，那就可以返回了
	
	x := 0
	sort.Ints(coins)
	for _, y := range coins {
		if y <= x+1 {
			x = x + y
		}else{
			break
		}
	}
	return x+1
}
// @lc code=end

