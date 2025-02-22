/*
 * @lc app=leetcode.cn id=2506 lang=golang
 *
 * [2506] 统计相似字符串对的数目
 */

// @lc code=start
func similarPairs(words []string) int {
    m := map[[26]int]int{}
	for _, word := range words {
		letters := [26]int{}
		for _, letter := range word {
			letters[letter - 'a'] = 1
		}
		m[letters]++
	}

	res := 0
	for _, v := range m {
		if v > 1 {
			res += v*(v-1)/2
		}
	}
	return res
}
// @lc code=end

