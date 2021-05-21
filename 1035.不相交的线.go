/*
 * @lc app=leetcode.cn id=1035 lang=golang
 *
 * [1035] 不相交的线
 */

// @lc code=start
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	//啊啊啊，看了答案才反应过来就是lcs问题 
	f1 := make([]int,len(nums2)+1)
	for i := 1 ; i <= len(nums1) ; i++ {
		f := make([]int,len(nums2)+1)
		for j := 1 ; j <= len(nums2) ; j++ {
			if nums1[i-1] == nums2[j-1] {
				f[j] = f1[j-1] + 1
			}else if f[j-1] > f1[j] {
				f[j] = f[j-1]
			}else{
				f[j] = f1[j]
			}
		}
		f1 = f
	}
	return f1[len(nums2)]
}
// @lc code=end

