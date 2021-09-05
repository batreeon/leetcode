func gcdSort(nums []int) bool {
	n := len(nums)
	numsCopy := make([]int,n)
	copy(numsCopy,nums)
	sort.Ints(numsCopy)
	index := make([]int,n)
	
}
func gcd(x,y int) int {
	if x < y {
		x,y = y,x
	}
	if y == 0 {
		return x
	}
	return gcd(y,x%y)
}