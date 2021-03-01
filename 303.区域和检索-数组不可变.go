/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */

// @lc code=start
type NumArray struct {
	nA []int
}


func Constructor(nums []int) NumArray {
	new := NumArray{make([]int,len(nums))}
	copy(new.nA[:],nums[:])
	return new
}


func (this *NumArray) SumRange(i int, j int) int {
	ans := this.nA[i]
	for k := i+1 ; k <= j ; k++ {
		ans += this.nA[k]
	}
	return ans
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
// @lc code=end

