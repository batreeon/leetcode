package main

import "strconv"

func translateNum(num int) int {
	// if num < 10 {
	// 	return 1
	// }
	// if num % 100 > 25 {
	// 	return translateNum(num/10)
	// }else{
	// 	return translateNum(num/100) + translateNum(num/10)
	// }
	str := strconv.Itoa(num)
	if len(str) < 2 {
		return 1
	}
	f := make([]int,len(str))
	f[0] = 1
	if str[:2] <= "25" && str[:2] >= "10" {
		f[1] = 2
	}else{
		f[1] = 1
	}
	for i := 2 ; i < len(str) ; i++ {
		f[i] += f[i-1]
		if str[i-1:i+1] <= "25" && str[i-1:i+1] >= "10" {
			f[i] += f[i-2]
		}
	}
	return f[len(str)-1]
}