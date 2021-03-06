type bigHeap []int
func (h bigHeap) Len() int {
	return len(h)
}
func (h bigHeap) Less(i,j int) bool {
	return h[i] > h[j]
}
func (h bigHeap) Swap(i,j int) {
	h[i],h[j] = h[j],h[i]
}
func (h *bigHeap) Push(x interface{}) {
	*h = append(*h,x.(int))
}
func (h *bigHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}
func (h bigHeap) Top() interface{} {
	return h[0]
}

type smallHeap []int
func (h smallHeap) Len() int {
	return len(h)
}
func (h smallHeap) Less(i,j int) bool {
	return h[i] < h[j]
}
func (h smallHeap) Swap(i,j int) {
	h[i],h[j] = h[j],h[i]
}
func (h *smallHeap) Push(x interface{}) {
	*h = append(*h,x.(int))
}
func (h *smallHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}
func (h smallHeap) Top() interface{} {
	return h[0]	//顶是下标0
}