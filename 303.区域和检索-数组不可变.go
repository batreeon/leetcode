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
	// new := NumArray{make([]int,len(nums))}
	// copy(new.nA[:],nums[:])
	// return new

	//前缀和
	n := len(nums)
	new := NumArray{make([]int,n+1)}
	for i := 0 ; i < n ; i++ {
		new.nA[i+1] = new.nA[i] + nums[i]	
		//为什么不直接new.nA[i] = new.nA[i-1] + nums[i]
		//这样也可以，但需要多写几条语句来处理当i==0时的情况
	}
	return new
}


func (this *NumArray) SumRange(i int, j int) int {
	// ans := this.nA[i]
	// for k := i+1 ; k <= j ; k++ {
	// 	ans += this.nA[k]
	// }
	// return ans

	//前缀和
	return this.nA[j+1]-this.nA[i]	//前j项和-前i-1项和
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
// @lc code=end

