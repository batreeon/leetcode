/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	next := -1
	for i := 0 ; i < len(nums2) ; i++ {
		for j := i+1 ; j < len(nums2) ; j++ {
			if nums2[j] > nums2[i] {
				next = nums2[j]
				break
			}
		}
		m[nums2[i]] = next
		next = -1
	}

	result := make([]int,len(nums1))
	for i := 0 ; i < len(nums1) ; i++ {
		result[i] = m[nums1[i]]
	}

	return result
}
// @lc code=end

