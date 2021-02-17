package main

import (
	"fmt"
)

func main() {
	a := [][]int{}
	a = append(a,[]int{})
	a[0] = append(a[0],[]int{1,2,3}...)
	a = append(a,[]int{})
	a[1] = append(a[1],[]int{4,5,6}...)
	fmt.Println(a[2])//报错，和python不一样
}