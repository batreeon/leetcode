/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1,l2 := len(nums1),len(nums2)
	i1,i2 := 0,0
	min := func(a,b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i1 < l1 && i2 < l2 && i1+i2 < (l1+l2)/2 {
		if nums1[i1] <= nums2[i2] {
			i1++
		}else{
			i2++
		}
	}
	if i1+i2 < (l1+l2)/2 {
		if i1 < l1 {
			for i1+i2 < (l1+l2)/2 {
				i1++
			}
		}else{
			for i1+i2 < (l1+l2)/2 {
				i2++
			}
		}
	}
	//i1,i2对应的值都还没有访问

	var ans float64
	if (l1+l2)%2 == 1 {	//总数为奇数
		if i1 == l1 {
			ans = float64(min(nums1[i2],nums2[i2+1]))
		}else if i2 == l2 {
			ans = float64(min(nums1[i1],nums2[i1+1]))
		}else{
			ans = float64(min(nums1[i1],nums2[i2]))
		}
	}else{
		if nums1[i1] < nums2[i2] {
			ans += float64(nums1[i1])
			i1++
		}else{
			ans += float64(nums2[i2])
			i2++
		}
		ans += float64(min(nums1[i1],nums2[i2]))
		ans = ans/2.0
	}
	return ans
}
// @lc code=end

