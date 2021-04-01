type stack struct {
	nums []int
	l    int
}

func (st *stack) push(x int) {
	st.nums = append(st.nums, x)
	st.l++
}

func (st *stack) top() int {
	return st.nums[st.l-1]
	st.l--
}

func (st *stack) pop() {
	st.nums = st.nums[:st.l-1]
}
func (st stack) isEmpty() bool {
	return st.l == 0
}