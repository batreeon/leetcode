func findMiddleIndex(nums []int) int {
	n := len(nums)
	sum := make([]int,n+1)
	for i := 1 ; i <= n ; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	for i := 1 ; i < n+1 ; i++ {
		if sum[i] == sum[n] - sum[i-1] {
			return i-1
		}
	}
	return -1
}