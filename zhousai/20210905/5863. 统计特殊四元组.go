func countQuadruplets(nums []int) int {
	n := len(nums)
	result := 0
	var sum int
	for i := 0 ; i < n ; i++ {
		sum += nums[i]
		for j := i+1 ; j < n ; j++ {
			sum += nums[j]
			for k := j+1 ; k < n ; k++ {
				sum += nums[k]
				for l := k+1 ; l < n ; l++ {
					if sum == nums[l] {
						result++
					}
				}
				sum -= nums[k]
			}
			sum -= nums[j]
		}
		sum -= nums[i]
	}

	return result
}