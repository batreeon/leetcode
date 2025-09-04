/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
	n := len(nums)

	result := [][]int{}
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] + nums[i+1] + nums[i+1] > 0 {
			break
		}
		if nums[i] + nums[n-2] + nums[n-1] < 0 {
			continue
		}

		for j,k := i+1, n-1; j < k; {
			s := nums[i]+nums[j]+nums[k]
			if s > 0 {
				k--
			}else if s < 0 {
				j++
			}else{
				result = append(result, []int{nums[i],nums[j],nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}

	return result
}
// @lc code=end

