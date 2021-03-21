/*
给你一个整数数组 nums （下标 从 0 开始 计数）以及两个整数：low 和 high ，请返回 漂亮数对 的数目。

漂亮数对 是一个形如 (i, j) 的数对，其中 0 <= i < j < nums.length 且 low <= (nums[i] XOR nums[j]) <= high 。
*/
func countPairs(nums []int, low int, high int) int {
	// l := len(nums)
	// ans := 0
	// for i := 0 ; i < l-1 ; i++ {
	//     for j := i+1 ; j < l ; j++ {
	//         res := nums[i]^nums[j]
	//         if res >= low && res <= high {
	//             ans++
	//         }
	//     }
	// }
	// return ans

}