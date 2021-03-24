/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start
func singleNumber(nums []int) int {
	// 策略：统计每一位1的总数，只有两种情况，0，3，4，
	// 若为4即只出现一次的元素中含有，则要保留，其他两种情况就不用保留了
	result := 0
	// 默认为int64
	for i := 0 ; i < 64 ; i++ {
		count := 0
		for j := 0 ; j < len(nums) ; j++ {
			count += (nums[j]>>i)&1
		}
		// 通过异或将需要保留的1，加到result上去
		// 从低到高位操作，(count%3)<<i只在最高位有1或0，实现在最高位添加1或不添加
		// 不会影响低位的数
		result ^= (count%3)<<i
	}
	return result
}
// @lc code=end

