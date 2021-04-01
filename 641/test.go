package main

import "fmt"

func main() {
	
	mcd := Constructor(3)
	mcd.InsertLast(1)
	mcd.InsertLast(2)
	mcd.InsertFront(3)
	mcd.InsertFront(4)
	fmt.Println(mcd.GetRear())

}