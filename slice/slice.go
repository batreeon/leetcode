package main

import "fmt"

func main() {
	s := make([]int, 0, 5)
	fmt.Println(len(s))

	s = append(s,1)
	fmt.Println(len(s))

	// copy(s[1:2],s[0:1])
	// s[0] = 2
	s = append([]int{2},s...)
	fmt.Println(len(s))
	// 这是s的cap就是在[]int{2}基础上计算的了
	fmt.Println(len(s))
}