/*
 * @lc app=leetcode.cn id=201 lang=golang
 *
 * [201] 数字范围按位与
 */

// @lc code=start
func rangeBitwiseAnd(left int, right int) int {
	var rangebitwiseand func(left,right int) int
	rangebitwiseand = func(left,right int) int {
		copyRight := right 
		for copyRight&(copyRight-1) != 0 {
			copyRight &= copyRight-1
		}
		// 找到right的最高位1，上述遍历完后，copyRight就是right的最高位1

		if copyRight&left == 0 {
			return 0
		}

		// 去掉最高位1
		mask := copyRight-1
		left &= mask
		right &= mask
		return copyRight+rangebitwiseand(left,right)
	}
	return rangebitwiseand(left,right)
	// ans := left
	// for ; left < right ; {
	// 	ans &= left+1
	// 	left++
	// }
	// return ans
}
// @lc code=end

