/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 */

// @lc code=start
func singleNumber(nums []int) []int {
	// 策略，将所有元素异或，得到的结果等于两个只出现一次元素的异或结果
	// 比如a,b只出现一次，最终异或结果为a^b,只要分别再和两个元素异或一次，就可分别得到另一个元素
	// 如何实现呢？
	temp := 0
	for _,num := range nums {
		temp ^= num
	}
	ans := []int{temp,temp}
	// 通过得到temp的最低位1，来对a和b进行区分。
	// 因为temp中某一位为1，说明a和b中，再改为一个有1，一个没1，根据想与结果进行区分
	// 而其他重复出现的元素会进入同一分支，对结果无影响
	// 太妙了太妙了，妙不可言
	temp = (temp&(temp-1))^temp
	for _,num := range nums {
		if temp&num == 0 {
			ans[0] ^= num
		}else{
			ans[1] ^= num
		}
	}
	return ans
}
// @lc code=end

