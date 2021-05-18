/*
 * @lc app=leetcode.cn id=1442 lang=golang
 *
 * [1442] 形成两个异或相等数组的三元组数目
 */

// @lc code=start
func countTriplets(arr []int) int {
	s := make([]int,len(arr)+1)
	for i := range arr {
		s[i+1] = s[i]^arr[i]
	}
	result := 0
	for i := 0 ; i < len(arr) ; i++ {
		for j := i+1 ; j <= len(arr) ; j++ {
			if s[i]^s[j] == 0 {
				// s[i]^s[j] == 0成立说明arr[i]^arr[i+1]^...arr[j-1]==0
				// 那么i取i,j可以取i+1,i+2...j-1,k可以去j-1
				// j总共有 (j-1) - (i+1) + 1种
				result += (j-i-1)
			}
		}
	}
	return result
}
// @lc code=end

