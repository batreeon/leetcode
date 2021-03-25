package main

import "fmt"

func main() {
	cI := Constructor(string("abc"), 2)
	fmt.Println(cI.Next())
	fmt.Println(cI.Next())
	fmt.Println(cI.Next())
	fmt.Println(cI.HasNext())
}