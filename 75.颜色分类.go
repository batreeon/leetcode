/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
type stack struct {
	s []int
	l int
}

func (st *stack) push(x int) {
	st.s = append(st.s, x)
	st.l++
}
func (st *stack) top() int {
	return st.s[st.l-1]
}
func (st *stack) pop() {
	st.s = st.s[:st.l-1]
	st.l--
}
func (st stack) isEmpty() bool {
	return st.l == 0
}

func sortColors(nums []int) {
	/*
		// 20210327单调栈尝试
		// 不行，比如[1,2,0]就不行
		ans := []int{}
		s := new(stack)
		for _,num := range nums {
			if s.isEmpty() {
				s.push(num)
			}else{
				for s.top() < num {
					ans = append(ans,s.top())
					s.pop()
					if s.isEmpty() {
						break
					}
				}
				s.push(num)
			}
		}
		for !(s.isEmpty()) {
			ans = append(ans,s.top())
			s.pop()
		}
		copy(nums,ans)
	*/

	// 记录排序后0和1的末尾
	index0end := -1
	index1end := -1
	for _, num := range nums {
		if num == 0 {
			index0end++
			index1end++
		}
		if num == 1 {
			index1end++
		}
	}
	l := len(nums)
	// 交换元素，将0归位
	for index0, i := -1, 0; index0 != index0end && i < l; i++ {
		if nums[i] == 0 {
			index0++
			nums[i], nums[index0] = nums[index0], nums[i]
		}
	}
	// 交换元素将1归位
	for index1, i := index0end, index0end+1; index1 != index1end && index1 < l; i++ {
		if nums[i] == 1 {
			index1++
			nums[i], nums[index1] = nums[index1], nums[i]
		}
	}
}

// @lc code=end

