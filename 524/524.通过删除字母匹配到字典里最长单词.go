/*
 * @lc app=leetcode.cn id=524 lang=golang
 *
 * [524] 通过删除字母匹配到字典里最长单词
 */

// @lc code=start
package main

import "sort"



func findLongestWord(s string, dictionary []string) string {
	letters := make([][]int,26)
	for i := 0; i < 26; i++ {
		letters[i] = []int{}
	}
	for i,b := range s {
		letters[int(b-'a')] = append(letters[int(b-'a')], i)
	}

	sort.Slice(dictionary, func(i,j int) bool {
		return len(dictionary[i]) > len(dictionary[j]) || (len(dictionary[i]) == len(dictionary[j]) && dictionary[i] < dictionary[j])
	})

	
	for _,word := range dictionary {
		i := 0
		maxIdx := -1
		for ; i < len(word); i++ {
			b := word[i]
			// if len(letters[int(b-'a')]) == 0 {
			// 	break
			// }
			if idx := sort.Search(len(letters[int(b-'a')]), func(j int)bool {return letters[int(b-'a')][j] > maxIdx}); idx != len(letters[int(b-'a')]) {
				maxIdx = letters[int(b-'a')][idx]
			}else{
				break
			}
		}
		if i == len(word) {
			return word
		}
	}

	return string("")
}
// @lc code=end

