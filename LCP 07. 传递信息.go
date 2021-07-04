func numWays(n int, relation [][]int, k int) int {
	m := map[int][]int{}
    var f func(source int,kk int,res *int)
	f = func(source int,kk int,res *int) {
		if kk == 1 {
			for _,dst := range m[source] {
				if dst == n-1 {
					(*res)++
				}
			}
		}else{
			for _,dst := range m[source] {
				f(dst,kk-1,res)
			}
		}
	}
	for _,r := range relation {
		if _,ok := m[r[0]] ; !ok {
			m[r[0]] = []int{}
		}
		m[r[0]] = append(m[r[0]],r[1])
	}
	var result int
	f(0,k,&result)
	return result
}