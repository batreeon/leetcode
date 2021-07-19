package main

func groupAnagrams(strs []string) [][]string {
    m := map[[26]int][]string{}
    getKey := func(str string) [26]int {
        res := [26]int{}
		for _,b := range str {
			res[int(b-'a')]++
		}
		return res
    }
	for _,str := range strs {
		key := getKey(str)
		if _,ok := m[key] ; !ok {
			m[key] = []string{}
		}
		m[key] = append(m[key],str)
	}
	result := [][]string{}
	for _,v := range m {
		result = append(result,v)
	}
	return result
}