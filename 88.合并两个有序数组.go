/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	i,j := 0,0
	// i < m是因为，超过m的都是0，他也会当作正常元素进行比较
	// j < n是为了防止越界，报错
	for ; i < m && j < n ; {
		if nums1[i] < nums2[j] {
			i++
		}else{
			copy(nums1[i+1:],nums1[i:])
			nums1[i] = nums2[j]
			m++
			i++
			j++
		}
	}
	// 有可能nums2中的没移完
	copy(nums1[m:],nums2[j:])
	// if i != m {
	// 	result = append(result,nums1[i:]...)
	// }
	// if j != n {
	// 	result = append(result,nums2[j:]...)
	// }
	// copy(nums1,result)
}
// @lc code=end

