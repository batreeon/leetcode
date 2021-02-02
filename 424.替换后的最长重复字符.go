/*
 * @lc app=leetcode.cn id=424 lang=golang
 *
 * [424] 替换后的最长重复字符
 */

// @lc code=start
func characterReplacement(s string, k int) int {
	left,right := 0,0
	cnt := make([]int,26)
	ans := 0
	maxcnt := 0

	var max func(x,y int) int
	max = func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}

	for ; right < len(s) ; right++ {
		cnt[s[right]-'A']++
		maxcnt = max(maxcnt,cnt[s[right]-'A'])
		for right-left+1-maxcnt > k {
			cnt[s[left]-'A']--
			left++
		}
		ans = max(ans,right-left+1)
	}
	return ans
}
// @lc code=end

