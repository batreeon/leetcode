/*
 * @lc app=leetcode.cn id=2325 lang=golang
 *
 * [2325] 解密消息
 */

// @lc code=start
func decodeMessage(key string, message string) string {
	m := make(map[byte]int)
    // space用来记录空格和重复字母数，这俩都会导致字母下标不等于其次序
	space := 0;
	for i,letter := range key {
		if byte(letter) == ' ' {
			space++
			continue
		}
		if _, ok := m[byte(letter)]; !ok {
			m[byte(letter)] = i-space;
		}else{
			space++
		}
	}

	result := make([]byte, len(message))
	for i, letter := range message {
		if byte(letter) == ' ' {
			result[i] = ' '
			continue
		}
		result[i] = byte(m[byte(letter)] + int('a'))
	}

	return string(result)
}
// @lc code=end

