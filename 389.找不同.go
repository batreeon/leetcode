/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	bs := []byte(s)
	mp := map[byte]int{}
	for _,ss := range bs {
		mp[ss]++
	}
	bt := []byte(t)
	var ans byte
	for _,tt := range bt {
		if mp[tt] > 0 {
			mp[tt]--
		}else{
			ans = tt
		}
	}
	return ans
}
// @lc code=end

