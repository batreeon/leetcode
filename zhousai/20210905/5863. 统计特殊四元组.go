func countQuadruplets(nums []int) int {
	n := len(nums)
	result := 0
	
	for i := 0 ; i < n ; i++ {
		for j := i+1 ; j < n ; j++ {
			for k := j+1 ; k < n ; k++ {
				for l := k+1 ; l < n ; l++ {
					if nums[i]+nums[j]+nums[k] == nums[l] {
						result++
					}
				}
			}
		}
	}

	return result
}