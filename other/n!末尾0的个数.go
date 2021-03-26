package main

import "fmt"

func findZeroAtTail(n int) int {
	five := 5
	ans := 0
	for n >= five {
		ans += n / five
		five *= 5
	}
	return ans
}

func main() {
	fmt.Println(findZeroAtTail(25))
}