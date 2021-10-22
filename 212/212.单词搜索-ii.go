/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */

// @lc code=start
package main

var orix = []int{-1, 0, 0, 1}
var oriy = []int{0, -1, 1, 0}
var r, c int

func findWords(board [][]byte, words []string) []string {
	result := []string{}
	r, c = len(board), len(board[0])

	for _, word := range words {
		b := word[0]

	loop:
		for i := 0; i < r; i++ {
			for j := 0; j < c; j++ {
				if board[i][j] == b {
					if findWord(board, i*c+j, word[1:], map[int]bool{i*c + j: true}) {
						wordCopy := word
						result = append(result, wordCopy)
						break loop
					}
				}
			}
		}
	}

	return result
}

func findWord(board [][]byte, begin int, word string, seen map[int]bool) bool {
	if word == "" {
		return true
	}
	b := word[0]
	x, y := begin/c, begin%c
	for i := 0; i < 4; i++ {
		newX, newY := x+orix[i], y+oriy[i]
		if newX >= 0 && newX < r && newY >= 0 && newY < c && !seen[newX*c+newY] && board[newX][newY] == b {
			seen[newX*c+newY] = true
			if findWord(board, newX*c+newY, word[1:], seen) {
				return true
			}
			delete(seen, newX*c+newY)
		}
	}
	return false
}

// @lc code=end
