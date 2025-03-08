/*
 * @lc app=leetcode.cn id=2597 lang=golang
 *
 * [2597] 美丽子集的数目
 */

// @lc code=start
func beautifulSubsets(nums []int, k int) int {
	result := 0
    backtrack(nums, map[int]int{}, k, &result)
	return result
}

func backtrack(nums []int, count map[int]int, k int, result *int) {
	if len(count) > 0 {
		(*result)++
	}
	for i, num := range nums {
		// 根据题义，nums是有重复元素的，是可以允许有重复结果出现的。

		// 这里count不能用set类型，因为可能有多个相同的元素在子数组里，如果使用set，
		// 那么当撤回一个元素时，只有当子数组中没有相同的元素时，才应该从set中将该值删掉，否则会出错。
		if count[num+k] > 0 || count[num-k] > 0 {
			continue
		}
		count[num]++
		if i != len(nums)-1 {
			backtrack(nums[i+1:], count, k, result)
		}else{
			backtrack([]int{}, count, k, result)
		}
		count[num]--
		if count[num] == 0 {
			delete(count, num)
		}
	}
}
// @lc code=end

