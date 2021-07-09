func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	times := 0
	result := -1
	for _,num := range nums {
		if times == 0 {
			result = num
			times = 1
		}else if result == num {
			times++
		}else {
			times--
		}
	}
	times = 0
	for _,num := range nums {
		if num == result {
			times++
		}
	}
	if times*2 <= len(nums) {
		return -1
	}
	return result
}