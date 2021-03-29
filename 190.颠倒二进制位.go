/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

// @lc code=start
package main
func reverseBits(num uint32) uint32 {
    // 20210329，思路，每次最右边的位，然后一个一个的去合并进去
	var ans uint32
	for i := 0 ; i < 32 ; i++ {
		ans <<= 1
		ans |= num&1
		num >>= 1
	}
	return ans
}
// @lc code=end

