/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

// @lc code=start
func near(w1, w2 string) bool {
	dif := 0
	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			dif++
			if dif > 1 {
				return false
			}
		}
	}
	if dif == 0 {
		return false
	}
	return true
}

func ladderLength(beginWord string, endWord string, wordList []string) int {

	//nearwordmap := make(map[string][]string)
	wait := []string{beginWord}
	var lastword string
	ans := 1
	n := 0
	for len(wait) != 0 {
		waitword := wait[0]
		for _, w := range wordList {
			if w != lastword && near(waitword, w) {
				//append(nearwordmap[wait],w)
				if near(w, endWord) {
					return ans + 1
				}
				wait = append(wait, w)
				n++
			}
		}
		lastword = waitword
		wait = wait[1:]
	}
	return 0
}

// @lc code=end

