/*
 * @lc app=leetcode.cn id=318 lang=golang
 *
 * [318] 最大单词长度乘积
 */

// @lc code=start
func maxProduct(words []string) int {
	// 掩码：小写字母最多只有26个，我们用数的低26位记录word含有的字母
	// masks记录掩码对应的最大单词长度，也就是说掩码相同的只记录最长的那个单词
	masks := map[int]int{}
	for _, w := range words {
		l := len(w)
		mask := 0
		for _, letter := range w {
			mask |= (1 << (letter - 'a'))
		}
		if length, ok := masks[mask]; !ok || length < l {
			masks[mask] = l
		}
	}

	result := 0
	for m1, l1 := range masks {
		for m2, l2 := range masks {
			if m1 & m2 == 0 && result < l1 * l2 {
				result = l1 * l2
			}
		}
	}

	return result
}
// @lc code=end

