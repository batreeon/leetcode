/*
 * @lc app=leetcode.cn id=284 lang=golang
 *
 * [284] 顶端迭代器
 */

// @lc code=start
/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *       
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
    iter *Iterator
	haspeek bool
	peekValue int
}

func Constructor(iter *Iterator) *PeekingIterator {
    return &PeekingIterator{iter,false,0}
}

func (this *PeekingIterator) hasNext() bool {
	if this.haspeek {
		return true
	}
    return this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
    if this.haspeek {
		this.haspeek = false
		return this.peekValue
	}
	return this.iter.next()
}

func (this *PeekingIterator) peek() int {
    if !this.haspeek {
		this.peekValue = this.iter.next()
		this.haspeek = true
	}
	return this.peekValue
}
// @lc code=end

