/*
 * @lc app=leetcode.cn id=1286 lang=golang
 *
 * [1286] 字母组合迭代器
 */

// @lc code=start
package main

type CombinationIterator struct {
	ci []string
}


func Constructor(characters string, combinationLength int) CombinationIterator {
	ans := []string{}
	var constructor func(res []byte,characters string,length int)
	constructor = func(res []byte,characters string,length int) {
		if length == 0 {
			resCopy := make([]byte,len(res))
			copy(resCopy,res)
			ans = append(ans,string(resCopy))
			return
		}
		n := len(characters)
		for i := 0 ; i < n-length+1 ; i++ {
			res = append(res,characters[i])
			constructor(res,characters[i+1:],length-1)
			res = res[:len(res)-1]
		}
	}

	constructor([]byte{},characters,combinationLength)
	return CombinationIterator{ans}
}


func (this *CombinationIterator) Next() string {
	ans := this.ci[0]
	this.ci = this.ci[1:]
	return ans
}


func (this *CombinationIterator) HasNext() bool {
	return len(this.ci) > 0
}


/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

