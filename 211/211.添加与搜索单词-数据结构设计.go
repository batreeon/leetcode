/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] 添加与搜索单词 - 数据结构设计
 */

// @lc code=start
package main

type WordDictionary struct {
	child []*WordDictionary
	isword bool
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
	child := make([]*WordDictionary,26)
	return WordDictionary{child:child}
}


func (this *WordDictionary) AddWord(word string)  {
	p := this
	for _,w := range word {
		if p.child[int(w-'a')] == nil {
			p.child[int(w-'a')] = &WordDictionary{child:make([]*WordDictionary,26)}
		}
		p = p.child[int(w-'a')]
	}
	p.isword = true
}


func (this *WordDictionary) Search(word string) bool {
	if word == "" {
		if this.isword {
			return true
		}
		return false
	}
	w := word[0]
	if w == '.' {
		for _,child := range this.child {
			if child != nil {
				if child.Search(word[1:]) {
					return true
				}
			}
		}
	}else{
		if this.child[int(w-'a')] != nil && this.child[int(w-'a')].Search(word[1:]) {
			return true
		}
	}
	return false
}

// func main() {
// 	this := Constructor()
// 	this.AddWord("bad")
// 	this.AddWord("dad")
// 	this.AddWord("mad")
// 	this.Search("pad")
// 	this.Search("bad")
// 	this.Search(".ad")
// 	this.Search("b..")
// }

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end

