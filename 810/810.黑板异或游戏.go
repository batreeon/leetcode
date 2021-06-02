/*
 * @lc app=leetcode.cn id=810 lang=golang
 *
 * [810] 黑板异或游戏
 */

// @lc code=start
package main
func xorGame(nums []int) bool {
	/*
	// 哇哇哇，心态崩了，没做出来，看了答案才发现我是小丑
	// 我来试试回溯，暴力奥
	xor := 0
	for _,num := range nums {
		xor ^= num
	}
	var backtrack func(nums []int,xor int,alice bool) bool 
	backtrack = func(nums []int,xor int,alice bool) bool {
		if alice && xor == 0 {
			return true
		}
		all0 := true
		for i,num := range nums {
			newnums := make([]int,len(nums)-1)
			copy(newnums[:i],nums[:i])
			copy(newnums[i:],nums[i+1:])
			xor ^= num
			if xor != 0 {
				all0 = false
				if alice && backtrack(newnums,xor,!alice) {
					return true
				}else if !alice && !backtrack(newnums,xor,!alice) {
					return false
				}
			}
			xor ^= num
		}
		if all0 {
			if alice {
				return false
			}else{
				return true
			}
		}
		return !alice
	}
	return backtrack(nums,xor,true)
	*/
	// 如果一开始所有元素的异或和为0，那么返回true

	// 一开始xor不为0，那么就要alice去掉一个数组，剩余的异或和不为0，令剩余元素异或和为xor1
	// 那么alice必胜，则自持bob去掉任意一个数字，都要得到新的xor2==0
	// 就有xor1^nums[0] = xor1^nums[1] = xor1^nums[2] = ...... = xor1^nums[n] = 0
	// 由上式得，(xor1^xor1^...^xor1)^(nums[0]^nums[1]^...^nums[n]) = 0
	// 且由于 nums[0]^nums[1]^...^nums[n] = xor1
	// 则上上式可得(xor1^xor1^...^xor1)^xor1 = 0
	// 由于xor1不为0，那么上式中xor1的个数要为偶数个，则bob面对的nums有奇数个数字（上式中原本有一个xor1是在括号外面的）
	// 则初始状态下alice面对的nums要有偶数个元素，就可以保证alice必胜

	// 那么上述结论靠谱吗？
	// 也就是说是否会出现，nums有偶数个元素，先手是否会输？
	// 如果先手会输，则会有xor^nums[i] = 0的情况出现，则跟上述的分析方法，最终会出现奇数个xor的异或结果为0的情况
	// 则条件必须为xor=0,这与我们开头第二行注释的前提条件相悖
	if len(nums) % 2 == 0 {
		return true
	}
	var xor int
	for _,num := range nums {
		xor ^= num 
	}
	return xor == 0
}
// @lc code=end

