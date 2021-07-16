func singleNumbers(nums []int) []int {
    temp := 0
	for _,num := range nums {
		temp ^= num
	}
	result := []int{temp,temp}
	// 得到最低位1
	temp = (temp&(temp-1))^temp
	for _,num := range nums {
		if temp&num == 0 {
			result[0] ^= num
		}else{
			result[1] ^= num
		}
	}
	return result
}