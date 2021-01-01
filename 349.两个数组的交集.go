/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	//比较两个集合大小是为了把较小的用map统计
	var ans []int
	shortnum := make(map[int]bool)
	if len(nums1) < len(nums2) {
		for _, v := range nums1 {
			shortnum[v] = true
		}
	} else {
		for _, v := range nums2 {
			shortnum[v] = true
		}
	}
	if len(nums1) < len(nums2) {
		for _, v := range nums2 {
			if _, ok := shortnum[v]; ok {
				ans = append(ans, v)
				delete(shortnum, v)
				//重复的元素只算一次，很奇怪，这是交集，但题目就是这样给的，没办法
			}
		}
	} else {
		for _, v := range nums1 {
			if _, ok := shortnum[v]; ok {
				ans = append(ans, v)
				delete(shortnum, v)
			}
		}
	}
	return ans
}

// @lc code=end

