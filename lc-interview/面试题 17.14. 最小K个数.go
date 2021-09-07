func smallestK(arr []int, k int) []int {
	start,end := 0,len(arr)-1
	p := partition(arr,start,end)
	for p != k-1 {
		if p > k-1 {
			end = p-1
		}else{
			start = p+1
		}
		p = partition(arr,start,end)
	}
	return arr[:k]
}

func partition(arr []int,start,end int) int {
	q := start-1
	x := arr[end]
	for i := start ; i <= end ; i++ {
		if arr[i] < x {
			q++
			arr[q],arr[start] = arr[start],arr[q]
		}
	}
	arr[q+1],arr[end] = arr[end],arr[q+1]
	return q+1
}