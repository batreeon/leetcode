func permutation(s string) []string {
	// 其实还可以先对s按字符序排序，那么就不需要用map来存结果了，就和常规的找排列差不多了
	var tb func(ss string,visited []bool,tmp []byte,resultm *map[string]bool)
	tb = func(ss string,visited []bool,tmp []byte,resultm *map[string]bool) {
		if len(tmp) == len(ss) {
			tmpCopy := make([]byte,len(tmp))
			copy(tmpCopy,tmp)
			(*resultm)[string(tmpCopy)] = true
		}
		for i := 0 ; i < len(ss) ; i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			tmp = append(tmp,ss[i])
			tb(ss,visited,tmp,resultm)
			tmp = tmp[:len(tmp)-1]
			visited[i] = false
		}
	}
	resultm := map[string]bool{}
	tb(s,make([]bool,len(s)),[]byte{},&resultm)
	result := []string{}
	for k := range resultm {
		result = append(result,k)
	}
	return result
}