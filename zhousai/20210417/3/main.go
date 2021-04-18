package main

import "math"

func getMaximumXor(nums []int, maximumBit int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	n := len(nums)
    result := []int{}
    
    res := math.Pow(2,float64(maximumBit))-1
	mask := res
	for i := n - 1; i >= 0; i-- {
        result = append(result,int(res)^int(mask)&xor)
		xor ^= nums[i]
	}

	return result
}