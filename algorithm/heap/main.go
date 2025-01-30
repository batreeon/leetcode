package main

// 算法导论最大堆实现
import "fmt"

type heap []int

// 用一个数组来存放堆
// 获取下标i的父亲，左儿子，右儿子
func (h heap) Parent(i int) int {
	parent := (i-1)/2
	return parent
	// if parent >= 0 {
	// 	return parent
	// }
	// return -1
}

func (h heap) Left(i int) int {
	left := i*2 + 1
	return left
}

func (h heap) Right(i int) int {
	right := i*2 + 2
	return right
}

func (h heap) MaxHeapify(i int) {
	// 前提，节点i的左右两棵子树都符合堆的特性
	l, r := h.Left(i), h.Right(i)
	largest := i
	if l < len(h) && h[l] > h[largest] {
		largest = l
	}
	if r < len(h) && h[r] > h[largest] {
		largest = r
	}
	if largest != i {
		h[i], h[largest] = h[largest], h[i]
		// 以largest为根的子树可能不符合最大堆了
		h.MaxHeapify(largest)
	}
}

func (h heap) BuildMaxHeap() {
	// 叶子节点自然符合堆的性质，只需要考察非叶子结点
	for i := len(h)/2 - 1; i >= 0; i-- {
		// i是从大到小遍历的，所以节点i的左右子树一定是符合堆的性质的
		h.MaxHeapify(i)
	}
}

func (h heap) HeapSort() []int {
	h.BuildMaxHeap()
	result := make([]int, len(h))
	for i := len(result) - 1; i >= 0; i-- {
		result[i] = h[0]
		h[0], h[i] = h[i], h[0]
		h = h[:len(h)-1]
		// 每次是把h[0]和h[len(h)-1]交换了一下，所以现在h[0]的两个儿子子树还都符合堆的性质
		h.MaxHeapify(0)
	}
	return result
}

// 下面两个要用*heap，因为要改变源结构
func (h *heap) Pop() int {
	pop := (*h)[0]
	(*h)[0], (*h)[len(*h)-1] = (*h)[len(*h)-1], (*h)[0]
	(*h) = (*h)[:len(*h)-1]
	// 左子树和右子树都是最大堆，所以只需要调整根节点
	(*h).MaxHeapify(0)
	return pop
}

func (h *heap) Push(x int) {
	(*h) = append((*h), x)
	for i := len(*h) - 1; i > 0 && (*h)[(*h).Parent(i)] < (*h)[i]; {
		(*h)[i], (*h)[(*h).Parent(i)] = (*h)[(*h).Parent(i)], (*h)[i]
		// i节点为根的子树始终是最大堆，i的父节点可能不满足最大堆性质
		// 但i的父节点一定是比i的子孙大的，因为在引入新节点前，整个树是符合最大堆的
		i = (*h).Parent(i)
	}
}

func main() {
	h := heap([]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7})
	h.BuildMaxHeap()
	fmt.Println(h.Pop())
	h.Push(11)
	h.Push(6)
	for len(h) > 0 {
		fmt.Println(h.Pop())
	}

	// fmt.Println(h.HeapSort())
	// HeapSort后，h就是有序序列，就不再是最大堆了
	// fmt.Println(h)	//h本身并没有被改变
}