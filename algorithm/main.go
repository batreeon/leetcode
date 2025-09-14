package main

import (
	"fmt"
	"leetcode/algorithm/lru"
)

func main() {
	l := lru.NewLru(2)
	fmt.Println(l.Get(1))
	l.Set(1, 1)
	fmt.Println(l.Get(1))
	l.Set(2, 2)
	l.Set(3, 3)
	fmt.Println(l.Get(1))
}
