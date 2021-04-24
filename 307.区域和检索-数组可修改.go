/*
 * @lc app=leetcode.cn id=307 lang=golang
 *
 * [307] 区域和检索 - 数组可修改
 */

// @lc code=start
type NumArray struct {
	/*
	nums *[]int
	sum *[]int
	change *[]int
	*/
	nums *[]int
}


func Constructor(nums []int) NumArray {
	
	/*
	n := len(nums)
	sum := make([]int,n)
	sum[0] = nums[0]
	for i := 1 ; i < n ; i++ {
		sum[i] = sum[i-1] + nums[i]
	}

	changes := []int
	return NumArray{nums:&nums,sum:&sum,changes:changes}
	*/

	return NumArray{nums:&nums}
}


func (this *NumArray) Update(index int, val int)  {
	
	/*
	change := val - (*this.nums)[index]
	(*this.nums)[index] = val
	n := len(*this.nums)
	*/

	(*this.nums)[index] = val
}


func (this *NumArray) SumRange(left int, right int) int {
	

	/*
	if left != 0 {
		return (*this.sum)[right] - (*this.sum)[left-1]
	}else{
		return (*this.sum)[right]
	}
	*/

	sum := 0
	for i := left ; i <= right ; i++ {
		sum += (*this.nums)[i]
	}
	return sum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
// @lc code=end

