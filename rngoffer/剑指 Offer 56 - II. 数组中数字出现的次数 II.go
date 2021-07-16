func singleNumber(nums []int) int {
	result := 0
	for i := 0 ; i < 64 ; i++ {
		count := 0
		for _,num := range nums {
			if (1<<i)&num != 0 {
				count++
			}
		}
		result ^= (1<<i)*(count%3)
	}
	return result
}