/*
 * @lc app=leetcode.cn id=3227 lang=golang
 *
 * [3227] 字符串元音游戏
 */

// @lc code=start
func doesAliceWin(s string) bool {
	// 当s有0个元音字母时，小红输
	// 当s有奇数个元音字母时，小红赢
	// 当s有偶数个元音字母时，小红移除奇数个，然后剩奇数个，小明再移除偶数个还剩奇数个，小红又能赢了
	// 所以只要s中含有元音字母，小红就能赢
    yuan := map[rune]struct{}{
		'a':struct{}{},
		'e':struct{}{},
		'i':struct{}{},
		'o':struct{}{},
		'u':struct{}{},
	}

	for _, r := range s {
		if _, ok := yuan[r]; ok {
			return true
		}
	}

	return false
}


// @lc code=end

