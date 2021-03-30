/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
// package main
// import "strings"

// 时间慢，但空间还好
// type Trie struct {
// 	m map[string]bool
// }


// /** Initialize your data structure here. */
// func Constructor() Trie {
// 	m := map[string]bool{}
// 	return Trie{m:m}
// }


// /** Inserts a word into the trie. */
// func (this *Trie) Insert(word string)  {
// 	this.m[word] = true
// }


// /** Returns if the word is in the trie. */
// func (this *Trie) Search(word string) bool {
// 	if _,ok := this.m[word] ; ok {
// 		return true
// 	}
// 	return false
// }


// /** Returns if there is any word in the trie that starts with the given prefix. */
// func (this *Trie) StartsWith(prefix string) bool {
// 	for key,_ := range this.m {
// 		if strings.HasPrefix(key,prefix) {
// 			return true
// 		}
// 	}
// 	return false
// }

// 空间换时间
type Trie struct {
	child []*Trie
	isWord bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
	child := make([]*Trie,26)
	return Trie{child:child}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	p := this
	for _,w := range word {
		if p.child[int(w-'a')] == nil {
			p.child[int(w-'a')] = &Trie{child:make([]*Trie,26)}
		}
		p = p.child[int(w-'a')]
	}
	p.isWord = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	p := this
	for _,w := range word {
		if p.child[int(w-'a')] == nil {
			return false
		}
		p = p.child[int(w-'a')]
	}
	if p.isWord {
		return true
	}
	return false
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	p := this
	for _,w := range prefix {
		if p.child[int(w-'a')] == nil {
			return false
		}
		p = p.child[int(w-'a')]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

