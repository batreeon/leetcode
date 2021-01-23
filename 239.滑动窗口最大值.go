/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
// func maxSlidingWindow(nums []int, k int) []int {
// 	partNums := []int{}
// 	ans := []int{}
// 	n := len(nums)
// 	for i := 0 ; i < n-k+1 ; i++ {
// 		partNums = nums[i:i+k]
// 		ans = append(ans,max(partNums))
// 	}
// 	return ans
// }

// func max(nums []int) int{
// 	max := nums[0]
// 	for _,v := range nums {
// 		if v > max {
// 			max = v
// 		}
// 	}
// 	return max
// }

var a []int
type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
//重新实现了Less函数，用来建最大堆，这个堆里的元素记录的不是实值，而是元素在a中的下标
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func maxSlidingWindow(nums []int, k int) []int {
    a = nums
    q := &hp{make([]int, k)}
    for i := 0; i < k; i++ {
        q.IntSlice[i] = i
    }
    heap.Init(q)

    n := len(nums)
    ans := make([]int, 1, n-k+1)
    ans[0] = nums[q.IntSlice[0]]
    for i := k; i < n; i++ {//i是右边界
        heap.Push(q, i)
        for q.IntSlice[0] <= i-k {
            heap.Pop(q)
		}//上一个窗口的最大值还在窗口中，就不删除它，否则就删掉它
		//太秀了，好秀的解答
        ans = append(ans, nums[q.IntSlice[0]])
    }
    return ans
}
// @lc code=end

