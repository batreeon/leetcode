/*
 * @lc app=leetcode.cn id=966 lang=golang
 *
 * [966] 元音拼写检查器
 */

// @lc code=start

func spellchecker(wordlist []string, queries []string) []string {
	yuanyin := map[byte]struct{}{
		'a':struct{}{},
		'e':struct{}{},
		'i':struct{}{},
		'o':struct{}{},
		'u':struct{}{},
	}
	result := []string{}

	sameWordlist := make(map[string]string)
	lowerWordlist := make(map[string]string)
	transWordlist := make(map[string]string)

	for _, word := range wordlist {
		sameWordlist[word] = word

		lowerWord := strings.ToLower(word)
		if _, ok := lowerWordlist[lowerWord]; !ok {
			lowerWordlist[lowerWord] = word
		}
		
		tw := transWord(yuanyin, word)
		if _, ok := transWordlist[tw]; !ok {
			transWordlist[tw] = word
		}
	}

	for _, query := range queries {
		if w, ok := sameWordlist[query]; ok {
			result = append(result, w)
			continue
		}

		lq := strings.ToLower(query)
		if w, ok := lowerWordlist[lq]; ok {
			result = append(result, w)
			continue
		}

		tq := transWord(yuanyin, query)
		if w, ok := transWordlist[tq]; ok {
			result = append(result, w)
			continue
		}

		result = append(result, "")
	}

	return result


}

func transWord(yuanyin map[byte]struct{}, word string) string {
	word = strings.ToLower(word)
	bs := []byte(word)
	for i, b := range bs {
		if _, ok := yuanyin[b]; ok {
			bs[i] = 'a'
		}
	}

	return string(bs)
}

/*
func sc(wordlist []string, query string) string {
	s1, s2, s3 := "", "", ""
	yuanyin := map[byte]struct{}{
		'a':struct{}{},
		'e':struct{}{},
		'i':struct{}{},
		'o':struct{}{},
		'u':struct{}{},
	}

	for _, word := range wordlist {
		if len(word) != len(query) {
			continue
		}
		if word == query {
			return word
		}

		lowerWord, lowerQuery := strings.ToLower(word), strings.ToLower(query)
		if s2 == "" {
			if lowerWord == lowerQuery {
				s2 = word
			}
		}

		if s2 == "" && s3 == "" {
			flag := true
			for i := range len(query) {
				if lowerWord[i] == lowerQuery[i] {
					continue
				}
				_, ok1 := yuanyin[lowerWord[i]]
				_, ok2 := yuanyin[lowerQuery[i]]
				if !ok1 || !ok2 {
					flag = false
					break
				}
			}

			if flag {
				s3 = word
			}
		}
	}

	if s1 != "" {
		return s1
	}
	if s2 != "" {
		return s2
	}
	if s3 != "" {
		return s3
	}
	return ""
}
*/
// @lc code=end

