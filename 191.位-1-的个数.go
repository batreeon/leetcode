/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(num uint32) int {
	nums := 0
	for i:=0;i<32;i++ {
		if (num >> i)&1 == 1 {
			nums++
		}
	}
	return nums
}
// @lc code=end

