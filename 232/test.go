package main

import "fmt"

func main() {
	mq := Constructor()
	mq.Push(1)
	mq.Push(2)
	fmt.Println(mq.Peek())
	fmt.Println(mq.Pop())
	fmt.Println(mq.Empty())
}