/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	n := len(nums)
	ans := math.MaxInt64
	for i := 0 ; i < n - 2 ; i++ {
		v1 := nums[i] - target
		for j := i+1 ; j < n- 1; j++ {
			v2 := v1 + nums[j]
			for k := j+1 ; k < n ; k++ {
				v := v2 + nums[k]
				if abs(ans) > abs(v){
					ans = v
				}
				if ans == 0 {
					return ans + target
				}
			}
		}
	}
	return ans + target
}
// @lc code=end

