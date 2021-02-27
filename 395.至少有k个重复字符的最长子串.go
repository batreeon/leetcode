/*
 * @lc app=leetcode.cn id=395 lang=golang
 *
 * [395] 至少有K个重复字符的最长子串
 */

// @lc code=start
//看的leetcode 宫水三叶的题解
func longestSubstring(s string, k int) int {
	n := len(s)
	ans := 0
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	for p := 1 ; p <= 26 ; p++ {
		//cnt用来记录每个窗口中，各字母出现的次数
		cnt := make([]int,26)
		//total,sum分别表示窗口中字母种类数和符合条件的字母种类数
		for l,r,total,sum := 0,0,0,0 ; r < n ; r++ {
			i := s[r]-'a'
			cnt[i]++
			if cnt[i] == 1 {
				total++
			}
			if cnt[i] == k {
				sum++
			}
			for total > p {
				li := s[l]-'a'
				cnt[li]--
				if cnt[li] == 0 {
					total--
				}
				if cnt[li] == k-1 {
					sum--
				}
				l++
			}
			if total == sum {//没有不符合条件的字母
				ans = max(ans,r-l+1)
			}
		}
	}
	return ans
}
//共有26种字母，分别考察1~26种，每种遍历一次s,时间复杂度位O(26n)=O(n)
//每找到一种就与ans比较取大
// @lc code=end

