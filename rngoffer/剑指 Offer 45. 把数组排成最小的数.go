package main

import (
	"sort"	
	"strconv"
)


func minNumber(nums []int) string {
	sort.Slice(nums,func(i,j int) bool {
		s1,s2 := strconv.Itoa(nums[i]),strconv.Itoa(nums[j])
		num1,_ := strconv.Atoi(s1+s2)
		num2,_ := strconv.Atoi(s2+s1)
		return num1 < num2
	})
	var result string
	for _,num := range nums {
		numS := strconv.Itoa(num)
		result += numS
	}
	return result
}