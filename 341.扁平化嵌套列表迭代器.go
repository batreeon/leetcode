/*
 * @lc app=leetcode.cn id=341 lang=golang
 *
 * [341] 扁平化嵌套列表迭代器
 */

// @lc code=start
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

 type NestedIterator struct {
    list []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	// 直接在构造函数里，把原结构展开为数组
    list := []int{}
	var constructor func(nestedList []*NestedInteger)
	constructor = func(nestedList []*NestedInteger) {
		for _,ni := range nestedList {
			if (*ni).IsInteger() {
				list = append(list,(*ni).GetInteger())
			}else{
				nlist := (*ni).GetList()
				constructor(nlist)
			}
		}
		return
	}

	constructor(nestedList)
	return &NestedIterator{list}
}

func (this *NestedIterator) Next() int {
    ans := this.list[0]
	this.list = this.list[1:]
	return ans
}

func (this *NestedIterator) HasNext() bool {
    return len(this.list) > 0
}
// @lc code=end

