package main

import (
	"fmt"
	"sort"
)

func numberOfWeakCharacters(properties [][]int) int {
	n := len(properties)
	sort.Slice(properties,func(i,j int) bool {return properties[i][0] < properties[j][0] || (properties[i][0] == properties[j][0] && properties[i][1] < properties[j][1])})
    
    properties = reverse(properties)

	now := -1
	result := 0
	for i := 0 ; i < n ; i++ {
		j := i
		for j+1 < n && properties[j+1][0] == properties[i][0] {
			j++
		}
		for k := i ; k <= j ; k++ {
			if properties[k][1] < now {
				result++
			}
		}
        for k := i ; k <= j ; k++ {
			now = max(now,properties[k][1])
		}
		i = j
	}
	return result
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func reverse(a [][]int) [][]int {
    for i,j := 0,len(a)-1 ; i < j ; i,j = i+1,j-1 {
        a[i],a[j] = a[j],a[i]
    }
    return a
}

func main() {
	p := [][]int{{7,9},{10,7},{6,9},{10,4},{7,5},{7,10}}
	res := numberOfWeakCharacters(p)
	fmt.Println(res)
}