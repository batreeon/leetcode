/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	nums1,nums2 := [26]int{},[26]int{}
	l1,l2 := len(s1),len(s2)
	for _,v := range s1 {
		nums1[v-'a']++
	}
	left,right := 0,0
	for ; right < l2 ; right++ {
		nums2[s2[right]-'a']++
		if right-left+1 > l1{
			nums2[s2[left]-'a']--
			left++
		}
		if nums1 == nums2 {
			return true
		}
	}
	return false
}
// @lc code=end

