package main

import (
	"sort"
)

func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties,func(i,j int) bool {return (properties[i][0] < properties[j][0] && properties[i][1] < properties[j][1]) || (properties[i][0] < properties[j][0]) || (properties[i][1] < properties[j][1])})
	// sort.Slice(properties,func(i,j int) bool {return (properties[i][0] < properties[j][0] && properties[i][1] < properties[j][1])})
	
	var result int
	n := len(properties)
	for i := 0 ; i < n ; i++ {
        if sort.Search(n-1-i, func(j int) bool {return properties[i][0] < properties[i+1:][j][0] && properties[i][1] < properties[i+1:][j][1]}) != n-1-i {
			result++
		}
	}
	return result
}