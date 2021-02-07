/*
 * @lc app=leetcode.cn id=665 lang=golang
 *
 * [665] 非递减数列
 */

// @lc code=start
func checkPossibility(nums []int) bool {
	// //这个太笨了，虽然是对的，但太慢且消耗内存
	// //依次删除一个元素，若删除后，得到一个递增数列，就返回true
	// if len(nums) == 0 || len(nums) == 1 {
	// 	return true
	// }
	// l := len(nums)
	// numsDel1 := make([]int,l-1)
	// for i := 0 ; i < l; i++ {
	// 	copy(numsDel1[:i],nums[:i])
	// 	copy(numsDel1[i:],nums[i+1:])
	// 	j := 0
	// 	for ; j < l-2 ; {
	// 		if numsDel1[j] > numsDel1[j+1] {
	// 			break
	// 		}
	// 		j++
	// 	}
	// 	if j == l-2 {
	// 		return true
	// 	}
	// }
	// return false

	illegal := 0
	for i := 0 ; i < len(nums)-1 ; i++ {
		if nums[i] > nums[i+1] {
			if i != 0 && i+1 != len(nums)-1 && nums[i-1] > nums[i+1] && nums[i] > nums[i+2] {
				return false
			}else{
				illegal++
			}
			if illegal > 1 {
				return false
			}
		}
	}
	return true
}
// @lc code=end

