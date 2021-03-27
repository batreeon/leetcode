type stack struct {
	s []int
	l int
}

func (st *stack) push(x int) {
	st.s = append(st.s,x)
	st.l++
}

func (st *stack) top() int {
	return st.s[st.l-1]
	st.l--
}

func (st *stack) pop() {
	st.s = st.s[:st.l-1]
}
func(st stack) isEmpty() bool {
	return st.l == 0
}