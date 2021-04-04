package main
/*
一个句子是由一些单词与它们之间的单个空格组成，且句子的开头和结尾没有多余空格。比方说，"Hello World" ，"HELLO" ，"hello world hello world" 都是句子。每个单词都 只 包含大写和小写英文字母。

如果两个句子 sentence1 和 sentence2 ，可以通过往其中一个句子插入一个任意的句子（可以是空句子）而得到另一个句子，那么我们称这两个句子是 相似的 。比方说，sentence1 = "Hello my name is Jane" 且 sentence2 = "Hello Jane" ，我们可以往 sentence2 中 "Hello" 和 "Jane" 之间插入 "my name is" 得到 sentence1 。

给你两个句子 sentence1 和 sentence2 ，如果 sentence1 和 sentence2 是相似的，请你返回 true ，否则返回 false 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sentence-similarity-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
import "strings"

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
    l1, l2 := len(sentence1), len(sentence2)
	if l1 < l2 {
		sentence1, sentence2 = sentence2, sentence1
        l1,l2 = l2,l1
	}
    if strings.HasPrefix(sentence1,sentence2) {
		if l1 == l2 {
			return true
		}else if sentence1[l2] == ' ' {
			return true
		}else{
			return false
		}
	}
	if strings.HasSuffix(sentence1,sentence2) {
		if l1 == l2 {
			return true
		}else if sentence1[l1-1-l2] == ' ' {
			return true
		}else{
			return false
		}
	}
	board := 0
	for i := 0; i < l2; i++ {
		if strings.HasPrefix(sentence1,sentence2[:i+1]) {
			continue
		}
		board = i
		break
	}
	if board == l2 {
		return true
	}
	if board == 0 {
		if strings.HasSuffix(sentence1,sentence2) {
			return true
		}
		return false
	}else{
		if sentence1[board-1] != ' ' {
			return false
		}else{
			if strings.HasSuffix(sentence1[board:],sentence2[board:]) {
				return true
			}
		}
	}
	return false
}