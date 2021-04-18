package main

import "math"

func countPoints(points [][]int, queries [][]int) []int {
    isin := func(x1, y1, x2, y2, r int) bool {
		return math.Pow(float64(x1-x2),2)+math.Pow(float64(y1-y2),2) <= math.Pow(float64(r),2)
	}

	result := make([]int,len(queries))

	for i := 0 ; i < len(queries) ; i++ {
		for j := 0 ; j < len(points) ; j++ {
			if isin(points[j][0],points[j][1],queries[i][0],queries[i][1],queries[i][2]) {
				result[i]++
			}
		}
	}

	return result
}